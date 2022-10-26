package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/coming-chat/go-aptos/aptosaccount"
	"github.com/coming-chat/go-aptos/aptosclient"
	txBuilder "github.com/coming-chat/go-aptos/transaction_builder"
)

//const rpcUrl = "https://fullnode.devnet.aptoslabs.com"

func main() {
	//mnemonic := "cargo emotion slot dentist client hint will penalty wrestle divide inform ranch"
	privateKey, err := hex.DecodeString("b60c04aeb238aa9994cd560d48491876e21baa8f29d3f95d485811766235bca6")
	if err != nil {
		log.Panic("failed to DecodeString", err)
	}
	account := aptosaccount.NewAccount(privateKey)
	fromAddress := "0x" + hex.EncodeToString(account.AuthKey[:])

	toAddress := "0x7ff115b61adcd12dd0248b716f7d49157da5093ecfc2f79a7d9ccdb22b73264e"
	amount := uint64(100000000)

	// Initialize the client
	restUrl := "https://fullnode.devnet.aptoslabs.com"
	client, err := aptosclient.Dial(context.Background(), restUrl)

	// Get Sender's account data and ledger info
	data, err := client.GetAccount(fromAddress)

	info, err := client.LedgerInfo()

	// Get gas price
	gasPrice, err := client.EstimateGasPrice()

	// Build paylod
	moduleName, err := txBuilder.NewModuleIdFromString("0x1::coin")
	toAddr, err := txBuilder.NewAccountAddressFromHex(toAddress)
	toAmountBytes := txBuilder.BCSSerializeBasicValue(amount)
	token, err := txBuilder.NewTypeTagStructFromString("0x1::aptos_coin::AptosCoin")
	payload := txBuilder.TransactionPayloadEntryFunction{
		ModuleName:   *moduleName,
		FunctionName: "transfer",
		TyArgs:       []txBuilder.TypeTag{*token},
		Args: [][]byte{
			toAddr[:], toAmountBytes,
		},
	}

	// Build transaction
	txn := &txBuilder.RawTransaction{
		Sender:                  account.AuthKey,
		SequenceNumber:          data.SequenceNumber,
		Payload:                 payload,
		MaxGasAmount:            2000,
		GasUnitPrice:            gasPrice,
		ExpirationTimestampSecs: info.LedgerTimestamp + 600,
		ChainId:                 uint8(info.ChainId),
	}

	// Sign raw transaction with account, and encode into data using BCS
	signedTxn, err := txBuilder.GenerateBCSTransaction(account, txn)

	// Submit transaction with BCS format.
	newTx, err := client.SubmitSignedBCSTransaction(signedTxn)
	if err != nil {
		log.Panic("failed to SubmitSignedBCSTransaction", err)
	}
	fmt.Printf("tx hash = %v\n", newTx.Hash)
}

func printLine(content string) {
	fmt.Printf("================= %s =================\n", content)
}

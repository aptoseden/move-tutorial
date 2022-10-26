package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/coming-chat/go-aptos/aptosaccount"
	"github.com/coming-chat/go-aptos/aptosclient"
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

	//toAddress := "0xcdbe33da8d218e97a9bec6443ba4a1b1858494f29142976d357f4770c384e015"
	//amount := uint64(1)

	// Initialize the client
	restUrl := "https://fullnode.devnet.aptoslabs.com"
	client, err := aptosclient.Dial(context.Background(), restUrl)
	if err != nil {
		log.Panic("failed to Dial ", err)
	}
	valdata, err := client.BalanceOf(fromAddress, "0x1::aptos_coin::AptosCoin")
	if err != nil {
		log.Panic("failed to get balance of ", err)
	}
	fmt.Println(valdata.Int64())

}

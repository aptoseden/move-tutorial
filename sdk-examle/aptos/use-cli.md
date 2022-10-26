# aptos cli使用笔记



参考链接：https://aptos.dev/nodes/local-testnet/using-cli-to-run-a-local-testnet

下载aptos cli

```sh
cd ~/bin
wget https://github.com/aptos-labs/aptos-core/releases/download/aptos-cli-v0.3.1a/aptos-cli-0.3.2-MacOSX-x86_64.zip --no-check-certificate
```

解压缩

```sh
unzip aptos-cli-0.3.2-MacOSX-x86_64.zip
which aptos
aptos help
```



获取客户端信息

```sh
$ aptos info
{
  "Result": {
    "build_branch": "devnet",
    "build_cargo_version": "cargo 1.62.1 (a748cf5a3 2022-06-08)",
    "build_commit_hash": "8399cd1c7b9662d3a6a09c28363c5f66f0839c41",
    "build_os": "macos-x86_64",
    "build_pkg_version": "0.3.2",
    "build_rust_channel": "1.62.1-x86_64-apple-darwin",
    "build_rust_version": "rustc 1.62.1 (e092d0b6b 2022-07-16)",
    "build_tag": "",
    "build_time": "2022-08-29 18:09:33 +00:00"
  }
}
```

启动节点

```sh
aptos node run-local-testnet --with-faucet

```

初始化本地local

```sh
aptos init --profile local --rest-url http://localhost:8080 --faucet-url http://localhost:8081

Configuring for profile local
Using command line argument for rest URL http://localhost:8080/
Using command line argument for faucet URL http://localhost:8081/
Enter your private key as a hex literal (0x...) [Current: None | No input: Generate new key (or keep one if present)]

No key given, generating key...
Account e718f50a9e4e6e77252a4b9a851aed2fd8bfa63f69a282cd5e5bc4862a1bb5b3 doesn't exist, creating it and funding it with 10000 coins
Aptos is now set up for account e718f50a9e4e6e77252a4b9a851aed2fd8bfa63f69a282cd5e5bc4862a1bb5b3!  Run `aptos help` for more information about commands
{
  "Result": "Success"
}
```

初始化

```sh

PROFILE=local
aptos init --profile $PROFILE --rest-url http://localhost:8080 --faucet-url http://localhost:8081
```



创建资源账户

```sh

aptos account create-resource-account --profile $PROFILE --seed 1
```



合约编译

```sh
aptos move compile --package-dir hello_blockchain --named-addresses hello_blockchain=$PROFILE 
```

响应

```sh
{
  "Result": [
    "e718f50a9e4e6e77252a4b9a851aed2fd8bfa63f69a282cd5e5bc4862a1bb5b3::message"
  ]
}
```



部署

```sh
aptos move publish --package-dir hello_blockchain --named-addresses hello_blockchain=$PROFILE  --profile $PROFILE
```



响应

```sh
{
  "Result": {
    "transaction_hash": "0x956fd6ef34a4a3768596d324f5520300836dc55812606b8fea4f0372e1869d6c",
    "gas_used": 182,
    "gas_unit_price": 1,
    "sender": "e718f50a9e4e6e77252a4b9a851aed2fd8bfa63f69a282cd5e5bc4862a1bb5b3",
    "sequence_number": 4,
    "success": true,
    "timestamp_us": 1665408942050405,
    "version": 201382,
    "vm_status": "Executed successfully"
  }
}
```



合约调用

 ```sh
aptos move run --function-id 28644355a529186d2a2402b0fc502bfc30c0e7fba2f4314ef143f8deb9c17fb5::message::set_message --args string:Hello!  --profile $PROFILE
 ```



响应

```sh
{
  "Result": {
    "transaction_hash": "0x205f5cfc23c365d0f6f60977a49b6159e7790311882286b38655a3a1cc450bf7",
    "gas_used": 39,
    "gas_unit_price": 1,
    "sender": "e718f50a9e4e6e77252a4b9a851aed2fd8bfa63f69a282cd5e5bc4862a1bb5b3",
    "sequence_number": 5,
    "success": true,
    "timestamp_us": 1665409355124101,
    "version": 206890,
    "vm_status": "Executed successfully"
  }
}
```

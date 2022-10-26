
初始化devnet账户

```sh
PROFILE=devnet
aptos account fund-with-faucet --profile $PROFILE --account $PROFILE
```

编译counter合约，需要注意counter所处的路径，笔者将其放在了aptos-core/aptos-move/move-examples下，因为它要引用相关的module。
```sh
aptos move compile --package-dir aptos-core/aptos-move/move-examples/counter/ --named-addresses MyCounterAddr=$PROFILE
```

部署合约
```sh
aptos move publish --package-dir aptos-core/aptos-move/move-examples/counter/ --named-addresses MyCounterAddr=$PROFILE  --profile $PROFILE
```

执行合约init_counter方法，需要将账户地址替换为自己的地址。
```sh
aptos move run --function-id 96f94d217c273061f5ea2b3aadea2f9dd9e6db5162655816dd679b05c5b20621::MyCounter::init_counter   --profile $PROFILE
```

执行合约的incr_counter方法
```sh
aptos move run --function-id 96f94d217c273061f5ea2b3aadea2f9dd9e6db5162655816dd679b05c5b20621::MyCounter::incr_counter   --profile $PROFILE
```
执行合约的incr_counter_by方法，携带参数。

```sh
aptos move run --function-id 96f94d217c273061f5ea2b3aadea2f9dd9e6db5162655816dd679b05c5b20621::MyCounter::incr_counter_by --args u64:5!   --profile $PROFILE
```

查看账户所拥有的资源
```
aptos account list --account $PROFILE --profile $PROFILE
```

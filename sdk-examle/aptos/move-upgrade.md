
# 可升级合约部署测试

### Move.toml
```toml
[package]
name = "pkgcounter"
version = "0.0.1"
upgrade_policy = "compatible"

[addresses]
counter = "_"

[dependencies]
AptosFramework = { local = "../../framework/aptos-framework" }
```

### counter.move
```move

module counter::MyCounter {
     use std::signer;

     struct Counter has key, store {
        value:u64,
     }
     public fun init(account: &signer){
        move_to(account, Counter{value:0});
     }
     public fun incr(account: &signer) acquires Counter {
        let counter = borrow_global_mut<Counter>(signer::address_of(account));
        counter.value = counter.value + 1;
     }

     public fun incr_by(account: &signer, increasement: u64) acquires Counter {
        let counter = borrow_global_mut<Counter>(signer::address_of(account));
        counter.value = counter.value + increasement;
     }

     public entry fun init_counter(account: signer){
        Self::init(&account)
     }

     public entry fun incr_counter(account: signer)  acquires Counter {
        Self::incr(&account)
     }
     
     public entry fun incr_counter_by(account: signer,increasement: u64)  acquires Counter {
        Self::incr_by(&account, increasement)
     }

     public entry fun incr_counter_by2(account: signer,increasement: u64, increasement2: u64)  acquires Counter {
        Self::incr_by(&account, increasement + increasement2)
     }
}
```


### 编译

```
PROFILE=local
aptos move compile --package-dir aptos-core/aptos-move/move-examples/counter/ --named-addresses coun
ter=$PROFILE
```

### 部署合约
```
aptos move publish --package-dir aptos-core/aptos-move/move-examples/counter/ --named-addresses counter=$PROFILE --profile $PROFILE
```

### 调用entry函数
```
aptos move run --function-id 4bb271917b3734bef8bcd4aec6d825dbb5c795c936d632901c329020d417b5cf::MyCounter::init_counter  --profile $PROFILE
```

### 多个参数调用
```
aptos move run --function-id 4bb271917b3734bef8bcd4aec6d825dbb5c795c936d632901c329020d417b5cf::MyCounter::incr_counter_by2 --args u64:3 --args u64:4  --profile $PROFILE
```

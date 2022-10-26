module MyCounterAddr::MyCounter {
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

     public entry fun get_counter(account: signer):u64  acquires Counter {
        let counter = borrow_global<Counter>(signer::address_of(&account));
        counter.value
     }
}
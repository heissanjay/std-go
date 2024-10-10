package main 

import (
  "fmt"
  "time"
  "sync"
)

/*
  sync package provides the basic synchronization primitives such as mutual exclusion locks.
  let have a look at its functions and types.

  O 
  - OnceFunc 
  - OnceValue 
  - OnceValues 
 C 
 - Cond type is condition variable that allows goroutines to wait untill a certain condition is true 
 - To create a condition, we need to pass a Locker interface implementation
 - Condition has 3 methods
   - Broadcast = Signals all waiting goroutines
   - Wait =   Block the current goroutine untill it receives the signal 
   - Signal = Singal a goroutine that is currently in waiting 

  - To create a condition variable, NewCond function is used.

 M 
 - Mutex 
   A Mutex type is a mutual exclusion lock. the zero value for mutex is unlocked mutex 
    - Lock 
    - TryLock (tried to lock and report whether it is succeeded or not)
    - Unlock [Unlocks the lock]
 - Map 
    Map is concurrently accessible map implementation in Go's sync package it does not require any external locking
    - Clear()
    - CompareAndDelete()
    - CompareAndSwap()
    - Delete()
    - Load()
    - LoadAndDelete()
    - LoadOrStore()
    - Range()
    - Store()
    - Swap()
 O 
 - Once 

 P 
 - Pool 

 R 
 - RWMutex  

 W 
 - WaitGroup 
 WaitGroup waits for a collections of goroutines to finish. 
 - The Main goroutine call `wg.Add()` with no of goroutines to wait for  and at the Same time wg.Wait() will be used to block untill all the goroutines has been finished
 - Each of the goroutines runs and when finishing they call wg.Done().

*/



func initConfiguration(){
  fmt.Println("Initiating configuration")
  time.Sleep(time.Second * 1)
  fmt.Println("Initialized configs")
}
  

func main(){
  //init := sync.OnceFunc(initConfiguration) // init can be run only once in the entire execution of the program
  //init() // run once 
  //init() // does nothing 
  //init() // absolutely does nothing 

  // For OnceValue let's go with the example given in the documentation 
/*
  once := sync.OnceValue(func() int{
    sum := 0
    for i := 0; i < 1000; i++ {
      sum += i
    }
    fmt.Println("Sum computed", sum)
    return sum 
  })

  now := time.Now()
  done := make(chan bool)
  for i := 0; i < 10; i++{
    go func(){
      const want = 499500
      got := once()
      if got != want {
        fmt.Println("Want", want, "GOT", got)
      }
      done <- true
    }()
  }
  for i := 0; i < 10; i++ {
    <- done 
  }

  close(done)
  fmt.Println("Time taken", time.Since(now))
  */ 
/*

  // Example for using Condition variable
  mu := &sync.Mutex{} // mutex implements the Locker Interface (Lock(), UnLock())
  cond := sync.NewCond(mu)
  var wg sync.WaitGroup
  wg.Add(2)

  go func(){
    defer wg.Done() 
    cond.L.Lock()
    fmt.Println("I am 1st goroutine, Waiting")
    cond.Wait()
    fmt.Println("Received signal - 1st goroutine")
    cond.L.Unlock()
  }()

  go func(){
    defer wg.Done()
    cond.L.Lock()
    fmt.Println("I am 2nd goroutine, Waiting")
    cond.Wait() 
    fmt.Println("Received the signal - 2nd goroutine")
    cond.L.Unlock()
  }()

  time.Sleep(time.Second * 2)
  cond.L.Lock()
  fmt.Println("Sending Signal")
  cond.Signal()
  //cond.Signal()
  cond.Broadcast()
  cond.L.Unlock()
  wg.Wait()
*/
/*
  var wg sync.WaitGroup 
  wg.Add(2)
  counter := 0
  //var mu sync.Mutex 
  go func(){
    defer wg.Done()
    //mu.Lock()
    for i := 0; i < 1000; i++{
      counter += 1 
    }
    //mu.Unlock()
  }()
  go func(){
   defer wg.Done()
   //mu.Lock()
   for i := 0; i < 500; i++ {
     counter -= 1
   }
   //mu.Unlock()
  }()
  wg.Wait()
  fmt.Println(counter)
*/ 
/*
 var mp sync.Map 

 mp.Store("India", "Delhi")
 mp.Store("United Kingdom", "London")


  ok := mp.CompareAndDelete("India", "Delhi")
  fmt.Println(ok)

 if value, ok := mp.Load("India"); ok {
    fmt.Println(value)
  }
 

val, ok = mp.LoadAndDelete("United Kingdom")
fmt.Println(val)
fmt.Println(ok)

fmt.Println(previous)
fmt.Println(ok)

mp.Clear()
mp.Range(func(key, value interface{})bool{
  fmt.Println(key, value)
  return true
})
*/
}

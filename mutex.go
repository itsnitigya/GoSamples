package main

import (
    "fmt"
    "math/rand"
    "sync"
    "sync/atomic"
    "time"
)

func main() {


	// state management using golang, to ensure at a time only one process/routine can access the state
    var state = make(map[int]int)

    var mutex = &sync.Mutex{}

    var readOps uint64
    var writeOps uint64

    for r := 0; r < 100; r++ {
        go func() {
            total := 0
            for {

                key := rand.Intn(5)
                mutex.Lock() // lock the state for usage
                total += state[key]
                mutex.Unlock() // unlock the state
                atomic.AddUint64(&readOps, 1)

                time.Sleep(time.Millisecond)
            }
        }()
    }

    for w := 0; w < 10; w++ {
        go func() {
            for {
                key := rand.Intn(5)
                val := rand.Intn(100)
                mutex.Lock() // waits for state to unlock and locks it again
                state[key] = val
                mutex.Unlock() // waits to execute the above given statement
                atomic.AddUint64(&writeOps, 1)
                time.Sleep(time.Millisecond)
            }
        }()
    }

    time.Sleep(time.Second) // should have used Sync groups instead

    readOpsFinal := atomic.LoadUint64(&readOps) // read the total number of ops maintained using atomic counters 
    fmt.Println("readOps:", readOpsFinal)
    writeOpsFinal := atomic.LoadUint64(&writeOps)
    fmt.Println("writeOps:", writeOpsFinal)

    mutex.Lock() // accessing the state safely 
    fmt.Println("state:", state)
    mutex.Unlock()
}
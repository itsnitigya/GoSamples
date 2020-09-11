package main

import (
    "fmt"
    "sync"
    "sync/atomic"
)

func main() {

    var ops uint64

    var wg sync.WaitGroup

    for i := 0; i < 50; i++ {
        wg.Add(1)

        go func() {
            for c := 0; c < 1000; c++ {

                atomic.AddUint64(&ops, 1) // ensures that ops is added atomically without goroutines messing with each other
            }
            wg.Done()
        }()
    }

    wg.Wait() // waits for the function to finish

    fmt.Println("ops:", ops) 
}
package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, wg *sync.WaitGroup) { // pass by reference

    defer wg.Done() // decreases count by 1

    fmt.Printf("Worker %d starting\n", id)

    time.Sleep(time.Second)
    fmt.Printf("Worker %d done\n", id)
}

func main() {

    var wg sync.WaitGroup

    for i := 1; i <= 5; i++ {
        wg.Add(1) // increases count by 1
        go worker(i, &wg)
    }

    wg.Wait() // waits till the count goes to 0
}
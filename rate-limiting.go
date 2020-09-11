package main

import (
    "fmt"
    "time"
)

func main() {

    requests := make(chan int, 5)
    for i := 1; i <= 5; i++ {
        requests <- i
    }
    close(requests)

    limiter := time.Tick(200 * time.Millisecond) // make the limiter tick every 200 ms

    for req := range requests {
        <-limiter // block using the channel
        fmt.Println("request", req, time.Now())
    }

    burstyLimiter := make(chan time.Time, 3) 

    for i := 0; i < 3; i++ {
        burstyLimiter <- time.Now() 
    }

    go func() {
        for t := range time.Tick(200 * time.Millisecond) {
            burstyLimiter <- t
        }
    }()

    burstyRequests := make(chan int, 5)
    for i := 1; i <= 5; i++ {
        burstyRequests <- i
    }
    close(burstyRequests)
    for req := range burstyRequests {
        <-burstyLimiter // fires off first requests immediately 
        fmt.Println("request", req, time.Now())
    }
}
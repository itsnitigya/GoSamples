package main

import (
	"fmt"
	"time"
)

func f(from string) {
    for i := 0; i < 3; i++ {
        fmt.Println(from, ":", i)
    }
}

func worker(done chan bool) {
    fmt.Print("working...")
    time.Sleep(time.Second)
    fmt.Println("done")

    done <- true
}

func ping(pings chan<- string, msg string) {
    pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
    msg := <-pings
    pongs <- msg
}

func main() {
	f("direct")

	go f("goroutine") // Executes the given function concurrently with the calling function

	time.Sleep(time.Second)
	fmt.Println("done")

	messages := make(chan string) // Communicate between goroutines

	// <- chan (recieve) 
	go func() { messages <- "ping" }()
	msg := <-messages 

	messages := make(chan string, 2) // Channels are unbuffered by default, this will accept only 2 strings
	
	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)

	done := make(chan bool, 1)
	go worker(done)
	<-done // Blocking till execution is finished

	pings := make(chan string, 1) // Passing messages between 2 channels
    pongs := make(chan string, 1)
    ping(pings, "passed message")
    pong(pings, pongs)
	fmt.Println(<-pongs)
	
	c1 := make(chan string)
	c2 := make(chan string)
	
	go func() {
        time.Sleep(1 * time.Second)
        c1 <- "one"
    }()
    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "two"
	}()
	
	for i := 0; i < 2; i++ {
        select { // Using the select statement to wait for the channels to recieve messages, Can do timeouts using select
			case msg1 := <-c1:
				fmt.Println("received", msg1)
			case msg2 := <-c2:
				fmt.Println("received", msg2)
        }
	}
	
	// Non blocking channel operations
	messages := make(chan string)
    signals := make(chan bool)
    select { // no blocking select
		case msg := <-messages:
			fmt.Println("received message", msg) 
		default:
			fmt.Println("no message received")
	}
	
	msg := "hi"
    select {
		case messages <- msg:
			fmt.Println("sent message", msg)
		default:
			fmt.Println("no message sent")
	}
	
	obs := make(chan int, 5)
    done := make(chan bool)
    go func() {
        for {
            j, more := <-jobs
            if more { // More depends if the channel is closed or not 
                fmt.Println("received job", j)
            } else {
                fmt.Println("received all jobs")
                done <- true
                return
            }
        }
    }()

    for j := 1; j <= 3; j++ {
        jobs <- j
        fmt.Println("sent job", j)
    }
    close(jobs) // close jobs, sets more to false
    fmt.Println("sent all jobs")
	<-done // Wait for jobs to finish
	

	queue := make(chan string, 2)
    queue <- "one"
    queue <- "two"
    close(queue)
    for elem := range queue {
		fmt.Println(elem) // Range over channels
	}
	
	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
	time.Sleep(2 * time.Second)

	// Timers are for when you want to do something once in the future - 
	// tickers are for when you want to do something repeatedly at regular intervals

	ticker := time.NewTicker(500 * time.Millisecond)
    done := make(chan bool)
    go func() {
        for {
            select {
            case <-done:
                return
            case t := <-ticker.C:
                fmt.Println("Tick at", t)
            }
        }
    }()
    time.Sleep(1600 * time.Millisecond) // Waiting for the goroutine to run
    ticker.Stop() // stops after 3 ticks
    done <- true
    fmt.Println("Ticker stopped")
}

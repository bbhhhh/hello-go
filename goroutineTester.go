package main

import (
	"fmt"
	"time"
)

var ch1 = make(chan int)

func main() {
	go Consumer()
	Producer()
}

func Producer() {
	for i := 0; i < 10; i++ {
		ch1 <- i
		fmt.Printf("Sent %d to ch1\n", i)
		time.Sleep(1000 * time.Millisecond)
	}
}

func Consumer() {
	for {
		s := <-ch1
		fmt.Printf("Received %d from ch1\n", s)
		//time.Sleep(100 * time.Millisecond)
	}
}

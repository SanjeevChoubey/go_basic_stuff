package main

import (
	"fmt"
	"math/rand"
	"time"
)

func producer(ch chan<- int) {
	for {
		time.Sleep(time.Duration(rand.Intn(7000)) * time.Millisecond)
		ch <- rand.Intn(10)
	}
}

func consumer(ch <-chan int, str string) {
	for n := range ch {
		fmt.Printf("%s--> %d\n", str, n)
	}
}

func fanout(chA <-chan int, chB, chC chan<- int) {
	for n := range chA {
		if n > 5 {
			chC <- n
		} else {
			chB <- n
		}
	}
}
func main() {
	chA := make(chan int)
	chB := make(chan int)
	chC := make(chan int)

	go producer(chA)
	go consumer(chB, "Channel B")
	go consumer(chC, "channel c")

	fanout(chA, chB, chC)
}

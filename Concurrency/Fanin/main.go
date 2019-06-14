package main

import (
	"fmt"
	"math/rand"
	"time"
)

func producer1(ch chan<- int) {
	for {
		n := rand.Intn(5)
		time.Sleep(
			time.Duration(rand.Intn(3000)) * time.Microsecond,
		)
		ch <- n

	}
}

func producer2(ch chan<- int) {
	for {
		n := rand.Intn(10)
		time.Sleep(
			time.Duration(rand.Intn(3000)) * time.Microsecond,
		)
		ch <- n

	}
}

func fanIn(ch1, ch2 <-chan int, ch3 chan<- int) {
	for {
		var n int
		select {
		case n = <-ch1:
			ch3 <- n
		case n = <-ch2:
			ch3 <- n

		}

	}
}

func consumer(ch <-chan int) {
	for n := range ch {
		fmt.Printf("<- %d\n", n)
	}

}
func main() {

	// fmt.Println("Processors : ", runtime.NumCPU())
	// return
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go producer1(ch1)
	go producer2(ch2)

	go consumer(ch3)
	fanIn(ch1, ch2, ch3)

}

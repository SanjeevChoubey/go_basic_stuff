package main

import (
	"fmt"
	"sync"
)

type queueList struct {
	values []int
	lock   sync.RWMutex
}

func (q *queueList) new() {
	q.values = []int{}
}

func (q *queueList) enqueue(value int) {
	q.lock.Lock()
	defer q.lock.Unlock()
	q.values = append(q.values, value)

}

func (q *queueList) dequeue() {
	q.lock.Lock()
	defer q.lock.Unlock()
	size := len(q.values)
	if size > 0 {
		q.values = q.values[1:size]
	}

}

func main() {
	q := queueList{}
	q.new()

	q.enqueue(28)
	q.enqueue(95)
	q.enqueue(65)
	q.enqueue(75)
	q.enqueue(9885)
	fmt.Println("queue list after enqueue is: ", q.values)
	q.dequeue()
	fmt.Println("queue list  after dequeue is: ", q.values)

}
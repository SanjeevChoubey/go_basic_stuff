package main

import (
	"fmt"
	"sync"
)

type stackList struct {
	value []int
	lock  sync.RWMutex
}

func (s *stackList) new() *stackList {
	s.value = []int{}
	return s
}

func (s *stackList) push(value int) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.value = append(s.value, value)
}

func (s *stackList) pop() {
	s.lock.Lock()
	defer s.lock.Unlock()
	size := len(s.value)
	s.value = s.value[0:(size - 1)]
}

func main() {
	s := stackList{}
	s.new()
	s.push(10)
	s.push(20)
	s.push(30)
	s.push(40)
	fmt.Println("Value after pushed", s.value)
	s.pop()
	fmt.Println("Value after pop", s.value)

}

package main

import (
	"fmt"
	"sync"
)

type list struct {
	head *node
	tail *node
	size int
	lock sync.RWMutex
}

type node struct {
	prv   *node
	value int
	next  *node
}

func (l *list) append(value int) {
	l.lock.Lock()
	defer l.lock.Unlock()
	n := &node{nil, value, nil}
	if l.head == nil {
		l.head = n
	} else {
		n.prv = l.tail
		l.tail.next = n

	}
	l.tail = n
}

func (n *node) nextElement() *node {
	return n.next
}

func (n *node) previousElement() *node {
	return n.prv
}

func (l *list) traverse(mode string) {
	switch mode {
	case "forward":
		fmt.Println("List in forward dircetion is as follows:")
		n := l.head
		for {
			if n == nil {
				break
			}
			fmt.Println(n.value)
			n = n.nextElement()
		}

	case "backward":
		fmt.Println("List in backward dircetion is as follows:")
		n := l.tail
		for {
			if n == nil {
				break
			}
			fmt.Println(n.value)
			n = n.previousElement()
		}

	}

}

func main() {
	l := list{}
	l.append(2)
	l.append(25)
	l.append(28)
	l.append(24)
	l.append(21)
	l.traverse("forward")
	l.traverse("backward")
}

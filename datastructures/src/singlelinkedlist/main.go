package main

import (
	"fmt"
)

type list struct {
	head *Node
	tail *Node
}
type Node struct {
	value int
	next  *Node
}

// get next node from list
func (n *Node) Next() *Node {
	return n.next
}

// add an element to list
func (l *list) push(value int) {
	node := &Node{value: value}
	if l.head == nil {
		l.head = node
	} else {
		l.tail.next = node
	}
	l.tail = node
}

// get the first node from list
func (l *list) first() *Node {
	return l.head
}

//Remove/pop an element
func (l *list) pop(value int) {
	n := l.first() // keep current node
	var pn *Node
	for {
		// if node is nil come out
		if n == nil {
			break
		}
		// if node value is same as passed value,
		if n.value == value {
			// when pop up head node
			if l.head == n {
				l.head = n.Next()
			}
			// when we pop up tail node
			if l.tail == n {
				l.tail = pn
			}
			//Take out this node
			if pn != nil {
				pn.next = n.Next()
			}

			break
		}
		pn = n
		n = n.Next()

	}
}

// Traverse through list
func (l *list) traverse() {
	n := l.first()
	for {
		fmt.Println(n.value)
		n = n.Next()
		if n == nil {
			break
		}
	}
}

func main() {
	l := list{}
	l.push(1)
	l.push(5)
	l.push(34)
	l.push(44)
	l.push(74)
	l.push(77)
	l.push(100)
	fmt.Println("Intial array")
	l.traverse()
	//Pop Elemnet
	fmt.Println("head Pop Up")
	l.pop(1) // Pop up head
	l.traverse()
	fmt.Println("tail Pop Up")
	l.pop(100) // pop up tail
	l.traverse()
	fmt.Println("other Pop Up")
	l.pop(34) // middle element
	l.traverse()

}

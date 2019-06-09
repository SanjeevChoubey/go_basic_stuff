package main

import (
  "fmt"
)
type list struct{
  head *Node
  tail *Node
}
type Node struct{
  value int
  next *Node
}
// get next node from list
func (n *Node) Next() *Node{
  return n.next
}

// add an element to list
func (l *list) push(value int) {
  node := &Node{value:value}
  if l.head == nil{
    l.head = node
  }else{
    l.tail.next = node
  }
  l.tail = node
}
// get the first node from list
func (l *list) first() *Node{
  return l.head
}
//Remove/pop an element
func (l *list) pop(value int){

  for n:= l.first(); n.Next() == nil;n.Next(){
    if n.Next().value == value{
      n.next = n.Next().Next()
      break
    }
  }
}
// Traverse through list
func (l *list) traverse(){
  n := l.first()
  for {
    fmt.Println(n.value)
    n = n.Next()
    if n ==  nil{
      break
    }
 }
}

func main(){
  l := list{}
  l.push(1)
  l.push(5)
  l.push(34)
  l.push(44)
  l.push(74)
  l.push(77)
  l.push(100)
  l.traverse()
  //Pop Elemnet
  l.pop(77)
  l.traverse()

}

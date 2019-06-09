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
func (l *list) pop(value int){

}
func main(){
  l := list{}
  l.push(1)
  l.push(5)
  l.push(34)
  n := l.first()
 for {
    fmt.Println(n.value)
    n = n.Next()
    if n ==  nil{
      break
    }
 }
}

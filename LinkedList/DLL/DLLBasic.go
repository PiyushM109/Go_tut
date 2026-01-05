package main

import "fmt"

type Node struct {
	value int
	next  *Node
	prev  *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) InsertAtEnd(value int) {
	newNode := &Node{value: value}

	if l.head == nil {
		l.head = newNode
		return
	}

	current := l.head
	for current.next != nil {
		current = current.next
	}

	current.next = newNode
	newNode.prev = current
}

func (l *LinkedList) Print() {
	current := l.head
	for current != nil {
		fmt.Printf("%d -> ", current.value)
		current = current.next
	}
	fmt.Println("nil")
}
func (l *LinkedList) ReverseDLL() {
	if l.head == nil || l.head.next == nil {
		return // nothing to reverse
	}

	curr := l.head
	var prev *Node = nil

	for curr != nil {
		next := curr.next
		curr.next = prev
		curr.prev = next
		prev = curr
		curr = next
	}

	l.head = prev
}

func main() {
	list := LinkedList{}
	list.InsertAtEnd(10)
	list.InsertAtEnd(20)
	list.InsertAtEnd(30)
	list.Print()

	list.ReverseDLL()
	list.Print()
}

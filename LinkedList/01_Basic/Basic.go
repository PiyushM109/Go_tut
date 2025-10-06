package main

import "fmt"

type Node struct {
	value int
	next  *Node
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
}

func (l *LinkedList) Print() {
	current := l.head
	for current != nil {
		fmt.Printf("%d -> ", current.value)
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	list := LinkedList{}
	list.InsertAtEnd(10)
	list.InsertAtEnd(20)
	list.InsertAtEnd(30)

	list.Print()
}

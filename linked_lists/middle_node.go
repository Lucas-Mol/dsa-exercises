package main

import "fmt"

//It's using the structs from linked_list.go. Please run it together or create a module

func middleNode(l DoublyLinkedList) *Node {
	head := l.head
	ahead := head

	for ahead != nil && ahead.next != nil {
		ahead = ahead.next.next
		head = head.next
	}

	return head
}

func main() {
	list := DoublyLinkedList{}
	list.AddToTail(1)
	list.AddToTail(2)
	list.AddToTail(3)
	list.AddToTail(4)
	list.AddToTail(5)
	list.AddToTail(6)
	list.AddToTail(7)
	list.AddToTail(8)
	list.AddToTail(9)
	list.AddToTail(10)

	fmt.Printf("%v\n", middleNode(list).value)
}

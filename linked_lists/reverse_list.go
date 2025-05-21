package main

import "fmt"

//It's using the structs from linked_list.go. Please run it together or create a module

func reverseList(l DoublyLinkedList) *DoublyLinkedList {
	var newList *Node

	for l.head != nil {
		nextNode := l.head.next
		l.head.next = newList
		newList = l.head
		l.head = nextNode
	}

	return &DoublyLinkedList{head: newList}
}

func main() {
	list := DoublyLinkedList{}
	list.AddToTail(1)
	list.AddToTail(2)
	list.AddToTail(3)

	reversedList := reverseList(list)
	for node := reversedList.head; node != nil; node = node.next {
		fmt.Println(node.value)
	}
}

package main

type Node struct {
	value any
	next  *Node
	prev  *Node
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
}

func (l *DoublyLinkedList) AddToHead(v any) {
	newNode := &Node{value: v}
	newNode.next = l.head
	if l.head == nil {
		l.tail = newNode
	} else {
		l.head.prev = newNode
	}
	l.head = newNode
}

func (l *DoublyLinkedList) AddToTail(v any) {
	newNode := &Node{value: v}
	newNode.prev = l.tail
	if l.tail == nil {
		l.head = newNode
	} else {
		l.tail.next = newNode
	}
	l.tail = newNode
}

func (l *DoublyLinkedList) RemoveFromHead() any {
	if l.head == nil {
		return nil
	}

	removedValue := l.head.value
	l.head = l.head.next
	if l.head != nil {
		l.head.prev = nil
	} else {
		l.tail = nil
	}
	return removedValue
}

func (l *DoublyLinkedList) RemoveFromTail() any {
	if l.tail == nil {
		return nil
	}

	removedValue := l.tail.value
	l.tail = l.tail.prev
	if l.tail != nil {
		l.tail.prev = nil
	} else {
		l.head = nil
	}
	return removedValue
}

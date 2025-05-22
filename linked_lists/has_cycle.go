package main

import "fmt"

func hasCycle(l DoublyLinkedList) bool {
	slow := l.head
	fast := l.head

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			return true
		}
	}
	return false
}

// Must have this function to test
func (l *DoublyLinkedList) addNodeToTail(newNode *Node) {
	newNode.prev = l.tail
	if l.tail == nil {
		l.head = newNode
	} else {
		l.tail.next = newNode
	}
	l.tail = newNode
}

func main() {
	list := DoublyLinkedList{}
	list.AddToTail(1)
	list.AddToTail(2)
	node := &Node{value: 3}
	list.addNodeToTail(node)
	list.AddToTail(4)
	list.addNodeToTail(node)

	fmt.Println(hasCycle(list))
}

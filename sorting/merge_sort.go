package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

func findMiddle(head *Node) *Node {
	if head == nil || head.next == nil {
		return head
	}

	slow := head
	fast := head
	var prev *Node

	for fast != nil && fast.next != nil {
		prev = slow
		slow = slow.next
		fast = fast.next.next
	}

	return prev
}

func merge(l1 *Node, l2 *Node) *Node {
	head := new(Node)
	tail := head

	for l1 != nil && l2 != nil {
		if l1.value < l2.value {
			tail.next = l1
			l1 = l1.next
		} else {
			tail.next = l2
			l2 = l2.next
		}
		tail = tail.next
	}
	if l1 != nil {
		tail.next = l1
	}

	if l2 != nil {
		tail.next = l2
	}

	return head.next
}

func MergeSort(head *Node) *Node {
	if head == nil || head.next == nil {
		return head
	}

	middle := findMiddle(head)
	afterMiddle := middle.next
	middle.next = nil

	left := MergeSort(head)
	right := MergeSort(afterMiddle)

	sorted := merge(left, right)
	return sorted
}

func main() {
	node7 := Node{value: 7}
	node1 := Node{value: 1, next: &node7}
	node3 := Node{value: 3, next: &node1}
	node9 := Node{value: 9, next: &node3}

	sortedHead := MergeSort(&node9)

	for sortedHead != nil {
		fmt.Println(sortedHead.value)
		sortedHead = sortedHead.next
	}
}

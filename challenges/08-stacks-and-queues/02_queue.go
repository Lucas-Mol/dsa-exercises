/*
A queue is an abstract data type that maintains the order in which elements were added to it, allowing the oldest elements to be removed from the front and new elements to be added to the rear. This is called a First-In-First-Out (FIFO) data structure because the first element added to the queue (i.e., the one that has been waiting the longest) is always the first one to be removed.

A basic queue has the following operations:

Enqueue: add a new element to the end of the queue.
Dequeue: remove the element from the front of the queue and return it.
In this challenge, you must first implement a queue using two stacks. Then process  queries, where each query is one of the following  types:

1 x: Enqueue element  into the end of the queue.
2: Dequeue the element at the front of the queue.
3: Print the element at the front of the queue.
For example, a series of queries might be as follows:

image

Function Description

Complete the put, pop, and peek methods in the editor below. They must perform the actions as described above.

Input Format

The first line contains a single integer, , the number of queries.

Each of the next  lines contains a single query in the form described in the problem statement above. All queries start with an integer denoting the query , but only query  is followed by an additional space-separated value, , denoting the value to be enqueued.

Constraints

It is guaranteed that a valid answer always exists for each query of types  and .
Output Format

For each query of type , return the value of the element at the front of the fifo queue on a new line.

Sample Input

10
1 42
2
1 14
3
1 28
3
1 60
1 78
2
2
Sample Output

14
14
Explanation

image
*/

package main

import (
	"fmt"
	"strconv"
)

type Queue[T any] struct {
	items []T
}

func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	if len(q.items) == 0 {
		var zeroItem T
		return zeroItem, false
	}

	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

func (q *Queue[T]) Peek() (T, bool) {
	if len(q.items) == 0 {
		var zeroItem T
		return zeroItem, false
	}

	item := q.items[0]
	return item, true
}

type QueueOperation struct {
	op    int
	value int
}

func NewQueueOperation(input []string) *QueueOperation {
	var op int
	var value int

	if len(input) == 1 {
		op, _ = strconv.Atoi(input[0])
	} else if len(input) > 1 {
		op, _ = strconv.Atoi(input[0])
		value, _ = strconv.Atoi(input[1])
	}

	return &QueueOperation{
		op,
		value,
	}
}

func operateQueue(queue *Queue[int], operation QueueOperation) {
	switch operation.op {
	case 1:
		queue.Enqueue(operation.value)
	case 2:
		queue.Dequeue()
	case 3:
		value, exists := queue.Peek()
		if exists {
			fmt.Println(value)
		}
	}
}

// func main() {
// 	reader := bufio.NewReader(os.Stdin)
// 	val, _, _ := reader.ReadLine()
// 	n, _ := strconv.Atoi(string(val))
// 	var queue Queue[int]

// 	for range n {
// 		val, _, _ := reader.ReadLine()
// 		input := strings.Split(string(val), " ")
// 		operationQueue := NewQueueOperation(input)
// 		operateQueue(&queue, *operationQueue)
// 	}
// }

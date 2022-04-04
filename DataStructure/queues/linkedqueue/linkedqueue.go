package linkedqueue

import "DataStructure/lists/linkedlist"

type Queue struct {
	head   *linkedlist.Node
	tail   *linkedlist.Node
	length int
}

func NewQueue() *Queue {
	return &Queue{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}

func (q *Queue) Enqueue(v interface{}) {
	node := &linkedlist.Node{
		Data: nil,
		Next: nil,
	}
	if q.tail == nil {
		q.tail = node
		q.head = node
	} else {
		q.tail.Next = node
		q.tail = node
	}
	q.length++
}

func (q *Queue) Dequeue() interface{} {
	if q.head == nil {
		return nil
	}
	v := q.head.Data
	q.head = q.head.Next
	q.length--
	return v
}

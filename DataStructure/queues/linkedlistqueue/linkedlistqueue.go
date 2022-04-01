// Package linkedlistqueue implement the queue base from linkedlist
// also can implement queue from linked
package linkedlistqueue

import "DataStructure/lists/linkedlist"

type Queue struct {
	list *linkedlist.List
	head int
	tail int
}

func New() *Queue {
	return &Queue{
		list: linkedlist.New(),
		head: 0,
		tail: 0,
	}
}

func (q *Queue) Enqueue(value interface{}) {
	q.list.Append(value)
	q.tail++
}

func (q *Queue) Dequeue() interface{} {
	if q.head == q.tail {
		return nil
	}
	value := q.list.SearchByIndex(q.head)
	q.head++
	return value
}

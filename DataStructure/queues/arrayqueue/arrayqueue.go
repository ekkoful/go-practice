package arrayqueue

import "DataStructure/lists/arraylist"

type Queue struct {
	list *arraylist.ArrayList
	head int
	tail int
}

func New() *Queue {
	return &Queue{
		list: arraylist.New(),
		head: 0,
		tail: 0,
	}
}

func (q *Queue) Enqueue(value interface{}) {
	q.list.Add(value)
	q.tail++
}

func (q *Queue) Dequeue() interface{} {
	if q.head == q.tail {
		return nil
	}
	value, ok := q.list.Get(q.head)
	if !ok {
		return nil
	}
	q.head++
	return value
}

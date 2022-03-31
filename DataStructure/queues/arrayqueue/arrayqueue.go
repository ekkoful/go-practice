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

func (q *Queue) Enqueue(value ...interface{}) {
	q.list.Add(value...)
}

func (q *Queue) Dequeue() {
	q.list.Remove(q.head)
}

package linkedlistqueue

import "testing"

func TestQueue_Enqueue(t *testing.T) {
	queue := New()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	if actualValue := queue.tail; actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	if actualValue := queue.head; actualValue != 0 {
		t.Errorf("Got %v expected %v", actualValue, 0)
	}
	t.Logf("queue list: %#v, queue head: %d, queue tail: %d", queue.list, queue.head, queue.tail)
}

func TestQueue_Dequeue(t *testing.T) {
	queue := New()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	if actualValue := queue.tail; actualValue != 4 {
		t.Errorf("Got %v expected %v", actualValue, 4)
	}
	if actualValue := queue.head; actualValue != 0 {
		t.Errorf("Got %v expected %v", actualValue, 0)
	}
	t.Logf("queue list: %#v, queue head: %d, queue tail: %d", queue.list, queue.head, queue.tail)

	// queue.Dequeue()
	if actualValue := queue.head; actualValue != 0 {
		t.Errorf("Got %v expected %v", actualValue, 0)
	}
	if actualValue := queue.Dequeue(); actualValue != 1 {
		t.Errorf("Got %v expected %v", actualValue, 1)
	}

	// queue.Dequeue()
	if actualValue := queue.head; actualValue != 1 {
		t.Errorf("Got %v expected %v", actualValue, 1)
	}
	if actualValue := queue.Dequeue(); actualValue != 2 {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}

	// queue.Dequeue()
	if actualValue := queue.head; actualValue != 2 {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}
	if actualValue := queue.Dequeue(); actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}

	// queue.Dequeue()
	if actualValue := queue.head; actualValue != 3 {
		t.Errorf("Got %v expected %v", actualValue, 3)
	}
	if actualValue := queue.Dequeue(); actualValue != 4 {
		t.Errorf("Got %v expected %v", actualValue, 4)
	}

	if actualValue := queue.tail; actualValue != 4 {
		t.Errorf("Got %v expected %v", actualValue, 4)
	}
	if actualValue := queue.head; actualValue != 4 {
		t.Errorf("Got %v expected %v", actualValue, 0)
	}
	t.Logf("queue list: %#v, queue head: %d, queue tail: %d", queue.list, queue.head, queue.tail)
}

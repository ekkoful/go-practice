package arrayqueue

import (
	"testing"
)

func TestQueue(t *testing.T) {
	queue := New()
	queue.Enqueue(1, 4, 5, 3)
	queue.Enqueue(2)
	queue.Dequeue()
}

package arraystack

import (
	"DataStructure/lists/arraylist"
)

type Stack struct {
	list *arraylist.ArrayList
}

func New() *Stack {
	stack := &Stack{
		list: arraylist.New(),
	}
	return stack
}

func (stack *Stack) Push(value ...interface{}) {
	stack.list.Add(value)
}

func (stack *Stack) Pop() (value interface{}, ok bool) {
	value, ok = stack.list.Get(stack.list.Size() - 1)
	stack.list.Remove(stack.list.Size() - 1)
	return
}

func (stack *Stack) Empty() bool {
	return stack.list.IsEmpty()
}

func (stack *Stack) Size() int {
	return stack.list.Size()
}

func (stack *Stack) Clear() {
	stack.list.Clear()
}

func (stack *Stack) Peek() (value interface{}, ok bool) {
	return stack.list.Get(stack.list.Size() - 1)
}

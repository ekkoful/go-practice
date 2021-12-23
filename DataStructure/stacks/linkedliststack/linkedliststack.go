package linkedliststack

import "DataStructure/lists/linkedlist"

type Stack struct {
	list *linkedlist.List
}

func New() *Stack {
	stack := &Stack{
		list: linkedlist.New(),
	}
	return stack
}

func (stack *Stack) Push(value ...interface{}) {
	stack.list.Append(value)
}

func (stack *Stack) Pop() (value interface{}, ok bool) {
	value = stack.list.SearchByIndex(stack.list.Length() - 1)
	stack.list.RemoveByIndex(stack.list.Length() - 1)
	return
}

func (stack *Stack) Empty() bool {
	return stack.list.IsEmpty()
}

func (stack *Stack) Size() int {
	return stack.list.Length()
}

func (stack *Stack) Clear() {
	stack.list = linkedlist.New()
}

func (stack *Stack) Peek() (value interface{}) {
	return stack.list.SearchByIndex(stack.list.Length() - 1)
}

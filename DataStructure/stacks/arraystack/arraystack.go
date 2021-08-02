package arraystack

import (
	"DataStructure/lists/arraylist"
)

type Stack struct {
	stack *arraylist.ArrayList
}

func New() *Stack {
	stack := &Stack{}
	return stack
}

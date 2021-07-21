package main

import (
	"fmt"
)

type Elem interface{}

type Node struct {
	Data Elem
	Next *Node
}

type List struct {
	headNode *Node
}

func New() *List {
	return &List{}
}

func (list *List) IsEmpty() bool {
	if list.headNode == nil {
		return true
	} else {
		return false
	}
}

func (list *List) Add(data Elem) *Node {
	node := &Node{
		Data: data,
	}
	node.Next = list.headNode
	list.headNode = node
	return node
}

func (list *List) Append(data Elem) *Node {
	node := &Node{
		Data: data,
	}
	if list.IsEmpty() {
		list.headNode = node
	} else {
		cur := list.headNode
		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = node
	}
	return node
}

func (list *List) ShowList() {
	if !list.IsEmpty() {
		cur := list.headNode
		// 只要当前Node不为空，就进行遍历
		for cur != nil {
			fmt.Println(cur.Data)
			cur = cur.Next
		}
	}
}

func (list *List) Length() int {
	cur := list.headNode
	count := 0
	for cur != nil {
		count++
		cur = cur.Next
	}
	return count
}

func (list *List) Insert(pos int, data Elem) {

}

func (list *List) SearchByIndex() Elem {

}

func (list *List) SearchByValue() bool {

}

func (list *List) RemoveByIndex() {

}

func (list *List) RemoveByValue() {

}

func main() {
	l := New()
	fmt.Println(l.IsEmpty())
	l.Add(1)
	l.Append(2)
	l.ShowList()
	fmt.Println(l.Length())
}

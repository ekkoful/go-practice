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

//在链表指定位置前插入元素
func (list *List) Insert(index int, data Elem) {
	//如果index < 0, 则进头部插入
	if index < 0 {
		list.Add(data)
	} else if index > list.Length() { //如果index大于链表长度，进行尾部插入
		list.Append(data)
	} else {
		pre := list.headNode
		count := 0
		for count < (index - 1) { //用于控制位移的链表数目
			count++
			pre = pre.Next
		}
		// 当循环退出后，pre指向index-1的位置
		node := &Node{
			Data: data,
		}
		node.Next = pre.Next
		pre.Next = node
	}
}

func (list *List) SearchByIndex(index int) Elem {
	if index <= 0 {
		return nil
	} else if index > list.Length() {
		return nil
	} else {
		cur := list.headNode
		count := 0
		for count < (index - 1) {
			count++
			cur = cur.Next
		}
		return cur.Data
	}
}

func (list *List) SearchByValue(data Elem) bool {
	cur := list.headNode
	for cur != nil {
		if cur.Data == data {
			return true
		}
		cur = cur.Next
	}
	return false
}

func (list *List) RemoveByIndex(index int) bool {
	if index <= 0 {
		return false
	} else if index > list.Length() {
		return false
	} else {
		cur := list.headNode
		count := 1
		for count < (index - 1) {
			count++
			cur = cur.Next
		}
		cur.Next = cur.Next.Next
		return false
	}
}

func (list *List) RemoveByValue(data Elem) bool {
	cur := list.headNode
	for cur != nil {
		if cur.Data == data {
			cur.Next = cur.Next.Next
			return true
		}
		cur = cur.Next
	}
	return false
}

func main() {
	l := New()
	fmt.Println(l.IsEmpty())
	l.Add(1)
	l.Append(2)
	l.Append(3)
	l.Append(4)
	l.Append(5)
	l.Append(6)
	l.ShowList()
	fmt.Println(l.Length())
	fmt.Println(l.SearchByIndex(3))
	fmt.Println(l.SearchByValue(1))
	fmt.Println(l.RemoveByIndex(2))
	l.ShowList()
	fmt.Println(l.RemoveByValue(5))
	l.ShowList()
}

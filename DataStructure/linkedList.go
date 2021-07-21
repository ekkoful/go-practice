package linkedList

type Elem interface{}

type Node struct {
	Data Elem
	Next *Node
}

type List struct {
	headNode *Node
}

func (list *List) New() {
	return &List{
		headNode: &Node{
			Data: nil,
			Next: *Node,
		},
	}
}

func (list *List) IsEmpty() bool {
	if list.headNode == nil {
		return true 
	} else if {
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
	} else if {
		curr := list.headNode
		for curr.Next != nil {
			curr = curr.Next
		}
		curr.Next = Node
	}
}

func main() {
	l = linkedList.New()
}
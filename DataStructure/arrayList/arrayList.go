package arraylist

type ArrayList struct {
	elements []interface{}
	size     int
}

const (
	growthFactor = float32(2.0)
	shrinkFactor = float32(0.25)
)

func New(values ...interface{}) *ArrayList {
	list := &ArrayList{}
	//list.Element = make([]interface{}, 10)
	if len(values) > 0 {
		list.Add(values)
	}
	return list
}

func (list *ArrayList) Add(values ...interface{}) {
	list.growBy(len(values))
	for _, value := range values {
		list.elements[list.size] = value
		list.size++
	}
}

func (list *ArrayList) resize(cap int) {
	newElements := make([]interface{}, cap, cap)
	copy(newElements, list.elements)
	list.elements = newElements
}

func (list *ArrayList) growBy(n int) {
	currentCapacity := cap(list.elements)
	if list.size+n >= currentCapacity {
		newCapacity := int(growthFactor * float32(currentCapacity+n))
		list.resize(newCapacity)
	}
}

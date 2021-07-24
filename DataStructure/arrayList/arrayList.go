package arraylist

type ArrayList struct {
	elements []interface{}
	size     int
}

const (
	GROWTH_FACTOR = float32(2.0)
	//shrinkFactor = float32(0.25)
	SHRINK_FACTOR = float32(0.25)
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

func (list *ArrayList) Remove(index int) {
	if !list.withinRange(index) {
		return
	}

	list.elements[index] = nil //cleanup reference
	copy(list.elements[index:], list.elements[index+1:list.size])
	list.size--

	list.shrink()
}

func (list *ArrayList) Get(index int) (interface{}, bool) {

	if !list.withinRange(index) {
		return nil, false
	}

	return list.elements[index], true
}

func (list *ArrayList) Contains(values ...interface{}) bool {

	for _, searchValue := range values {
		found := false
		for _, element := range list.elements {
			if element == searchValue {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func (list *ArrayList) withinRange(index int) bool {
	return index >= 0 && index < list.size
}

func (list *ArrayList) resize(cap int) {
	newElements := make([]interface{}, cap)
	copy(newElements, list.elements)
	list.elements = newElements
}

func (list *ArrayList) growBy(n int) {
	currentCapacity := cap(list.elements)
	if list.size+n >= currentCapacity {
		newCapacity := int(GROWTH_FACTOR * float32(currentCapacity+n))
		list.resize(newCapacity)
	}
}

func (list *ArrayList) shrink() {

	if SHRINK_FACTOR == 0.0 {
		return
	}
	currentCapacity := cap(list.elements)
	if list.size <= int(SHRINK_FACTOR*float32(currentCapacity)) {
		list.resize(list.size)
	}
}

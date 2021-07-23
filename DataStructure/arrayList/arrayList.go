package arraylist

type ArrayList struct {
	element []interface{}
	size    int
}

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
		list.element[list.size] = value
		list.size++
	}
}

func (list *ArrayList) growBy(s int) {

}

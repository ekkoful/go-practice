package arraylist

type ArrayList struct {
	Element []interface{}
	Size    int
}

func New(values ...interface{}) *ArrayList {
	list := &ArrayList{}
	list.Element = make([]interface{}, 10)

}

package arraylist

import (
	"testing"
)

func TestListNew(t *testing.T) {
	list1 := New()
	if actualValue := list1.IsEmpty(); actualValue != true {
		t.Errorf("Got %v expected %v", actualValue, true)
	}

	list2 := New(1, "b")
	if actualValue := list2.Size(); actualValue != 2 {
		t.Errorf("Got %v expected %v", actualValue, 2)
	}

}

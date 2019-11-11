package ordered_queue

import (
	"testing"
)

type MyData struct {
	A int
	B string
}

func (this MyData) LessThan(other QueueElement) bool {
	if this.A < other.(MyData).A {
		return true
	}
	return false
}

func TestLoadConfigurationFromPathInvalidPath(t *testing.T) {
	q := OrderedQueue{}
	q.Push(MyData{3, "asdf"})
	q.Push(MyData{7, "qwer"})
	q.Push(MyData{1, "zxcv"})
	if q.Size() != 3 {
		t.Errorf("The queue should contain 3 elements")
	}
	if value, err := q.Pop(); err != nil {
		t.Errorf("Value should not be nil")
	} else if castedValue := value.(MyData); castedValue.A != 1 {
		t.Errorf("Value A should be 1, instead it is %d", castedValue.A)
	} else if castedValue.B != "zxcv" {
		t.Errorf("Value A should be zxcv, instead it is %s", castedValue.B)
	}
	if value, err := q.Pop(); err != nil {
		t.Errorf("Value should not be nil")
	} else if castedValue := value.(MyData); castedValue.A != 3 {
		t.Errorf("Value A should be 3, instead it is %d", castedValue.A)
	} else if castedValue.B != "asdf" {
		t.Errorf("Value A should be asdf, instead it is %s", castedValue.B)
	}
	if value, err := q.Pop(); err != nil {
		t.Errorf("Value should not be nil")
	} else if castedValue := value.(MyData); castedValue.A != 7 {
		t.Errorf("Value A should be 7, instead it is %d", castedValue.A)
	} else if castedValue.B != "qwer" {
		t.Errorf("Value A should be qwer, instead it is %s", castedValue.B)
	}
	if _, err := q.Pop(); err == nil {
		t.Errorf("Pop() on empty queue should return an error")
	}
}

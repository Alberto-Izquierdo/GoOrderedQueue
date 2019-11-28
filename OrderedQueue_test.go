package ordered_queue

import (
	"testing"
)

func TestQueueWithInts(t *testing.T) {
	q := OrderedQueue{}
	if err := q.Push(3); err != nil {
		t.Errorf("Pushing correct values should not return an error")
	}
	if err := q.Push(7); err != nil {
		t.Errorf("Pushing correct values should not return an error")
	}
	if err := q.Push(1); err != nil {
		t.Errorf("Pushing correct values should not return an error")
	}
	if q.Size() != 3 {
		t.Errorf("The queue should contain 3 elements")
	}
	if value, err := q.Pop(); err != nil {
		t.Errorf("Value should not be nil")
	} else if castedValue := value.(int); castedValue != 1 {
		t.Errorf("First value should be 1, instead it is %d", castedValue)
	}
	if value, err := q.Pop(); err != nil {
		t.Errorf("Value should not be nil")
	} else if castedValue := value.(int); castedValue != 3 {
		t.Errorf("First value should be 3, instead it is %d", castedValue)
	}
	if value, err := q.Pop(); err != nil {
		t.Errorf("Value should not be nil")
	} else if castedValue := value.(int); castedValue != 7 {
		t.Errorf("First value should be 7, instead it is %d", castedValue)
	}
	if _, err := q.Pop(); err == nil {
		t.Errorf("Pop() on empty queue should return an error")
	}
}

type NewStruct struct {
	value int
}

func TestQueueWithWrongStruct(t *testing.T) {
	q := OrderedQueue{}
	var value NewStruct
	err := q.Push(value)
	if err == nil {
		t.Errorf("Pushing wrong struct should return an error")
	}
}

func TestQueueWithStrings(t *testing.T) {
	q := OrderedQueue{}
	if err := q.Push("asdf"); err != nil {
		t.Errorf("Pushing correct values should not return an error")
	}
	if err := q.Push("zxcv"); err != nil {
		t.Errorf("Pushing correct values should not return an error")
	}
	if err := q.Push("qwer"); err != nil {
		t.Errorf("Pushing correct values should not return an error")
	}
	if q.Size() != 3 {
		t.Errorf("The queue should contain 3 elements")
	}
	if value, err := q.Pop(); err != nil {
		t.Errorf("Value should not be nil")
	} else if castedValue := value.(string); castedValue != "asdf" {
		t.Errorf("First value should be asdf, instead it is %s", castedValue)
	}
	if value, err := q.Pop(); err != nil {
		t.Errorf("Value should not be nil")
	} else if castedValue := value.(string); castedValue != "qwer" {
		t.Errorf("First value should be qwer, instead it is %s", castedValue)
	}
	if value, err := q.Pop(); err != nil {
		t.Errorf("Value should not be nil")
	} else if castedValue := value.(string); castedValue != "zxcv" {
		t.Errorf("First value should be zxcv, instead it is %s", castedValue)
	}
	if _, err := q.Pop(); err == nil {
		t.Errorf("Pop() on empty queue should return an error")
	}
}

type MyData struct {
	A int
	B string
}

func (this MyData) LessThan(other QueueElement) bool {
	return this.A < other.(MyData).A
}

func TestQueueWithCustomStruct(t *testing.T) {
	q := OrderedQueue{}
	if err := q.Push(MyData{3, "asdf"}); err != nil {
		t.Errorf("Pushing correct values should not return an error")
	}
	if err := q.Push(MyData{7, "qwer"}); err != nil {
		t.Errorf("Pushing correct values should not return an error")
	}
	if err := q.Push(MyData{1, "zxcv"}); err != nil {
		t.Errorf("Pushing correct values should not return an error")
	}
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

func TestQueueWithWrongValues(t *testing.T) {
	q := OrderedQueue{}
	if err := q.Push(3); err != nil {
		t.Errorf("Pushing correct values should not return an error")
	}
	if err := q.Push("asdf"); err == nil {
		t.Errorf("Pushing wrong values should return an error")
	}
	if q.Size() != 1 {
		t.Errorf("The queue should contain 1 element")
	}
}

func TestQueueWithMultipleTypes(t *testing.T) {
	q := OrderedQueue{}
	{
		vector := []int{1, 5, 3}
		for _, element := range vector {
			if err := q.Push(element); err != nil {
				t.Errorf("Pushing correct values should not return an error")
			}
		}
		if q.Size() != 3 {
			t.Errorf("The queue should contain 3 elements, instead int contains %d elements", q.Size())
		}
		for _, _ = range vector {
			if _, err := q.Pop(); err != nil {
				t.Errorf("Pop() should not return an error")
			}
		}
		if q.Size() != 0 {
			t.Errorf("The queue should contain 0 elements")
		}
	}
	{
		vector := []int8{1, 5, 3}
		for _, element := range vector {
			if err := q.Push(element); err != nil {
				t.Errorf("Pushing correct values should not return an error")
			}
		}
		if q.Size() != 3 {
			t.Errorf("The queue should contain 3 elements, instead int contains %d elements", q.Size())
		}
		for _, _ = range vector {
			if _, err := q.Pop(); err != nil {
				t.Errorf("Pop() should not return an error")
			}
		}
		if q.Size() != 0 {
			t.Errorf("The queue should contain 0 elements")
		}
	}
	{
		vector := []int16{1, 5, 3}
		for _, element := range vector {
			if err := q.Push(element); err != nil {
				t.Errorf("Pushing correct values should not return an error")
			}
		}
		if q.Size() != 3 {
			t.Errorf("The queue should contain 3 elements, instead int contains %d elements", q.Size())
		}
		for _, _ = range vector {
			if _, err := q.Pop(); err != nil {
				t.Errorf("Pop() should not return an error")
			}
		}
		if q.Size() != 0 {
			t.Errorf("The queue should contain 0 elements")
		}
	}
	{
		vector := []int32{1, 5, 3}
		for _, element := range vector {
			if err := q.Push(element); err != nil {
				t.Errorf("Pushing correct values should not return an error")
			}
		}
		if q.Size() != 3 {
			t.Errorf("The queue should contain 3 elements, instead int contains %d elements", q.Size())
		}
		for _, _ = range vector {
			if _, err := q.Pop(); err != nil {
				t.Errorf("Pop() should not return an error")
			}
		}
		if q.Size() != 0 {
			t.Errorf("The queue should contain 0 elements")
		}
	}
	{
		vector := []int64{1, 5, 3}
		for _, element := range vector {
			if err := q.Push(element); err != nil {
				t.Errorf("Pushing correct values should not return an error")
			}
		}
		if q.Size() != 3 {
			t.Errorf("The queue should contain 3 elements, instead int contains %d elements", q.Size())
		}
		for _, _ = range vector {
			if _, err := q.Pop(); err != nil {
				t.Errorf("Pop() should not return an error")
			}
		}
		if q.Size() != 0 {
			t.Errorf("The queue should contain 0 elements")
		}
	}
	{
		vector := []uint{1, 5, 3}
		for _, element := range vector {
			if err := q.Push(element); err != nil {
				t.Errorf("Pushing correct values should not return an error")
			}
		}
		if q.Size() != 3 {
			t.Errorf("The queue should contain 3 elements, instead int contains %d elements", q.Size())
		}
		for _, _ = range vector {
			if _, err := q.Pop(); err != nil {
				t.Errorf("Pop() should not return an error")
			}
		}
		if q.Size() != 0 {
			t.Errorf("The queue should contain 0 elements")
		}
	}
	{
		vector := []uint8{1, 5, 3}
		for _, element := range vector {
			if err := q.Push(element); err != nil {
				t.Errorf("Pushing correct values should not return an error")
			}
		}
		if q.Size() != 3 {
			t.Errorf("The queue should contain 3 elements, instead int contains %d elements", q.Size())
		}
		for _, _ = range vector {
			if _, err := q.Pop(); err != nil {
				t.Errorf("Pop() should not return an error")
			}
		}
		if q.Size() != 0 {
			t.Errorf("The queue should contain 0 elements")
		}
	}
	{
		vector := []uint16{1, 5, 3}
		for _, element := range vector {
			if err := q.Push(element); err != nil {
				t.Errorf("Pushing correct values should not return an error")
			}
		}
		if q.Size() != 3 {
			t.Errorf("The queue should contain 3 elements, instead int contains %d elements", q.Size())
		}
		for _, _ = range vector {
			if _, err := q.Pop(); err != nil {
				t.Errorf("Pop() should not return an error")
			}
		}
		if q.Size() != 0 {
			t.Errorf("The queue should contain 0 elements")
		}
	}
	{
		vector := []uint32{1, 5, 3}
		for _, element := range vector {
			if err := q.Push(element); err != nil {
				t.Errorf("Pushing correct values should not return an error")
			}
		}
		if q.Size() != 3 {
			t.Errorf("The queue should contain 3 elements, instead int contains %d elements", q.Size())
		}
		for _, _ = range vector {
			if _, err := q.Pop(); err != nil {
				t.Errorf("Pop() should not return an error")
			}
		}
		if q.Size() != 0 {
			t.Errorf("The queue should contain 0 elements")
		}
	}
	{
		vector := []uint64{1, 5, 3}
		for _, element := range vector {
			if err := q.Push(element); err != nil {
				t.Errorf("Pushing correct values should not return an error")
			}
		}
		if q.Size() != 3 {
			t.Errorf("The queue should contain 3 elements, instead int contains %d elements", q.Size())
		}
		for _, _ = range vector {
			if _, err := q.Pop(); err != nil {
				t.Errorf("Pop() should not return an error")
			}
		}
		if q.Size() != 0 {
			t.Errorf("The queue should contain 0 elements")
		}
	}
	{
		vector := []float32{1.0, 5.0, 3.0}
		for _, element := range vector {
			if err := q.Push(element); err != nil {
				t.Errorf("Pushing correct values should not return an error")
			}
		}
		if q.Size() != 3 {
			t.Errorf("The queue should contain 3 elements, instead int contains %d elements", q.Size())
		}
		for _, _ = range vector {
			if _, err := q.Pop(); err != nil {
				t.Errorf("Pop() should not return an error")
			}
		}
		if q.Size() != 0 {
			t.Errorf("The queue should contain 0 elements")
		}
	}
	{
		vector := []float64{1.0, 5.0, 3.0}
		for _, element := range vector {
			if err := q.Push(element); err != nil {
				t.Errorf("Pushing correct values should not return an error")
			}
		}
		if q.Size() != 3 {
			t.Errorf("The queue should contain 3 elements, instead int contains %d elements", q.Size())
		}
		for _, _ = range vector {
			if _, err := q.Pop(); err != nil {
				t.Errorf("Pop() should not return an error")
			}
		}
		if q.Size() != 0 {
			t.Errorf("The queue should contain 0 elements")
		}
	}
}

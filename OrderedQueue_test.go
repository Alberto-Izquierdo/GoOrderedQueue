package ordered_queue

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueueWithInts(t *testing.T) {
	q := OrderedQueue{}
	assert.Nil(t, q.Push(3), "Pushing correct values should not return an error")
	assert.Nil(t, q.Push(7), "Pushing correct values should not return an error")
	assert.Nil(t, q.Push(1), "Pushing correct values should not return an error")
	assert.Equal(t, 3, q.Size(), "The queue should contain 3 elements")
	value, err := q.Pop()
	assert.Nil(t, err, "Error should be nil when poping elements from queue with elements")
	assert.Equal(t, 1, value.(int), "First value should be 1, instead it is %d", value.(int))
	value, err = q.Pop()
	assert.Nil(t, err, "Error should be nil when poping elements from queue with elements")
	assert.Equal(t, 3, value.(int), "First value should be 3, instead it is %d", value.(int))
	value, err = q.Pop()
	assert.Nil(t, err, "Error should be nil when poping elements from queue with elements")
	assert.Equal(t, 7, value.(int), "First value should be 7, instead it is %d", value.(int))
	value, err = q.Pop()
	assert.NotNil(t, err, "Pop() on empty queue should return an error")
}

type NewStruct struct {
	value int
}

func TestQueueWithWrongStruct(t *testing.T) {
	q := OrderedQueue{}
	var value NewStruct
	assert.NotNil(t, q.Push(value), "Pushing a struct that do not implement the \"QueueElement\" should return an error")
}

func TestQueueWithStrings(t *testing.T) {
	q := OrderedQueue{}
	assert.Nil(t, q.Push("asdf"), "Pushing correct values should not return an error")
	assert.Nil(t, q.Push("zxcv"), "Pushing correct values should not return an error")
	assert.Nil(t, q.Push("qwer"), "Pushing correct values should not return an error")
	assert.Equal(t, 3, q.Size(), "The queue should contain 3 elements")
	value, err := q.Pop()
	assert.Nil(t, err, "Error should be nil when poping elements from queue with elements")
	assert.Equal(t, "asdf", value.(string), "First value should be \"asdf\", instead it is %s", value.(string))
	value, err = q.Pop()
	assert.Nil(t, err, "Error should be nil when poping elements from queue with elements")
	assert.Equal(t, "qwer", value.(string), "First value should be \"qwer\", instead it is %s", value.(string))
	value, err = q.Pop()
	assert.Nil(t, err, "Error should be nil when poping elements from queue with elements")
	assert.Equal(t, "zxcv", value.(string), "First value should be \"zxcv\", instead it is %s", value.(string))
	value, err = q.Pop()
	assert.NotNil(t, err, "Pop() on empty queue should return an error")
}

type MyData struct {
	A int
	B string
}

func (this MyData) LessThan(other interface{}) bool {
	return this.A < other.(MyData).A
}
func (this MyData) Equals(other interface{}) bool {
	return this.A == other.(MyData).A
}

func TestQueueWithCustomStruct(t *testing.T) {
	q := OrderedQueue{}
	assert.Nil(t, q.Push(MyData{3, "asdf"}), "Pushing correct values should not return an error")
	assert.Nil(t, q.Push(MyData{7, "qwer"}), "Pushing correct values should not return an error")
	assert.Nil(t, q.Push(MyData{1, "zxcv"}), "Pushing correct values should not return an error")
	assert.Equal(t, 3, q.Size(), "The queue should contain 3 elements")
	value, err := q.Pop()
	assert.Nil(t, err, "Error should be nil when poping elements from queue with elements")
	assert.Equal(t, 1, value.(MyData).A, "Value A should be 1, instead it is %d", value.(MyData).A)
	assert.Equal(t, "zxcv", value.(MyData).B, "Value A should be zxcv, instead it is %s", value.(MyData).B)
	value, err = q.Pop()
	assert.Nil(t, err, "Error should be nil when poping elements from queue with elements")
	assert.Equal(t, 3, value.(MyData).A, "Value A should be 3, instead it is %d", value.(MyData).A)
	assert.Equal(t, "asdf", value.(MyData).B, "Value A should be asdf, instead it is %s", value.(MyData).B)
	value, err = q.Pop()
	assert.Nil(t, err, "Error should be nil when poping elements from queue with elements")
	assert.Equal(t, 7, value.(MyData).A, "Value A should be 7, instead it is %d", value.(MyData).A)
	assert.Equal(t, "qwer", value.(MyData).B, "Value A should be qwer, instead it is %s", value.(MyData).B)
	value, err = q.Pop()
	assert.NotNil(t, err, "Pop() on empty queue should return an error")
}

func TestQueueWithWrongValues(t *testing.T) {
	q := OrderedQueue{}
	assert.Nil(t, q.Push(3), "Pushing correct values should not return an error")
	assert.NotNil(t, q.Push("asdf"), "Pushing values of a different type from the original should return an error")
	assert.Equal(t, 1, q.Size(), "The queue should contain 1 element")
}

func TestQueueWithMultipleTypes(t *testing.T) {
	q := OrderedQueue{}
	{
		vector := []int{1, 5, 3}
		for _, element := range vector {
			assert.Nil(t, q.Push(element), "Pushing correct values should not return an error")
		}
		assert.Equal(t, 3, q.Size(), "The queue should contain 3 elements, instead int contains %d elements", q.Size())
		for _, _ = range vector {
			_, err := q.Pop()
			assert.Nil(t, err, "Pop() should not return an error")
		}
		assert.Equal(t, 0, q.Size(), "The queue should contain 0 elements")
	}
	{
		vector := []int8{1, 5, 3}
		for _, element := range vector {
			assert.Nil(t, q.Push(element), "Pushing correct values should not return an error")
		}
		assert.Equal(t, 3, q.Size(), "The queue should contain 3 elements, instead int contains %d elements", q.Size())
		for _, _ = range vector {
			_, err := q.Pop()
			assert.Nil(t, err, "Pop() should not return an error")
		}
		assert.Equal(t, 0, q.Size(), "The queue should contain 0 elements")
	}
	{
		vector := []int16{1, 5, 3}
		for _, element := range vector {
			assert.Nil(t, q.Push(element), "Pushing correct values should not return an error")
		}
		assert.Equal(t, 3, q.Size(), "The queue should contain 3 elements, instead int contains %d elements", q.Size())
		for _, _ = range vector {
			_, err := q.Pop()
			assert.Nil(t, err, "Pop() should not return an error")
		}
		assert.Equal(t, 0, q.Size(), "The queue should contain 0 elements")
	}
	{
		vector := []int32{1, 5, 3}
		for _, element := range vector {
			assert.Nil(t, q.Push(element), "Pushing correct values should not return an error")
		}
		assert.Equal(t, 3, q.Size(), "The queue should contain 3 elements, instead int contains %d elements", q.Size())
		for _, _ = range vector {
			_, err := q.Pop()
			assert.Nil(t, err, "Pop() should not return an error")
		}
		assert.Equal(t, 0, q.Size(), "The queue should contain 0 elements")
	}
	{
		vector := []int64{1, 5, 3}
		for _, element := range vector {
			assert.Nil(t, q.Push(element), "Pushing correct values should not return an error")
		}
		assert.Equal(t, 3, q.Size(), "The queue should contain 3 elements, instead int contains %d elements", q.Size())
		for _, _ = range vector {
			_, err := q.Pop()
			assert.Nil(t, err, "Pop() should not return an error")
		}
		assert.Equal(t, 0, q.Size(), "The queue should contain 0 elements")
	}
	{
		vector := []uint{1, 5, 3}
		for _, element := range vector {
			assert.Nil(t, q.Push(element), "Pushing correct values should not return an error")
		}
		assert.Equal(t, 3, q.Size(), "The queue should contain 3 elements, instead int contains %d elements", q.Size())
		for _, _ = range vector {
			_, err := q.Pop()
			assert.Nil(t, err, "Pop() should not return an error")
		}
		assert.Equal(t, 0, q.Size(), "The queue should contain 0 elements")
	}
	{
		vector := []uint8{1, 5, 3}
		for _, element := range vector {
			assert.Nil(t, q.Push(element), "Pushing correct values should not return an error")
		}
		assert.Equal(t, 3, q.Size(), "The queue should contain 3 elements, instead int contains %d elements", q.Size())
		for _, _ = range vector {
			_, err := q.Pop()
			assert.Nil(t, err, "Pop() should not return an error")
		}
		assert.Equal(t, 0, q.Size(), "The queue should contain 0 elements")
	}
	{
		vector := []uint16{1, 5, 3}
		for _, element := range vector {
			assert.Nil(t, q.Push(element), "Pushing correct values should not return an error")
		}
		assert.Equal(t, 3, q.Size(), "The queue should contain 3 elements, instead int contains %d elements", q.Size())
		for _, _ = range vector {
			_, err := q.Pop()
			assert.Nil(t, err, "Pop() should not return an error")
		}
		assert.Equal(t, 0, q.Size(), "The queue should contain 0 elements")
	}
	{
		vector := []uint32{1, 5, 3}
		for _, element := range vector {
			assert.Nil(t, q.Push(element), "Pushing correct values should not return an error")
		}
		assert.Equal(t, 3, q.Size(), "The queue should contain 3 elements, instead int contains %d elements", q.Size())
		for _, _ = range vector {
			_, err := q.Pop()
			assert.Nil(t, err, "Pop() should not return an error")
		}
		assert.Equal(t, 0, q.Size(), "The queue should contain 0 elements")
	}
	{
		vector := []uint64{1, 5, 3}
		for _, element := range vector {
			assert.Nil(t, q.Push(element), "Pushing correct values should not return an error")
		}
		assert.Equal(t, 3, q.Size(), "The queue should contain 3 elements, instead int contains %d elements", q.Size())
		for _, _ = range vector {
			_, err := q.Pop()
			assert.Nil(t, err, "Pop() should not return an error")
		}
		assert.Equal(t, 0, q.Size(), "The queue should contain 0 elements")
	}
	{
		vector := []float32{1.0, 5.0, 3.0}
		for _, element := range vector {
			assert.Nil(t, q.Push(element), "Pushing correct values should not return an error")
		}
		assert.Equal(t, 3, q.Size(), "The queue should contain 3 elements, instead int contains %d elements", q.Size())
		for _, _ = range vector {
			_, err := q.Pop()
			assert.Nil(t, err, "Pop() should not return an error")
		}
		assert.Equal(t, 0, q.Size(), "The queue should contain 0 elements")
	}
	{
		vector := []float64{1.0, 5.0, 3.0}
		for _, element := range vector {
			assert.Nil(t, q.Push(element), "Pushing correct values should not return an error")
		}
		assert.Equal(t, 3, q.Size(), "The queue should contain 3 elements, instead int contains %d elements", q.Size())
		for _, _ = range vector {
			_, err := q.Pop()
			assert.Nil(t, err, "Pop() should not return an error")
		}
		assert.Equal(t, 0, q.Size(), "The queue should contain 0 elements")
	}
}

func TestClearAllElements(t *testing.T) {
	q := OrderedQueue{}
	vector := []int{1, 5, 3}
	for _, element := range vector {
		assert.Nil(t, q.Push(element), "Pushing correct values should not return an error")
	}
	assert.Equal(t, 3, q.Size(), "The queue should contain 3 elements, instead int contains %d elements", q.Size())
	q.ClearAllElements()
	assert.Equal(t, 0, q.Size(), "The queue should contain 0 elements, instead int contains %d elements", q.Size())
}

func TestRemoveElements(t *testing.T) {
	q := OrderedQueue{}
	vector := []int{1, 5, 3}
	for _, element := range vector {
		assert.Nil(t, q.Push(element), "Pushing correct values should not return an error")
	}
	assert.Equal(t, 3, q.Size(), "The queue should contain 3 elements, instead int contains %d elements", q.Size())
	removed, err := q.RemoveElement(3)
	assert.True(t, removed)
	assert.Nil(t, err)
	assert.Equal(t, 2, q.Size(), "The queue should contain 2 elements, instead int contains %d elements", q.Size())
	removed, err = q.RemoveElement(5)
	assert.True(t, removed)
	assert.Nil(t, err)
	assert.Equal(t, 1, q.Size(), "The queue should contain 1 elements, instead int contains %d elements", q.Size())
	removed, err = q.RemoveElement(1)
	assert.True(t, removed)
	assert.Nil(t, err)
	assert.Equal(t, 0, q.Size(), "The queue should contain 0 elements, instead int contains %d elements", q.Size())
	removed, err = q.RemoveElement(1)
	assert.NotNil(t, err)
}

func TestRemoveRandomElements(t *testing.T) {
	q := OrderedQueue{}
	for i := 0; i < 10; i++ {
		vector := []int{0, 1, 2, 3, 4, 5}
		for j := 0; j < 6; j++ {
			for _, element := range vector {
				assert.Nil(t, q.Push(element), "Pushing correct values should not return an error")
			}
			indexToRemove := 0
			if q.Size() > 1 {
				indexToRemove = rand.Intn(len(vector) - 1)
			}
			elementToRemove := vector[indexToRemove]
			copy(vector[indexToRemove:], vector[indexToRemove+1:])
			vector = vector[:len(vector)-1]
			removed, err := q.RemoveElement(elementToRemove)
			assert.True(t, removed)
			assert.Nil(t, err)
			assert.Equal(t, len(vector), q.Size(), "The queue should contain "+strconv.Itoa(len(vector))+" elements, instead int contains %d elements", q.Size())
			for index := range vector {
				element, err := q.Pop()
				assert.Equal(t, vector[index], element.(int))
				assert.Nil(t, err)
			}
		}
	}
}

func TestRemoveNonExistentElement(t *testing.T) {
	q := OrderedQueue{}
	vector := []int{0, 1, 3, 4, 5}
	for _, element := range vector {
		assert.Nil(t, q.Push(element), "Pushing correct values should not return an error")
	}
	removed, err := q.RemoveElement(2)
	assert.False(t, removed)
	assert.NotNil(t, err)
}

func TestRemoveDifferentType(t *testing.T) {
	q := OrderedQueue{}
	q.Push(1)
	removed, err := q.RemoveElement("asdf")
	assert.False(t, removed)
	assert.NotNil(t, err)
}

package ordered_queue

import (
	"errors"
	"reflect"
)

type QueueElement interface {
	LessThan(interface{}) bool
	Equals(interface{}) bool
}

type OrderedQueue struct {
	elements []interface{}
}

func (q *OrderedQueue) Push(newElement interface{}) (err error) {
	if q.Size() > 0 {
		if reflect.TypeOf(q.elements[0]) != reflect.TypeOf(newElement) {
			return errors.New("Trying to push an element of a different type from the original")
		}
	} else if !isElementValid(newElement) {
		return errors.New("Element's type not valid")
	}
	return q.insertElement(newElement)
}

func (q *OrderedQueue) Pop() (element interface{}, err error) {
	if q.Size() > 0 {
		element = q.elements[0]
		q.elements = q.elements[1:]
	} else {
		err = errors.New("Queue already empty")
	}
	return element, err
}

func (q *OrderedQueue) Size() int {
	return len(q.elements)
}

func (q *OrderedQueue) insertElement(newElement interface{}) error {
	for index, value := range q.elements {
		isLower, err := isElementLowerThan(newElement, value)
		if isLower {
			q.elements = append(q.elements[:index], append([]interface{}{newElement}, q.elements[index:]...)...)
			return nil
		} else if err != nil {
			return err
		}
	}
	q.elements = append(q.elements, newElement)
	return nil
}

func (q *OrderedQueue) ClearAllElements() {
	q.elements = q.elements[:0]
}

func (q *OrderedQueue) RemoveElement(elementToRemove interface{}) (bool, error) {
	if len(q.elements) == 0 {
		return false, errors.New("Empty queue cannot remove elements")
	} else {
		if reflect.TypeOf(q.elements[0]) != reflect.TypeOf(elementToRemove) {
			return false, errors.New("Trying to remove an element of a different type from queue elements")
		}
	}
	index := q.findElement(elementToRemove)
	if index != -1 {
		copy(q.elements[index:], q.elements[index+1:])
		q.elements = q.elements[:len(q.elements)-1]
		return true, nil
	}
	return false, errors.New("Element not found")
}

func (q OrderedQueue) findElement(elementToRemove interface{}) int {
	lowerIndex := 0
	higherIndex := len(q.elements)
	for {
		indexToCheck := lowerIndex + (higherIndex-lowerIndex)/2
		equal, _ := areElementsEqual(q.elements[indexToCheck], elementToRemove)
		if equal {
			return indexToCheck
		}
		lower, _ := isElementLowerThan(elementToRemove, q.elements[indexToCheck])
		if lower {
			higherIndex = indexToCheck - 1
		} else {
			lowerIndex = indexToCheck + 1
		}
		if lowerIndex > higherIndex {
			return -1
		}
	}
}

func (q OrderedQueue) GetCurrentElements() []interface{} {
	return q.elements
}

func areElementsEqual(element1 interface{}, element2 interface{}) (bool, error) {
	if reflect.TypeOf(element1) != reflect.TypeOf(element2) {
		return false, errors.New("Element types are different")
	}
	switch element1Casted := element1.(type) {
	case int:
		if element1Casted == element2.(int) {
			return true, nil
		}
		return false, nil
	case int8:
		if element1Casted == element2.(int8) {
			return true, nil
		}
		return false, nil
	case int16:
		if element1Casted == element2.(int16) {
			return true, nil
		}
		return false, nil
	case int32:
		if element1Casted == element2.(int32) {
			return true, nil
		}
		return false, nil
	case int64:
		if element1Casted == element2.(int64) {
			return true, nil
		}
		return false, nil
	case uint:
		if element1Casted == element2.(uint) {
			return true, nil
		}
		return false, nil
	case uint8:
		if element1Casted == element2.(uint8) {
			return true, nil
		}
		return false, nil
	case uint16:
		if element1Casted == element2.(uint16) {
			return true, nil
		}
		return false, nil
	case uint32:
		if element1Casted == element2.(uint32) {
			return true, nil
		}
		return false, nil
	case uint64:
		if element1Casted == element2.(uint64) {
			return true, nil
		}
		return false, nil
	case float32:
		if element1Casted == element2.(float32) {
			return true, nil
		}
		return false, nil
	case float64:
		if element1Casted == element2.(float64) {
			return true, nil
		}
		return false, nil
	case string:
		if element1Casted == element2.(string) {
			return true, nil
		}
		return false, nil
	case QueueElement:
		if element1Casted.Equals(element2.(QueueElement)) {
			return true, nil
		}
		return false, nil
	}
	return false, errors.New("Trying to push an element that cannot be compared or does not implement the QueueElement interface")
}

func isElementLowerThan(element1 interface{}, element2 interface{}) (bool, error) {
	if reflect.TypeOf(element1) != reflect.TypeOf(element2) {
		return false, errors.New("Element types are different")
	}
	switch element1Casted := element1.(type) {
	case int:
		if element1Casted < element2.(int) {
			return true, nil
		}
		return false, nil
	case int8:
		if element1Casted < element2.(int8) {
			return true, nil
		}
		return false, nil
	case int16:
		if element1Casted < element2.(int16) {
			return true, nil
		}
		return false, nil
	case int32:
		if element1Casted < element2.(int32) {
			return true, nil
		}
		return false, nil
	case int64:
		if element1Casted < element2.(int64) {
			return true, nil
		}
		return false, nil
	case uint:
		if element1Casted < element2.(uint) {
			return true, nil
		}
		return false, nil
	case uint8:
		if element1Casted < element2.(uint8) {
			return true, nil
		}
		return false, nil
	case uint16:
		if element1Casted < element2.(uint16) {
			return true, nil
		}
		return false, nil
	case uint32:
		if element1Casted < element2.(uint32) {
			return true, nil
		}
		return false, nil
	case uint64:
		if element1Casted < element2.(uint64) {
			return true, nil
		}
		return false, nil
	case float32:
		if element1Casted < element2.(float32) {
			return true, nil
		}
		return false, nil
	case float64:
		if element1Casted < element2.(float64) {
			return true, nil
		}
		return false, nil
	case string:
		if element1Casted < element2.(string) {
			return true, nil
		}
		return false, nil
	case QueueElement:
		if element1Casted.LessThan(element2.(QueueElement)) {
			return true, nil
		}
		return false, nil
	}
	return false, errors.New("Trying to push an element that cannot be compared or does not implement the QueueElement interface")
}

func isElementValid(element interface{}) bool {
	switch element.(type) {
	case int:
		return true
	case int8:
		return true
	case int16:
		return true
	case int32:
		return true
	case int64:
		return true
	case uint:
		return true
	case uint8:
		return true
	case uint16:
		return true
	case uint32:
		return true
	case uint64:
		return true
	case float32:
		return true
	case float64:
		return true
	case string:
		return true
	case QueueElement:
		return true
	}
	return false
}

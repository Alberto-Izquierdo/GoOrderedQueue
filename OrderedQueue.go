package ordered_queue

import (
	"errors"
	"reflect"
)

type QueueElement interface {
	LessThan(interface{}) bool
}

type OrderedQueue struct {
	elements []interface{}
}

func (q *OrderedQueue) Push(newElement interface{}) (err error) {
	if q.Size() > 0 {
		if reflect.TypeOf(q.elements[0]) != reflect.TypeOf(newElement) {
			return errors.New("Trying to push an element of a different type from the original")
		}
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
	switch newElementCasted := newElement.(type) {
	case int:
		for index, value := range q.elements {
			if newElementCasted < value.(int) {
				q.elements = append(q.elements[:index], append([]interface{}{newElement}, q.elements[index:]...)...)
				return nil
			}
		}
		q.elements = append(q.elements, newElementCasted)
	case int8:
		for index, value := range q.elements {
			if newElementCasted < value.(int8) {
				q.elements = append(q.elements[:index], append([]interface{}{newElement}, q.elements[index:]...)...)
				return nil
			}
		}
		q.elements = append(q.elements, newElementCasted)
	case int16:
		for index, value := range q.elements {
			if newElementCasted < value.(int16) {
				q.elements = append(q.elements[:index], append([]interface{}{newElement}, q.elements[index:]...)...)
				return nil
			}
		}
		q.elements = append(q.elements, newElementCasted)
	case int32:
		for index, value := range q.elements {
			if newElementCasted < value.(int32) {
				q.elements = append(q.elements[:index], append([]interface{}{newElement}, q.elements[index:]...)...)
				return nil
			}
		}
		q.elements = append(q.elements, newElementCasted)
	case int64:
		for index, value := range q.elements {
			if newElementCasted < value.(int64) {
				q.elements = append(q.elements[:index], append([]interface{}{newElement}, q.elements[index:]...)...)
				return nil
			}
		}
		q.elements = append(q.elements, newElementCasted)
	case uint:
		for index, value := range q.elements {
			if newElementCasted < value.(uint) {
				q.elements = append(q.elements[:index], append([]interface{}{newElement}, q.elements[index:]...)...)
				return nil
			}
		}
		q.elements = append(q.elements, newElementCasted)
	case uint8:
		for index, value := range q.elements {
			if newElementCasted < value.(uint8) {
				q.elements = append(q.elements[:index], append([]interface{}{newElement}, q.elements[index:]...)...)
				return nil
			}
		}
		q.elements = append(q.elements, newElementCasted)
	case uint16:
		for index, value := range q.elements {
			if newElementCasted < value.(uint16) {
				q.elements = append(q.elements[:index], append([]interface{}{newElement}, q.elements[index:]...)...)
				return nil
			}
		}
		q.elements = append(q.elements, newElementCasted)
	case uint32:
		for index, value := range q.elements {
			if newElementCasted < value.(uint32) {
				q.elements = append(q.elements[:index], append([]interface{}{newElement}, q.elements[index:]...)...)
				return nil
			}
		}
		q.elements = append(q.elements, newElementCasted)
	case uint64:
		for index, value := range q.elements {
			if newElementCasted < value.(uint64) {
				q.elements = append(q.elements[:index], append([]interface{}{newElement}, q.elements[index:]...)...)
				return nil
			}
		}
		q.elements = append(q.elements, newElementCasted)
	case float32:
		for index, value := range q.elements {
			if newElementCasted < value.(float32) {
				q.elements = append(q.elements[:index], append([]interface{}{newElement}, q.elements[index:]...)...)
				return nil
			}
		}
		q.elements = append(q.elements, newElementCasted)
	case float64:
		for index, value := range q.elements {
			if newElementCasted < value.(float64) {
				q.elements = append(q.elements[:index], append([]interface{}{newElement}, q.elements[index:]...)...)
				return nil
			}
		}
		q.elements = append(q.elements, newElementCasted)
	case string:
		for index, value := range q.elements {
			if newElementCasted < value.(string) {
				q.elements = append(q.elements[:index], append([]interface{}{newElement}, q.elements[index:]...)...)
				return nil
			}
		}
		q.elements = append(q.elements, newElementCasted)
	case QueueElement:
		for index, value := range q.elements {
			if newElementCasted.LessThan(value.(QueueElement)) {
				q.elements = append(q.elements[:index], append([]interface{}{newElementCasted}, q.elements[index:]...)...)
				return nil
			}
		}
		q.elements = append(q.elements, newElementCasted)
	default:
		return errors.New("Trying to push an element that cannot be compared or does not implement the QueueElement interface")
	}
	return nil
}

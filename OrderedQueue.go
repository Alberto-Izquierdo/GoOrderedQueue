package ordered_queue

import "errors"

type QueueElement interface {
	LessThan(QueueElement) bool
}

type OrderedQueue struct {
	elements []QueueElement
}

func (q *OrderedQueue) Push(newElement QueueElement) {
	for index, value := range q.elements {
		if newElement.LessThan(value) {
			q.elements = append(q.elements[:index], append([]QueueElement{newElement}, q.elements[index:]...)...)
			return
		}
	}
	q.elements = append(q.elements, newElement)
}

func (q *OrderedQueue) Pop() (element QueueElement, err error) {
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

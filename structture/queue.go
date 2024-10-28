package structture

// Queue 队列
type Queue struct {
	elements []interface{}
}

func (q *Queue) Push(val interface{}) {
	q.elements = append(q.elements, val)
}

func (q *Queue) Pop() (interface{}, bool) {
	if len(q.elements) == 0 {
		return nil, false
	}

	v := q.elements[0]
	q.elements = q.elements[1:]
	return v, true
}

func (q *Queue) Front() (interface{}, bool) {
	if len(q.elements) == 0 {
		return nil, false
	}

	return q.elements[0], true
}

func (q *Queue) Back() (interface{}, bool) {
	if len(q.elements) == 0 {
		return nil, false
	}

	return q.elements[len(q.elements)-1], true
}

func (q *Queue) IsEmpty() bool {
	return len(q.elements) == 0
}

func (q *Queue) Size() int {
	return len(q.elements)
}

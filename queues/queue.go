package queues

type Queue struct {
	first *node
	last  *node
}

func (q *Queue) IsEmpty() bool {
	return q.first == nil
}

func (q *Queue) Enqueue(item interface{}) {
	oldLast := q.last
	q.last = &node{item: item}

	if q.IsEmpty() {
		q.first = q.last
	} else {
		oldLast.next = q.last
	}
}

func (q *Queue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}

	oldFirst := q.first
	q.first = oldFirst.next

	return oldFirst.item
}

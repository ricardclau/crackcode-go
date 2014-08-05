package queues

type QueueLinkedList struct {
	first *node
	last  *node
}

func (q *QueueLinkedList) IsEmpty() bool {
	return q.first == nil
}

func (q *QueueLinkedList) Enqueue(item interface{}) {
	oldLast := q.last
	q.last = &node{item: item}

	if q.IsEmpty() {
		q.first = q.last
	} else {
		oldLast.next = q.last
	}
}

func (q *QueueLinkedList) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}

	oldFirst := q.first
	q.first = oldFirst.next

	return oldFirst.item
}

package linkedlists

func deleteDupes(list *SingleLinkedList) {
	if list.top == nil {
		return
	}

	currentNode := list.top
	for currentNode != nil {
		item := currentNode.item
		elimination := currentNode

		for elimination.next != nil {
			if elimination.next.item == item {
				elimination.next = elimination.next.next
				continue
			}

			elimination = elimination.next
		}

		currentNode = currentNode.next
	}
}

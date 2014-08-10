package linkedlists

func deleteNode(node *singlelinkednode) bool {
	if node == nil || node.next == nil {
		return false
	}

	tmp := node.next
	node.item = tmp.item
	node.next = tmp.next

	return true
}

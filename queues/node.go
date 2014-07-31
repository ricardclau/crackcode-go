// The most similar to a Generic that exists in Go
package queues

type node struct {
	item interface{}
	next *node
}

package single

// Node represents a single element in a singly linked list.
// Each Node contains some data and a pointer to the next Node in the list.
type Node struct {
	data int   // The data stored in the node
	next *Node // Pointer to the next node in the list
}

// NewNode creates a new Node with default zero values.
// Returns a pointer to the newly created Node.
func NewNode() *Node {
	return &Node{}
}

// NewNodeWithData creates a new Node with the given data and
// the next pointer set to nil.
// data: The data to be stored in the Node.
// Returns a pointer to the newly created Node.
func NewNodeWithData(data int) *Node {
	return &Node{
		data: data,
		next: nil,
	}
}

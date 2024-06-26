package linkedlist

// Node represents a single element in a singly linked list.
// Each Node contains some data and a pointer to the next Node in the list.
type ListNode struct {
	data int       // The data stored in the node
	prev *ListNode // Pointer to the previous node in the list
	next *ListNode // Pointer to the next node in the list
}

// NewEmptyNode creates a new empty Node with default values.
// Returns a pointer to the newly created Node.
func NewEmptyListNode() *ListNode {
	return &ListNode{}
}

// NewNodeWithData creates a new Node with the given data and
// the next pointer set to nil.
// data: The data to be stored in the Node.
// Returns a pointer to the newly created Node.
func NewListNodeWithData(data int) *ListNode {
	return &ListNode{
		data: data,
		prev: nil,
		next: nil,
	}
}

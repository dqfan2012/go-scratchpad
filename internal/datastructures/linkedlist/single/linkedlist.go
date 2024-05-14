package single

// List interface defines the operations for a singly linked list.
type List interface {
	ClearList()
	DeleteAtPosition(int) (int, bool)
	DeleteHead() (int, bool)
	DeleteTail() (int, bool)
	IsEmpty() bool
	IsValuePresent(int) bool
	InsertAtPosition(int, int)
	InsertHead(int)
	InsertTail(int)
	SetHeadIfEmptyOrInvalidPosition(*Node, int) bool
	Size() int
}

// LinkedList represents a singly linked list.
type LinkedList struct {
	head   *Node
	length int
}

// NewEmptyLinkedList creates an empty linked list.
func NewEmptyLinkedList() LinkedList {
	return LinkedList{
		head:   nil,
		length: 0,
	}
}

// NewLinkedListWithHead creates a linked list with an initial head node.
func NewLinkedListWithHead(data int) LinkedList {
	newNode := NewNodeWithData(data)
	return LinkedList{
		head:   newNode,
		length: 1,
	}
}

// CheckForEmptyListOrInvalidPosition checks to see if the list is clear or the
// position is invalid
func (list *LinkedList) CheckForEmptyListOrInvalidPosition(position int) bool {
	return list.IsEmpty() || position < 0 || position >= list.length
}

// ClearList clears the linked list.
func (list *LinkedList) ClearList() {
	list.head = nil
	list.length = 0
}

// DeleteAtPosition deletes the node at the specified position.
func (list *LinkedList) DeleteAtPosition(position int) (int, bool) {
	if list.CheckForEmptyListOrInvalidPosition(position) {
		return 0, false
	}

	var value int
	if position == 0 {
		value = list.head.data
		list.head = list.head.next
		list.length--
		return value, true
	}

	temp := list.head
	for i := 0; i < position-1; i++ {
		temp = temp.next
	}

	value = temp.next.data
	temp.next = temp.next.next
	list.length--

	return value, true
}

// DeleteHead deletes the head node.
func (list *LinkedList) DeleteHead() (int, bool) {
	if list.IsEmpty() {
		return 0, false
	}

	value := list.head.data
	list.head = list.head.next
	list.length--

	return value, true
}

// DeleteTail deletes the tail node.
func (list *LinkedList) DeleteTail() (int, bool) {
	if list.IsEmpty() {
		return 0, false
	}

	if list.head.next == nil {
		value := list.head.data
		list.head = nil
		list.length--
		return value, true
	}

	temp := list.head
	for temp.next.next != nil {
		temp = temp.next
	}

	value := temp.next.data
	temp.next = nil
	list.length--

	return value, true
}

// IsEmpty checks if the list is empty.
func (list *LinkedList) IsEmpty() bool {
	return list.length == 0
}

// InsertAtPosition inserts a node at the specified position.
func (list *LinkedList) InsertAtPosition(data int, position int) {
	newNode := NewNodeWithData(data)

	if list.SetHeadIfEmptyOrInvalidPosition(newNode, position) {
		return
	}

	temp := list.head
	for i := 0; i < position-1; i++ {
		temp = temp.next
	}

	newNode.next = temp.next
	temp.next = newNode
	list.length++
}

// InsertHead inserts a node at the head of the list.
func (list *LinkedList) InsertHead(data int) {
	newNode := NewNodeWithData(data)

	if list.SetHeadIfEmptyOrInvalidPosition(newNode, 0) {
		return
	}

	newNode.next = list.head
	list.head = newNode
	list.length++
}

// InsertTail inserts a node at the tail of the list.
func (list *LinkedList) InsertTail(data int) {
	newNode := NewNodeWithData(data)

	if list.SetHeadIfEmptyOrInvalidPosition(newNode, 0) {
		list.head = newNode
		list.length++
		return
	}

	temp := list.head
	for temp.next != nil {
		temp = temp.next
	}

	temp.next = newNode
	list.length++
}

// IsValuePresent checks if a value is present in the list.
func (list *LinkedList) IsValuePresent(data int) bool {
	if !list.IsEmpty() {
		temp := list.head

		for temp != nil {
			if temp.data == data {
				return true
			}

			temp = temp.next
		}
	}

	return false
}

// SetHeadIfEmptyOrInvalidPosition sets the head if the list is empty or position is invalid.
func (list *LinkedList) SetHeadIfEmptyOrInvalidPosition(node *Node, position int) bool {
	if list.CheckForEmptyListOrInvalidPosition(position) {
		list.head = node
		list.length++
		return true
	}

	return false
}

// Size returns the size of the list.
func (list *LinkedList) Size() int {
	return list.length
}

package single

type List interface {
	ClearList()
	DeleteAtPosition(int) (int, bool)
	DeleteHead() (int, bool)
	DeleteTail() (int, bool)
	IsEmpty() bool
	IsPresent(int) bool
	InsertAtPosition(int, int)
	InsertHead(int)
	InsertTail(int)
	Size() int
}

type LinkedList struct {
	head   *Node
	length int
}

func NewEmptyLinkedList() LinkedList {
	return LinkedList{
		head:   nil,
		length: 0,
	}
}

func NewLinkedListWithHead(data int) LinkedList {
	return LinkedList{
		head: &Node{
			data: data,
			next: nil,
		},
		length: 1,
	}
}

func (list *LinkedList) ClearList() {
	list.head = nil
	list.length = 0
}

func (list *LinkedList) DeleteAtPosition(position int) (int, bool) {
	if !list.IsEmpty() {
		itemBeforeRemoval := list.head

		// Grab the second to last element
		for i := 0; i < position-1; i++ {
			itemBeforeRemoval = itemBeforeRemoval.next
		}

		// Grab the last element.
		itemToRemove := itemBeforeRemoval.next

		// Get the value of the removed element
		value := itemToRemove.data

		// Replace the "next" value for the item right before
		// the item we're removing
		itemBeforeRemoval.next = itemToRemove.next
		list.length--

		return value, true
	}

	return 0, false
}

func (list *LinkedList) DeleteHead() (int, bool) {
	if !list.IsEmpty() {
		value := list.head.data
		list.head = list.head.next
		list.length--

		return value, true
	}

	return 0, false
}

func (list *LinkedList) DeleteTail() (int, bool) {
	if !list.IsEmpty() {
		nextToLast := list.head

		// Grab the second to last element
		for i := 0; i < list.length-1; i++ {
			nextToLast = nextToLast.next
		}

		// Grab the last element.
		last := nextToLast.next

		// Get the value of the last element
		value := last.data

		// Remove the last element
		nextToLast.next = nil
		list.length--

		return value, true
	}

	return 0, false
}

func (list *LinkedList) IsEmpty() bool {
	return list.length == 0
}

func (list *LinkedList) IsPresent(data int) bool {
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

func (list *LinkedList) InsertAtPosition(data int, position int) {
	newNode := NewNodeWithData(data)

	if list.IsEmpty() {
		list.head = newNode
	} else {
		temp := list.head

		for i := 0; i < position-1; i++ {
			temp = temp.next
		}

		newNode.next = temp.next
		temp.next = newNode
	}

	list.length++
}

func (list *LinkedList) InsertHead(data int) {
	newNode := NewNodeWithData(data)

	if list.IsEmpty() {
		list.head = newNode
	} else {
		newNode.next = list.head
		list.head = newNode
	}

	list.length++
}

func (list *LinkedList) InsertTail(data int) {
	newNode := NewNodeWithData(data)

	if list.IsEmpty() {
		list.head = newNode
	} else {
		temp := list.head

		for temp.next != nil {
			temp = temp.next
		}

		temp.next = newNode
	}

	list.length++
}

func (list *LinkedList) Size() int {
	return list.length
}

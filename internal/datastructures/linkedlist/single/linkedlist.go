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
	SetHeadIfEmptyOrInvalidPosition(*Node, int) bool
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
	newNode := NewNodeWithData(data)
	return LinkedList{
		head:   newNode,
		length: 1,
	}
}

func (list *LinkedList) ClearList() {
	list.head = nil
	list.length = 0
}

func (list *LinkedList) DeleteAtPosition(position int) (int, bool) {
	if list.IsEmpty() || position < 0 || position >= list.length {
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

func (list *LinkedList) DeleteHead() (int, bool) {
	if list.IsEmpty() {
		return 0, false
	}

	value := list.head.data
	list.head = list.head.next
	list.length--

	return value, true
}

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

func (list *LinkedList) IsEmpty() bool {
	return list.length == 0
}

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

func (list *LinkedList) InsertHead(data int) {
	newNode := NewNodeWithData(data)

	if list.SetHeadIfEmptyOrInvalidPosition(newNode, 0) {
		return
	}

	newNode.next = list.head
	list.head = newNode
	list.length++
}

func (list *LinkedList) InsertTail(data int) {
	newNode := NewNodeWithData(data)

	if list.SetHeadIfEmptyOrInvalidPosition(newNode, 0) {
		return
	}

	temp := list.head
	for temp.next != nil {
		temp = temp.next
	}

	temp.next = newNode
	list.length++
}

func (list *LinkedList) SetHeadIfEmptyOrInvalidPosition(node *Node, position int) bool {
	if list.IsEmpty() || position < 0 {
		list.head = node
		list.length++
		return true
	}

	return false
}

func (list *LinkedList) Size() int {
	return list.length
}

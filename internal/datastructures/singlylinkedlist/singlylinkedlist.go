package singlylinkedlist

type List interface {
	IsEmpty() bool
	InsertAtBeginning(int)
	InsertAtEnd(int)
	InsertAtPosition(int, int)
	RemoveFromBeginning()
	RemoveFromEnd()
	RemoveFromPosition(int)
}

type SinglyLinkedList struct {
	head *Node
}

func (list *SinglyLinkedList) IsEmpty() bool {
	return list.head == nil
}

func (list *SinglyLinkedList) InsertAtBeginning(data int) {
	newNode := NewNodeWithData(data)

	if list.IsEmpty() {
		list.head = newNode
	} else {
		//fmt.Println("Before:")
		//fmt.Println(newNode)
		//fmt.Println(list.head)

		newNode.next = list.head
		list.head = newNode

		//fmt.Println("--------------")
		//fmt.Println("After:")
		//fmt.Println(newNode)
		//fmt.Println(list.head)
	}
}

func (list *SinglyLinkedList) InsertAtEnd(data int) {
	newNode := NewNodeWithData(data)

	if list.IsEmpty() {
		list.head = newNode
	} else {
		temp := list.head

		// Loop until we get to the last item in the list
		for temp.next != nil {
			temp = temp.next
		}

		temp.next = newNode
	}
}

func (list *SinglyLinkedList) InsertAtPosition(data int, position int) {
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
}

func (list *SinglyLinkedList) RemoveFromBeginning() {

}

func (list *SinglyLinkedList) RemoveFromEnd() {

}

func (list *SinglyLinkedList) RemoveFromPosition(position int) {

}

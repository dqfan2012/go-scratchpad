package singlylinkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertAtBeginning(t *testing.T) {
	expectedValue := 3
	linkedList := &SinglyLinkedList{}

	linkedList.InsertAtBeginning(expectedValue)

	assert.Equal(t, expectedValue, linkedList.head.data)
}

func TestInsertAtEnd(t *testing.T) {
	value1ToInsert := 3
	value2ToInsert := 5
	expectedValue := 7
	linkedList := &SinglyLinkedList{}

	linkedList.InsertAtBeginning(value1ToInsert)
	linkedList.InsertAtBeginning(value2ToInsert)
	linkedList.InsertAtEnd(expectedValue)

	temp := linkedList.head
	for temp.next != nil {
		temp = temp.next
	}

	assert.Equal(t, expectedValue, temp.data)
}

func TestInsertAtPosition(t *testing.T) {
	expectedValue := 5
	position := 3
	linkedList := &SinglyLinkedList{}

	linkedList.InsertAtBeginning(3)
	linkedList.InsertAtBeginning(6)
	linkedList.InsertAtBeginning(9)
	linkedList.InsertAtBeginning(12)
	linkedList.InsertAtBeginning(1)
	linkedList.InsertAtBeginning(4)
	linkedList.InsertAtPosition(expectedValue, position)

	// Get the item inserted at the specific position
	temp := linkedList.head
	for i := 0; i < position; i++ {
		temp = temp.next
	}

	assert.Equal(t, expectedValue, temp.data)
}

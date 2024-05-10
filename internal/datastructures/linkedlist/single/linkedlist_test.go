package single

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyLinkedList(t *testing.T) {
	LinkedList := NewEmptyLinkedList()

	expectedListLength := 0
	actualListLength := LinkedList.length
	actualHead := LinkedList.head

	assert.Equal(t, expectedListLength, actualListLength)
	assert.True(t, reflect.ValueOf(actualHead).IsNil())
}

func TestNewListWithHead(t *testing.T) {
	expectedValue := 15
	LinkedList := NewLinkedListWithHead(expectedValue)

	expectedListLength := 1
	actualListLength := LinkedList.length
	actualHead := LinkedList.head

	assert.Equal(t, expectedListLength, actualListLength)
	assert.Equal(t, expectedValue, actualHead.data)
}

func TestInsertAtBeginning(t *testing.T) {
	expectedValue := 3
	linkedList := &LinkedList{}

	linkedList.InsertHead(expectedValue)

	assert.Equal(t, expectedValue, linkedList.head.data)
}

func TestInsertAtEnd(t *testing.T) {
	value1ToInsert := 3
	value2ToInsert := 5
	expectedValue := 7
	linkedList := &LinkedList{}

	linkedList.InsertHead(value1ToInsert)
	linkedList.InsertHead(value2ToInsert)
	linkedList.InsertTail(expectedValue)

	temp := linkedList.head
	for temp.next != nil {
		temp = temp.next
	}

	assert.Equal(t, expectedValue, temp.data)
}

func TestInsertAtPosition(t *testing.T) {
	expectedValue := 5
	position := 3
	linkedList := &LinkedList{}

	linkedList.InsertHead(3)
	linkedList.InsertHead(6)
	linkedList.InsertHead(9)
	linkedList.InsertHead(12)
	linkedList.InsertHead(1)
	linkedList.InsertHead(4)
	linkedList.InsertAtPosition(expectedValue, position)

	// Get the item inserted at the specific position
	temp := linkedList.head
	for i := 0; i < position; i++ {
		temp = temp.next
	}

	assert.Equal(t, expectedValue, temp.data)
}

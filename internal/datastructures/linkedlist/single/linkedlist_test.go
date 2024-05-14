package single

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedListInitialization(t *testing.T) {
	linkedList := NewEmptyLinkedList()
	assert.Equal(t, 0, linkedList.length)
	assert.Nil(t, linkedList.head)

	linkedListWithHead := NewLinkedListWithHead(10)
	assert.Equal(t, 1, linkedListWithHead.length)
	assert.Equal(t, 10, linkedListWithHead.head.data)
}

func TestInsertHead(t *testing.T) {
	linkedList := NewEmptyLinkedList()
	linkedList.InsertHead(10)
	linkedList.InsertHead(20)

	assert.Equal(t, 2, linkedList.length)
	assert.Equal(t, 20, linkedList.head.data)
	assert.Equal(t, 10, linkedList.head.next.data)
}

func TestInsertTail(t *testing.T) {
	linkedList := NewEmptyLinkedList()
	linkedList.InsertTail(10)
	linkedList.InsertTail(20)

	assert.Equal(t, 2, linkedList.length)
	assert.Equal(t, 10, linkedList.head.data)
	assert.Equal(t, 20, linkedList.head.next.data)
}

func TestInsertAtPosition(t *testing.T) {
	linkedList := NewEmptyLinkedList()
	linkedList.InsertHead(10)
	linkedList.InsertHead(20)
	linkedList.InsertAtPosition(15, 1)

	assert.Equal(t, 3, linkedList.length)
	assert.Equal(t, 20, linkedList.head.data)
	assert.Equal(t, 15, linkedList.head.next.data)
	assert.Equal(t, 10, linkedList.head.next.next.data)
}

func TestDeleteHead(t *testing.T) {
	linkedList := NewEmptyLinkedList()
	linkedList.InsertHead(10)
	linkedList.InsertHead(20)
	value, ok := linkedList.DeleteHead()

	assert.True(t, ok)
	assert.Equal(t, 20, value)
	assert.Equal(t, 1, linkedList.length)
	assert.Equal(t, 10, linkedList.head.data)
}

func TestDeleteTail(t *testing.T) {
	linkedList := NewEmptyLinkedList()
	linkedList.InsertHead(10)
	linkedList.InsertHead(20)
	value, ok := linkedList.DeleteTail()

	assert.True(t, ok)
	assert.Equal(t, 10, value)
	assert.Equal(t, 1, linkedList.length)
	assert.Equal(t, 20, linkedList.head.data)
}

func TestDeleteAtPosition(t *testing.T) {
	linkedList := NewEmptyLinkedList()
	linkedList.InsertHead(10)
	linkedList.InsertHead(20)
	linkedList.InsertHead(30)
	value, ok := linkedList.DeleteAtPosition(1)

	assert.True(t, ok)
	assert.Equal(t, 20, value)
	assert.Equal(t, 2, linkedList.length)
	assert.Equal(t, 30, linkedList.head.data)
	assert.Equal(t, 10, linkedList.head.next.data)
}

func TestIsEmpty(t *testing.T) {
	linkedList := NewEmptyLinkedList()
	assert.True(t, linkedList.IsEmpty())

	linkedList.InsertHead(10)
	assert.False(t, linkedList.IsEmpty())
}

func TestIsValuePresent(t *testing.T) {
	linkedList := NewEmptyLinkedList()
	linkedList.InsertHead(10)
	linkedList.InsertHead(20)

	assert.True(t, linkedList.IsValuePresent(10))
	assert.False(t, linkedList.IsValuePresent(30))
}

func TestSize(t *testing.T) {
	linkedList := NewEmptyLinkedList()
	assert.Equal(t, 0, linkedList.Size())

	linkedList.InsertHead(10)
	linkedList.InsertHead(20)
	assert.Equal(t, 2, linkedList.Size())
}

package singlylinkedlist

import "github.com/stretchr/testify/mock"

type MockSinglyLinkedList struct {
	mock.Mock
}

func newMockSinglyLinkedList() *MockSinglyLinkedList { return &MockSinglyLinkedList{} }

func (m *MockSinglyLinkedList) IsEmpty() bool {
	args := m.Called()
	return args.Bool(0)
}

func (m *MockSinglyLinkedList) InsertAtBeginning(data int) {
	m.Called(data)
}

func (m *MockSinglyLinkedList) InsertAtEnd(data int) {
	m.Called(data)
}

func (m *MockSinglyLinkedList) InsertAtPosition(data int, position int) {
	m.Called(data, position)
}

func (m *MockSinglyLinkedList) RemoveFromBeginning() {
	m.Called()
}

func (m *MockSinglyLinkedList) RemoveFromEnd() {
	m.Called()
}

func (m *MockSinglyLinkedList) RemoveFromPosition(position int) {
	m.Called(position)
}

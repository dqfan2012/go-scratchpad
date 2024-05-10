package single

type Node struct {
	data int
	next *Node
}

func NewNode() *Node {
	return &Node{}
}

func NewNodeWithData(data int) *Node {
	return &Node{
		data: data,
		next: nil,
	}
}

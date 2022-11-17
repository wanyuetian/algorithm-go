package graph

type Node[T any] struct {
	Value T
	In    int
	Out   int
	Nexts []*Node[T]
	Edges []*Edge[T]
}

func NewNode[T any](value T) *Node[T] {
	return &Node[T]{
		Value: value,
		In:    0,
		Out:   0,
		Nexts: make([]*Node[T], 0),
		Edges: make([]*Edge[T], 0),
	}
}

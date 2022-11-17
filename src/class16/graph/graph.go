package graph

type Graph[T any] struct {
	Nodes map[int]*Node[T]
	Edges map[*Edge[T]]struct{}
}

func NewGraph[T any]() *Graph[T] {
	return &Graph[T]{
		Nodes: make(map[int]*Node[T]),
		Edges: make(map[*Edge[T]]struct{}),
	}
}

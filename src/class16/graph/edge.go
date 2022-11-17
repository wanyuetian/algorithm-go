package graph

type Edge[T any] struct {
	Weight int
	From   *Node[T]
	To     *Node[T]
}

func NewEdge[T any](weight int, from *Node[T], to *Node[T]) *Edge[T] {
	return &Edge[T]{
		Weight: weight,
		From:   from,
		To:     to,
	}
}

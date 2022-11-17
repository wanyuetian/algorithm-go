package main

import (
	"fmt"
	"math"
	"sort"
)

type DirectedGraphNode struct {
	Label      int
	Neighbours []*DirectedGraphNode
}

type Record struct {
	Node *DirectedGraphNode
	Deep int
}

func topSort(graph []*DirectedGraphNode) []*DirectedGraphNode {
	order := make(map[*DirectedGraphNode]*Record)
	for i := range graph {
		process(graph[i], order)
	}
	records := make([]*Record, 0)
	for node := range order {
		records = append(records, order[node])
	}
	sort.Slice(records, func(i, j int) bool {
		if records[i].Deep >= records[j].Deep {
			return true
		}
		return false
	})
	ans := make([]*DirectedGraphNode, 0)
	for i := range records {
		ans = append(ans, records[i].Node)
	}
	return ans
}

func process(node *DirectedGraphNode, order map[*DirectedGraphNode]*Record) *Record {
	if _, ok := order[node]; ok {
		return order[node]
	}
	follow := 0
	for i := range node.Neighbours {
		follow = int(math.Max(float64(follow), float64(process(node.Neighbours[i], order).Deep)))
	}
	ans := &Record{
		Node: node,
		Deep: follow + 1,
	}
	order[node] = ans
	return ans
}

func main() {
	graph := make([]*DirectedGraphNode, 0)
	n1 := &DirectedGraphNode{
		Label:      1,
		Neighbours: nil,
	}
	n2 := &DirectedGraphNode{
		Label:      2,
		Neighbours: []*DirectedGraphNode{n1},
	}
	n3 := &DirectedGraphNode{
		Label:      3,
		Neighbours: []*DirectedGraphNode{n1, n2},
	}
	n4 := &DirectedGraphNode{
		Label:      4,
		Neighbours: []*DirectedGraphNode{n3},
	}

	n5 := &DirectedGraphNode{
		Label:      5,
		Neighbours: []*DirectedGraphNode{n3, n4},
	}
	graph = []*DirectedGraphNode{n1, n2, n3, n4, n5}
	res := topSort(graph)
	for i := range res {
		fmt.Println(res[i].Label)
	}
}

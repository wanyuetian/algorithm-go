package main

import "fmt"

type DirectedGraphNode struct {
	Label      int
	Neighbours []*DirectedGraphNode
}

// https://www.lintcode.com/problem/127/
func topSort(graph []*DirectedGraphNode) []*DirectedGraphNode {
	ans := make([]*DirectedGraphNode, 0)
	m := make(map[*DirectedGraphNode]int)
	for i := range graph {
		m[graph[i]] = 0
	}
	for i := range graph {
		for j := range graph[i].Neighbours {
			m[graph[i].Neighbours[j]]++
		}
	}
	zeroQueue := make([]*DirectedGraphNode, 0)
	for node := range m {
		if m[node] == 0 {
			zeroQueue = append(zeroQueue, node)
		}
	}
	for len(zeroQueue) > 0 {
		cur := zeroQueue[0]
		zeroQueue = zeroQueue[1:]
		ans = append(ans, cur)
		for i := range cur.Neighbours {
			m[cur.Neighbours[i]]--
			if m[cur.Neighbours[i]] == 0 {
				zeroQueue = append(zeroQueue, cur.Neighbours[i])
			}
		}
	}
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

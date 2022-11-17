package main

import (
	"fmt"
	"kuaishou.com/algorithm-go/src/class16/graph"
)

func dfs[T any](start *graph.Node[T]) {
	if start == nil {
		return
	}
	stack := make([]*graph.Node[T], 0)
	set := make(map[*graph.Node[T]]struct{})
	stack = append(stack, start)
	set[start] = struct{}{}
	fmt.Println(start.Value)

	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		for i := range cur.Nexts {
			if _, ok := set[cur.Nexts[i]]; !ok {
				stack = append(stack, cur)
				stack = append(stack, cur.Nexts[i])
				fmt.Println(cur.Nexts[i].Value)
				break
			}
		}

	}
}

func main() {
	g := graph.NewGraph[string]()
	n1 := graph.NewNode("node1")
	n2 := graph.NewNode("node2")
	n3 := graph.NewNode("node3")
	n4 := graph.NewNode("node4")
	n5 := graph.NewNode("node5")
	g.Nodes[1] = n1
	g.Nodes[2] = n2
	g.Nodes[3] = n3
	g.Nodes[4] = n4
	g.Nodes[5] = n5

	e1_2 := graph.NewEdge(1, n1, n2)
	e1_3 := graph.NewEdge(1, n1, n3)
	n1.Edges = append(n1.Edges, e1_2)
	n1.Out++
	n2.In++
	n1.Edges = append(n1.Edges, e1_3)
	n1.Out++
	n3.In++
	n1.Nexts = append(n1.Nexts, n2)
	n1.Nexts = append(n1.Nexts, n3)
	e2_3 := graph.NewEdge(1, n2, n3)
	n2.Edges = append(n1.Edges, e2_3)
	n2.Out++
	n3.In++
	e2_4 := graph.NewEdge(1, n2, n4)
	n2.Edges = append(n1.Edges, e2_4)
	n2.Out++
	n4.In++
	e2_5 := graph.NewEdge(1, n2, n5)
	n2.Edges = append(n1.Edges, e2_5)
	n2.Out++
	n5.In++
	n2.Nexts = append(n2.Nexts, n3)
	n2.Nexts = append(n2.Nexts, n4)
	n2.Nexts = append(n2.Nexts, n5)
	e3_5 := graph.NewEdge(1, n3, n5)
	n3.Edges = append(n1.Edges, e3_5)
	n3.Out++
	n5.In++
	n3.Nexts = append(n3.Nexts, n5)
	e4_1 := graph.NewEdge(1, n4, n1)
	n4.Edges = append(n1.Edges, e4_1)
	n4.Out++
	n1.In++
	n4.Nexts = append(n4.Nexts, n1)
	g.Edges[e1_2] = struct{}{}
	g.Edges[e1_3] = struct{}{}
	g.Edges[e2_3] = struct{}{}
	g.Edges[e2_4] = struct{}{}
	g.Edges[e2_5] = struct{}{}
	g.Edges[e3_5] = struct{}{}
	g.Edges[e4_1] = struct{}{}

	dfs(g.Nodes[1])
}

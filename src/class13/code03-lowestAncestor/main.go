package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

type Info struct {
	FindA bool
	FindB bool
	Ans   *Node
}

func lowestAncestor1(head *Node, a, b *Node) *Node {
	if head == nil || a == nil || b == nil {
		return nil
	}
	m := make(map[*Node]*Node)
	m[head] = nil
	getParentMap(head, &m)

	nodeSet := make(map[*Node]struct{})
	nodeSet[a] = struct{}{}
	cur := m[a]
	for cur != nil {
		nodeSet[cur] = struct{}{}
		cur = m[cur]
	}
	cur = b
	var ans *Node
	for cur != nil {
		if _, ok := nodeSet[cur]; ok {
			ans = cur
			break
		}
		cur = m[cur]
	}

	return ans
}

func getParentMap(head *Node, m *map[*Node]*Node) {
	if head.Left != nil {
		(*m)[head.Left] = head
		getParentMap(head.Left, m)
	}
	if head.Right != nil {
		(*m)[head.Right] = head
		getParentMap(head.Right, m)
	}
}

func lowestAncestor2(head *Node, a, b *Node) *Node {
	return process(head, a, b).Ans
}

func process(head, a, b *Node) *Info {
	if head == nil {
		return &Info{
			FindA: false,
			FindB: false,
			Ans:   nil,
		}
	}

	leftInfo := process(head.Left, a, b)
	rightInfo := process(head.Right, a, b)
	var findA, findB bool
	var ans *Node
	if leftInfo.Ans != nil {
		ans = leftInfo.Ans
	}
	if rightInfo.Ans != nil {
		ans = rightInfo.Ans
	}
	if head == a {
		findA = true
	}
	if head == b {
		findB = true
	}
	if leftInfo.FindA || rightInfo.FindA {
		findA = true
	}
	if leftInfo.FindB || rightInfo.FindB {
		findB = true
	}
	if ans == nil {
		if findA && findB {
			ans = head
		}
	}

	return &Info{
		FindA: findA,
		FindB: findB,
		Ans:   ans,
	}
}

func pickRandomOne(head *Node) *Node {
	if head == nil {
		return nil
	}
	arr := make([]*Node, 0)
	var pre func(head *Node)
	pre = func(node *Node) {
		if node == nil {
			return
		}
		arr = append(arr, node)
		pre(node.Left)
		pre(node.Right)
	}
	pre(head)
	return arr[rand.Intn(len(arr))]
}

func generateRandomBST(maxLevel, maxValue int) *Node {
	return generate(1, maxLevel, maxValue)
}

func generate(level int, maxLevel int, maxValue int) *Node {
	if level > maxLevel || rand.Float64() < 0.5 {
		return nil
	}
	head := &Node{Val: rand.Intn(maxValue) + 1}
	head.Left = generate(level+1, maxLevel, maxValue)
	head.Right = generate(level+1, maxLevel, maxValue)
	return head
}

func main() {
	rand.Seed(time.Now().UnixNano())
	maxLevel := 10
	maxValue := 100
	testTimes := 100000
	for i := 0; i < testTimes; i++ {
		head := generateRandomBST(maxLevel, maxValue)
		a := pickRandomOne(head)
		b := pickRandomOne(head)
		r1 := lowestAncestor1(head, a, b)
		r2 := lowestAncestor2(head, a, b)
		if r1 != r2 {
			fmt.Println("Oops")
		}
	}
	fmt.Println("Nice")
}

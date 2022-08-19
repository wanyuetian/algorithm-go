package main

type Node struct {
	Val    int
	Left   *Node
	Right  *Node
	Parent *Node
}

func getSuccessorNode(node *Node) *Node {
	if node == nil {
		return nil
	}
	if node.Right != nil {
		target := node.Right
		for target.Left != nil {
			target = target.Left
		}
		return target
	}
	parent := node.Parent
	for parent != nil && parent.Right == node {
		node = node.Parent
		parent = node.Parent
	}
	return parent
}

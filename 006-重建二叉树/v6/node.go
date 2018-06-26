package main

type Node struct {
	left  *Node
	right *Node
	value int
}

func (node *Node) PreOrderTraversal(result []int) []int {
	if node != nil {
		result = append(result, node.value)
	}
	if node.left != nil {
		result = node.left.PreOrderTraversal(result)
	}
	if node.right != nil {
		result = node.right.PreOrderTraversal(result)
	}

	return result
}


func (node *Node) InOrderTraversal(result []int) []int {
	if node.left != nil {
		result = node.left.PreOrderTraversal(result)
	}
	if node != nil {
		result = append(result, node.value)
	}
	if node.right != nil {
		result = node.right.PreOrderTraversal(result)
	}

	return result

}
func (node *Node) BreadthFirstTraversal(result []int) []int {
	queue := []*Node{node}
	var n *Node

	for len(queue) > 0 {
		n = queue[0]
		queue = queue[1:]

		result = append(result, n.value)

		if n.left != nil {
			queue = append(queue, n.left)
		}
		if n.right != nil {
			queue = append(queue, n.right)
		}
	}

	return result
}

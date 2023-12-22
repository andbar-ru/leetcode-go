package leetcode

type Node struct {
	Val       int
	Neighbors []*Node
}

func slice2graph(slice [][]int) *Node {
	if len(slice) == 0 {
		return nil
	}
	nodes := make([]*Node, len(slice))
	for i, vertices := range slice {
		nodes[i] = &Node{
			Val:       i + 1,
			Neighbors: make([]*Node, len(vertices)),
		}
	}
	for i, vertices := range slice {
		node := nodes[i]
		for j, val := range vertices {
			node.Neighbors[j] = nodes[val-1]
		}
	}
	return nodes[0]
}

func graph2slice(node *Node) [][]int {
	if node == nil {
		return [][]int{}
	}
	adjList := make([][]int, 0)
	queue := []*Node{node}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		for len(adjList) < node.Val {
			adjList = append(adjList, nil)
		}
		if adjList[node.Val-1] != nil {
			continue
		}
		for _, neighbor := range node.Neighbors {
			adjList[node.Val-1] = append(adjList[node.Val-1], neighbor.Val)
			if len(adjList) < neighbor.Val || adjList[neighbor.Val-1] == nil {
				queue = append(queue, neighbor)
			}
		}
	}
	return adjList
}

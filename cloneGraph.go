package leetcode

/*
133. Clone Graph

Given a reference of a node in a connected undirected graph.

Return a deep copy (clone) of the graph.

Each node in the graph contains a value (int) and a list (List[Node]) of its neighbors.
```
class Node {
	public int val;
	public List<Node> neighbors;
}
```

Test case format:

For simplicity, each node's value is the same as the node's index (1-indexed). For example, the first node with val == 1, the second node with val == 2, and so on. The graph is represented in the test case using an adjacency list.

An adjacency list is a collection of unordered lists used to represent a finite graph. Each list describes the set of neighbors of a node in the graph.

The given node will always be the first node with val = 1. You must return the copy of the given node as a reference to the cloned graph.

Example 1:
Input: adjList = [[2,4],[1,3],[2,4],[1,3]]
Output: [[2,4],[1,3],[2,4],[1,3]]
Explanation: There are 4 nodes in the graph.
1st node (val = 1)'s neighbors are 2nd node (val = 2) and 4th node (val = 4).
2nd node (val = 2)'s neighbors are 1st node (val = 1) and 3rd node (val = 3).
3rd node (val = 3)'s neighbors are 2nd node (val = 2) and 4th node (val = 4).
4th node (val = 4)'s neighbors are 1st node (val = 1) and 3rd node (val = 3).

Example 2:
Input: adjList = [[]]
Output: [[]]
Explanation: Note that the input contains one empty list. The graph consists of only one node with val = 1 and it does not have any neighbors.

Example 3:
Input: adjList = []
Output: []
Explanation: This an empty graph, it does not have any nodes.

Constraints:
* The number of nodes in the graph is in the range [0, 100].
* 1 <= Node.val <= 100
* Node.val is unique for each node.
* There are no repeated edges and no self-loops in the graph.
* The Graph is connected and all nodes can be visited starting from the given node.
*/

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
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

	nodes := make([]*Node, len(adjList))
	for i, vertices := range adjList {
		nodes[i] = &Node{
			Val:       i + 1,
			Neighbors: make([]*Node, len(vertices)),
		}
	}
	for i, vertices := range adjList {
		node := nodes[i]
		for j, val := range vertices {
			node.Neighbors[j] = nodes[val-1]
		}
	}
	return nodes[0]
}

func cloneGraphV2(node *Node) *Node {
	if node == nil {
		return nil
	}

	val2node := make(map[int]*Node)

	var traverse func(*Node, *Node)
	traverse = func(node, newNode *Node) {
		if _, ok := val2node[node.Val]; ok {
			return
		}
		val2node[node.Val] = newNode
		for _, neighbor := range node.Neighbors {
			nNode, ok := val2node[neighbor.Val]
			if !ok {
				nNode = &Node{Val: neighbor.Val}
			}
			newNode.Neighbors = append(newNode.Neighbors, nNode)
			if !ok {
				traverse(neighbor, nNode)
			}
		}
	}

	traverse(node, &Node{Val: node.Val})
	return val2node[node.Val]
}

// This variant fails on leetcode for some unclear reasons.
func cloneGraphV3(node *Node) *Node {
	if node == nil {
		return nil
	}

	type N struct {
		node, newNode *Node
	}

	val2node := make(map[int]*Node)

	queue := []N{{node, new(Node)}}
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		if _, ok := val2node[n.node.Val]; ok {
			continue
		}
		n.newNode.Val = n.node.Val
		val2node[n.node.Val] = n.newNode
		for _, neighbor := range n.node.Neighbors {
			nNode, ok := val2node[neighbor.Val]
			if !ok {
				nNode = new(Node)
				nNode.Val = neighbor.Val
			}
			n.newNode.Neighbors = append(n.newNode.Neighbors, nNode)
			if !ok {
				queue = append(queue, N{neighbor, nNode})
			}
		}
	}

	return val2node[node.Val]
}

func cloneGraphTopRated(node *Node) *Node {
	if node == nil {
		return nil
	}

	var dfs func(*Node, []*Node)
	dfs = func(node *Node, copies []*Node) {
		newNode := new(Node)
		newNode.Val = node.Val
		copies[node.Val] = newNode
		for _, neighbor := range node.Neighbors {
			if copies[neighbor.Val] == nil {
				dfs(neighbor, copies)
			}
			newNode.Neighbors = append(newNode.Neighbors, copies[neighbor.Val])
		}
	}

	copies := make([]*Node, 101)
	dfs(node, copies)
	return copies[node.Val]
}

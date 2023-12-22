package leetcode

import (
	"reflect"
	"testing"
)

func TestCloneGraph(t *testing.T) {
	var traverse func(*Node, map[int]*Node)
	traverse = func(node *Node, m map[int]*Node) {
		if _, ok := m[node.Val]; ok {
			return
		}
		m[node.Val] = node
		for _, neighbor := range node.Neighbors {
			if _, ok := m[neighbor.Val]; !ok {
				traverse(neighbor, m)
			}
		}
	}

	var cases = []struct {
		node *Node
		want *Node
	}{
		{
			node: slice2graph([][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}}),
			want: slice2graph([][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}}),
		},
		{
			node: slice2graph([][]int{{}}),
			want: slice2graph([][]int{{}}),
		},
		{
			node: slice2graph([][]int{}),
			want: slice2graph([][]int{}),
		},
	}

	for _, c := range cases {
		result := cloneGraphV3(c.node)
		if result != nil && result == c.node {
			t.Errorf("cloneGraph returned not copy but a reference to the original node")
		}
		srcAdjList := graph2slice(c.node)
		resAdjList := graph2slice(result)
		wantAdjList := graph2slice(c.want)
		if !reflect.DeepEqual(resAdjList, srcAdjList) {
			t.Errorf("cloneGraph(%v) = %v, want %v", srcAdjList, resAdjList, wantAdjList)
		}

		if result == nil {
			continue
		}

		val2nodes := make(map[int]*Node)
		val2newNodes := make(map[int]*Node)
		traverse(c.node, val2nodes)
		traverse(result, val2newNodes)
		if len(val2nodes) != len(val2newNodes) {
			t.Errorf("sized of graphs are not equal: %d != %d", len(val2nodes), len(val2newNodes))
		}
		for i := 0; i < len(val2nodes); i++ {
			if val2nodes[i+1] == val2newNodes[i+1] {
				t.Errorf("node with value %d is not copied but referenced", val2nodes[i+1].Val)
			}
		}
	}
}

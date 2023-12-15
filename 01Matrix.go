package leetcode

/*
542. 01 Matrix

Given an m x n binary matrix mat, return the distance of the nearest 0 for each cell.

The distance between two adjacent cells is 1.

Example 1:
Input: mat = [[0,0,0],[0,1,0],[0,0,0]]
Output: [[0,0,0],[0,1,0],[0,0,0]]

Example 2:
Input: mat = [[0,0,0],[0,1,0],[1,1,1]]
Output: [[0,0,0],[0,1,0],[1,2,1]]

Constraints:
* m == mat.length
* n == mat[i].length
* 1 <= m, n <= 10^4
* 1 <= m * n <= 10^4
* mat[i][j] is either 0 or 1.
* There is at least one 0 in mat.
*/

func _01Matrix(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	prevDistance := 1
	curDistance, curQueue := 2, make([][2]int, 0)
	nextDistance, nextQueue := 3, make([][2]int, 0)
	deltas := [][2]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}} // left, up, right, down

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 1 {
				hasAdjacent0 := false
				for _, d := range deltas {
					x, y := i+d[0], j+d[1]
					if x >= 0 && x < m && y >= 0 && y < n && mat[x][y] == 0 {
						hasAdjacent0 = true
						break
					}
				}
				if !hasAdjacent0 {
					mat[i][j] = curDistance
					curQueue = append(curQueue, [2]int{i, j})
				}
			}
		}
	}

	for len(curQueue) > 0 {
		var ij [2]int
		ij, curQueue = curQueue[0], curQueue[1:]
		i, j := ij[0], ij[1]
		hasAdjacentPrevDistance := false
		for _, d := range deltas {
			x, y := i+d[0], j+d[1]
			if x >= 0 && x < m && y >= 0 && y < n && mat[x][y] == prevDistance {
				hasAdjacentPrevDistance = true
				break
			}
		}
		if !hasAdjacentPrevDistance {
			mat[i][j] = nextDistance
			nextQueue = append(nextQueue, [2]int{i, j})
		}
		if len(curQueue) == 0 && len(nextQueue) > 0 {
			prevDistance = curDistance
			curDistance, curQueue, nextDistance, nextQueue = nextDistance, nextQueue, nextDistance+1, curQueue
		}
	}

	return mat
}

func _01MatrixTopRated(mat [][]int) [][]int {
	if mat == nil || len(mat) == 0 || len(mat[0]) == 0 {
		return [][]int{}
	}

	m, n := len(mat), len(mat[0])
	queue := make([][]int, 0)
	MAX_VALUE := m * n

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				queue = append(queue, []int{i, j})
			} else {
				mat[i][j] = MAX_VALUE
			}
		}
	}

	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for len(queue) > 0 {
		cell := queue[0]
		queue = queue[1:]
		for _, dir := range directions {
			r, c := cell[0]+dir[0], cell[1]+dir[1]
			if r >= 0 && r < m && c >= 0 && c < n && mat[r][c] > mat[cell[0]][cell[1]]+1 {
				queue = append(queue, []int{r, c})
				mat[r][c] = mat[cell[0]][cell[1]] + 1
			}
		}
	}

	return mat
}

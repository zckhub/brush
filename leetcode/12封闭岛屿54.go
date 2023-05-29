package main

import "fmt"

func main() {
	grid := [][]int{{1, 1, 1, 1, 1, 1, 1, 0}, {1, 0, 0, 0, 0, 1, 1, 0}, {1, 0, 1, 0, 1, 1, 1, 0}, {1, 0, 0, 0, 0, 1, 0, 1}, {1, 1, 1, 1, 1, 1, 1, 0}}
	res := closedIsland(grid)
	fmt.Println(res)
}

func closedIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	res := 0
	var dfs func(int, int)
	dfs = func(i int, j int) {
		if i < 0 || i >= m || j < 0 || j >= n {
			return
		}
		if grid[i][j] == 1 {
			return
		}
		grid[i][j] = 1
		dfs(i+1, j)
		dfs(i-1, j)
		dfs(i, j+1)
		dfs(i, j-1)
	}
	for i := 0; i < m; i++ {
		dfs(i, 0)
		dfs(i, n-1)
	}
	for j := 0; j < n; j++ {
		dfs(0, j)
		dfs(m-1, j)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				res++
				dfs(i, j)
			}
		}
	}

	return res
}

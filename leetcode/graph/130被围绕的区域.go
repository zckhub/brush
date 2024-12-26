package main

import "fmt"

func main() {
	board := [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}}
	solve1(board)
	fmt.Println(board)
}

var dirs = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func solve1(board [][]byte) {
	//把边缘的'o'遍历修改为'n'
	//把其他'o'遍历修改为'x'
	//最后把'n'改回'o'
	visited := make([][]int, len(board))
	for v := range visited {
		visited[v] = make([]int, len(board[0]))
	}
	//fmt.Println("viste",visited)
	for i := 0; i < len(board); i++ {
		dfs(board, visited, i, 0, 'N')
		dfs(board, visited, i, len(board[0])-1, 'N')
	}
	//fmt.Println("board",board)

	for j := 0; j < len(board[0]); j++ {
		dfs(board, visited, 0, j, 'N')
		dfs(board, visited, len(board)-1, j, 'N')
	}
	//fmt.Println("board",board)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if visited[i][j] == 1 {
				continue
			}
			dfs(board, visited, i, j, 'X')
		}
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'N' {
				board[i][j] = 'O'
			}
		}
	}
	return
}

func dfs(board [][]byte, visited [][]int, i int, j int, trans byte) {
	//fmt.Println("i=",i,"j=",j)
	if i > len(board)-1 || i < 0 || j > len(board[0])-1 || j < 0 || visited[i][j] == 1 || board[i][j] == 'X' {
		return
	}

	visited[i][j] = 1
	board[i][j] = trans
	for _, dir := range dirs {
		dfs(board, visited, i+dir[0], j+dir[1], trans)
	}
}

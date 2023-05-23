package main

import "fmt"

/*
51. N 皇后
困难
1.8K
相关企业

按照国际象棋的规则，皇后可以攻击与之处在同一行或同一列或同一斜线上的棋子。

n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。

给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。

每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。



示例 1：

输入：n = 4
输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
解释：如上图所示，4 皇后问题存在两个不同的解法。

示例 2：

输入：n = 1
输出：[["Q"]]

*/

func main() {
	n := 4
	res := solution(n)
	fmt.Println("res = ", res)
}

func solution(n int) [][]string {
	var res [][]int //结果集
	var queue []int //当前解
	column := make([]bool, n)
	diagnoalLeft := make([]bool, 2*n-1)
	diagnoalRight := make([]bool, 2*n-1)
	var help func(row int)
	//当前遍历第row行，选择第row行的第i列放置皇后
	help = func(row int) {
		if row == n {
			tmp := make([]int, n)
			copy(tmp, queue)
			res = append(res, tmp)
			return
		}
		for i := 0; i < n; i++ {
			if column[i] || diagnoalLeft[i-row+n-1] || diagnoalRight[row+i] {
				continue
			}
			column[i] = true
			diagnoalLeft[i-row+n-1] = true
			diagnoalRight[row+i] = true
			queue = append(queue, i)
			help(row + 1)
			queue = queue[:len(queue)-1]
			column[i] = false
			diagnoalLeft[i-row+n-1] = false
			diagnoalRight[row+i] = false
		}
	}
	help(0)
	fmt.Println("res = ", res)
	var resString [][]string
	for i := 0; i < len(res); i++ {
		tmpList := make([]string, n)
		for j := 0; j < n; j++ {
			tmp := ""

			for k := 0; k < n; k++ {
				if res[i][j] == k {
					tmp += "Q"
				} else {
					tmp += "."
				}
			}
			tmpList[j] = tmp
		}
		resString = append(resString, tmpList)
	}
	return resString
}

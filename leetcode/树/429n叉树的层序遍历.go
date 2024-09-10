package main

import "fmt"

func main() {
	root := Node{}
	res := levelOrder(&root)
	fmt.Println(res)
}

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	queue := make([]*Node, 0)
	queue = append(queue, root)
	res := make([][]int, 0)
	for len(queue) > 0 {
		//每次循环都要遍历一层
		level := make([]int, 0)
		n := len(queue)
		for i := 0; i < n; i++ {
			level = append(level, queue[i].Val)
			queue = append(queue, queue[i].Children...)

		}
		queue = queue[n:]
		res = append(res, level)

	}
	return res
}

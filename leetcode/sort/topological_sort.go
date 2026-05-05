package main

import "fmt"

func main() {
	// 示例1: 正常的有向无环图
	// 课程依赖关系: 0 <- 1, 0 <- 2, 1 <- 3, 2 <- 3
	numVertices1 := 4
	edges1 := [][]int{{3, 1}, {3, 2}, {1, 0}, {2, 0}}
	result1, ok1 := TopologicalSort(numVertices1, edges1)
	if ok1 {
		fmt.Println("拓扑排序结果:", result1)
	} else {
		fmt.Println("图中存在环，无法进行拓扑排序")
	}

	// 示例2: 包含环的图
	numVertices2 := 3
	edges2 := [][]int{{0, 1}, {1, 2}, {2, 0}}
	result2, ok2 := TopologicalSort(numVertices2, edges2)
	if ok2 {
		fmt.Println("拓扑排序结果:", result2)
	} else {
		fmt.Println("图中存在环，无法进行拓扑排序")
	}
}

// TopologicalSort 对有向图进行拓扑排序
// 输入: numVertices 顶点数, edges 有向边列表 [from, to]
// 输出: 拓扑排序结果, 是否存在环
func TopologicalSort(numVertices int, edges [][]int) ([]int, bool) {
	// 构建邻接表和入度数组
	graph := make([][]int, numVertices)
	inDegree := make([]int, numVertices)
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		inDegree[edge[1]]++
	}

	// 将所有入度为0的顶点加入队列
	queue := []int{}
	for i := 0; i < numVertices; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	// BFS遍历拓扑序
	result := []int{}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		result = append(result, node)

		// 将相邻顶点入度减1，入度为0则加入队列
		for _, neighbor := range graph[node] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	// 如果结果数量不等于顶点数，说明存在环
	if len(result) != numVertices {
		return nil, false
	}
	return result, true
}

package main

import "fmt"

func main() {
	numCourse := 5
	prere := [][]int{{1, 4}, {2, 4}, {3, 1}, {3, 2}}
	canFinish(numCourse, prere)
}

func canFinish(numCourse int, prere [][]int) bool {
	graph := make([][]int, numCourse)
	for i := range prere {
		graph[prere[i][1]] = append(graph[prere[i][1]], prere[i][0])
	}
	fmt.Println("graph", graph)
	path := make([]int, numCourse)
	visited := make([]int, numCourse)
	for j := 0; j < numCourse; j++ {
		if visited[j] > 0 {
			continue
		}
		if dGraph(j, graph, path, visited) == false {
			return false
		}
		fmt.Println("path", path, "visited", visited)
	}
	return true
}

func dGraph(index int, graph [][]int, path []int, visited []int) bool {
	if path[index] > 1 {
		return false
	}
	for i := range graph[index] {

		path[graph[index][i]]++
		visited[graph[index][i]]++
		if dGraph(graph[index][i], graph, path, visited) == false {
			return false
		}
		path[graph[index][i]]--
	}
	return true
}

package main

import (
	"coursera_algorithms/graphs"
	"fmt"
)

func main() {
	adj := [][]int{
		[]int{0, 1, 2},
		[]int{1, 2},
		[]int{2, 0, 3},
		[]int{3, 3},
	}

	gr := graphs.NewGraph(adj)
	gr.PrintGraph()
	fmt.Println(gr.Bfs(2))
	fmt.Println(gr.Dfs(2))
}

package main

import (
	"fmt"
)

func main() {
	gr := readWeightedGraphFromFile("dijkstraGraph.txt", "\t")
	gr.PrintGraph()
	paths := gr.ShortestPath(1)
	required := []int{7, 37, 59, 82, 99, 115, 133, 165, 188, 197}

	for _, val := range required {
		fmt.Print(paths[val], ",")
	}
	fmt.Println()

}

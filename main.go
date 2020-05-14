package main

import (
	"fmt"
)

func main() {
	graph := readWeightedGraphFromFileV2("g3.txt", " ")
	fmt.Println(graph.ShortestAllPairs())
}

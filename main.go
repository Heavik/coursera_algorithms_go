package main

import "fmt"

func main() {
	graph := readPrimMstGraphFromFile("prim_mst.txt")
	mstSum := graph.PrimMst()
	fmt.Println(mstSum)
}

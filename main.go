package main

import "fmt"

func main() {
	gr := readGraphFromFile("graph_data.txt")
	//gr.PrintGraph()
	fmt.Println(gr.ComputeScc(5))
}

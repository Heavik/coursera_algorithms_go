package main

import "fmt"

func main() {
	numbers := readBinaryPointsFromFile("clustering_big.txt")
	fmt.Println(clusteringBig(numbers, 24))
}

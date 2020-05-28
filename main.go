package main

import (
	"fmt"
)

func main() {
	coords := readCoordinates("tsp_huge.txt")
	//dist := getDistances(coords)
	fmt.Println(tspNearestNeighbor(coords))
}

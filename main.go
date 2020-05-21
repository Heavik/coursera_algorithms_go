package main

import (
	"fmt"
)

func main() {
	coords := readCoordinates("tsp.txt")
	dist := getDistances(coords)
	fmt.Println(tsp(dist))
}

package main

import (
	"fmt"
)

func main() {
	a, b := readTwoSat("2sat_1.txt")
	fmt.Println(isTwoSatSatisfy(a, b))
}

package main

import "fmt"

func main() {
	pairs := readPointPairsFromFile("clustering1.txt")
	fmt.Println(singleLink(pairs, 500, 4))
}

package main

import "fmt"

func main() {
	k := readKnapsackFromFile("knapsack_big.txt")
	fmt.Println(bigKnapsackOptimal(k))
}

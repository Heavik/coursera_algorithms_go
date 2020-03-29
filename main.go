package main

import (
	"fmt"
)

func main() {
	// arr := readNumbersFromFile("median.txt")
	// median := datastructs.FindMedian(arr)
	arr := readNumbersFromFile("numbers.txt")
	total := twoSumCount(arr, -10000, 10000)
	fmt.Println(total)
}

package main

import "fmt"

func main() {
	result := maxWeightSet(readNumbersFromFile("mwis.txt"))
	fmt.Println(result)
	filtered := []int{}
	test := []int{1, 2, 3, 4, 17, 117, 517, 997}

	for _, val := range test {
		if contains(result, val) {
			filtered = append(filtered, val)
		}
	}
	fmt.Println(filtered)
	fmt.Println(getDepth(huffman(readNumbersFromFile("huffman.txt"))))
}

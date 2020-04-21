package main

import "fmt"

func main() {
	tree := huffman(readNumbersFromFile("huffman.txt"))
	fmt.Println(getDepth(tree))
}

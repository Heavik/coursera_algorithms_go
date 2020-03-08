package main

import (
	"bufio"
	"coursera_algorithms/graphs"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readGraphFromFile(fileName string) graphs.Graph {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	adj := [][]int{}
	currentNode := -1
	currentList := -1
	notProcessed := make(map[int]bool)
	processed := make(map[int]bool)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		node, err := strconv.Atoi(line[0])
		if err != nil {
			fmt.Printf("Error: %v", err)
			return nil
		}
		adjNode, err := strconv.Atoi(line[1])
		if err != nil {
			fmt.Printf("Error: %v", err)
			return nil
		}
		if _, ok := processed[adjNode]; !ok {
			notProcessed[adjNode] = true
		}
		if node == currentNode {
			adj[currentList] = append(adj[currentList], adjNode)
		} else {
			currentNode = node
			currentList++
			adj = append(adj, []int{node, adjNode})
			processed[node] = true
			if _, ok := notProcessed[node]; ok {
				delete(notProcessed, node)
			}
		}
	}
	for key := range notProcessed {
		adj = append(adj, []int{key})
	}

	return graphs.NewGraph(adj)
}

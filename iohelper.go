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

	return graphs.NewGraph(adj, nil)
}

func readWeightedGraphFromFile(fileName string, delim string) graphs.Graph {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	adj := [][]int{}
	lengths := [][]int{}
	currentNode := -1
	notProcessed := make(map[int]bool)
	processed := make(map[int]bool)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), delim)
		fmt.Println(line)

		sourceNode, err := strconv.Atoi(line[0])
		if err != nil {
			fmt.Printf("Error: %v", err)
			return nil
		}
		adj = append(adj, []int{sourceNode})
		lengths = append(lengths, []int{sourceNode})
		currentNode++
		processed[sourceNode] = true
		if _, ok := notProcessed[sourceNode]; ok {
			delete(notProcessed, sourceNode)
		}
		for i := 1; i < len(line); i++ {
			if line[i] == "" {
				continue
			}
			node := strings.Split(line[i], ",")
			adjNode, err := strconv.Atoi(node[0])
			if err != nil {
				fmt.Printf("Error: %v", err)
				return nil
			}
			length, err := strconv.Atoi(node[1])
			if err != nil {
				fmt.Printf("Error: %v", err)
				return nil
			}
			if _, ok := processed[adjNode]; !ok {
				notProcessed[adjNode] = true
			}
			adj[currentNode] = append(adj[currentNode], adjNode)
			lengths[currentNode] = append(lengths[currentNode], length)
		}
	}
	for key := range notProcessed {
		adj = append(adj, []int{key})
		lengths = append(lengths, []int{key})
	}

	return graphs.NewGraph(adj, lengths)
}

func readNumbersFromFile(fileName string) []int {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil
	}
	defer file.Close()

	arr := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Printf("Error: %v", err)
			return nil
		}
		arr = append(arr, line)
	}
	return arr
}

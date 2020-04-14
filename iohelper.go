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

func readPrimMstGraphFromFile(fileName string) graphs.Graph {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	adj := [][]int{}
	lengths := [][]int{}
	adjIndex := -1
	lengthsIndex := -1
	adjIndexMap := make(map[int]int)
	lengthsIndexMap := make(map[int]int)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		sourceNode, err := strconv.Atoi(line[0])
		if err != nil {
			fmt.Printf("Error: %v", err)
			return nil
		}
		if _, ok := adjIndexMap[sourceNode]; !ok {
			adj = append(adj, []int{sourceNode})
			adjIndex++
			adjIndexMap[sourceNode] = adjIndex
		}
		destNode, err := strconv.Atoi(line[1])
		if err != nil {
			fmt.Printf("Error: %v", err)
			return nil
		}
		adj[adjIndexMap[sourceNode]] = append(adj[adjIndexMap[sourceNode]], destNode)
		if _, ok := adjIndexMap[destNode]; !ok {
			adj = append(adj, []int{destNode})
			adjIndex++
			adjIndexMap[destNode] = adjIndex
		}
		adj[adjIndexMap[destNode]] = append(adj[adjIndexMap[destNode]], sourceNode)
		length, err := strconv.Atoi(line[2])
		if err != nil {
			fmt.Printf("Error: %v", err)
			return nil
		}
		if _, ok := lengthsIndexMap[sourceNode]; !ok {
			lengths = append(lengths, []int{sourceNode})
			lengthsIndex++
			lengthsIndexMap[sourceNode] = lengthsIndex
		}
		lengths[lengthsIndexMap[sourceNode]] = append(lengths[lengthsIndexMap[sourceNode]], length)
		if _, ok := lengthsIndexMap[destNode]; !ok {
			lengths = append(lengths, []int{destNode})
			lengthsIndex++
			lengthsIndexMap[destNode] = lengthsIndex
		}
		lengths[lengthsIndexMap[destNode]] = append(lengths[lengthsIndexMap[destNode]], length)
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

func readJobsFromFile(fileName string) []job {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil
	}
	defer file.Close()

	jobs := []job{}
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	_, err = strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil
	}
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		weight, err := strconv.Atoi(line[0])
		if err != nil {
			fmt.Printf("Error: %v", err)
			return nil
		}
		length, err := strconv.Atoi(line[1])
		if err != nil {
			fmt.Printf("Error: %v", err)
			return nil
		}
		jobs = append(jobs, job{weight: weight, length: length})
	}
	return jobs
}

func readPointPairsFromFile(fileName string) []*pair {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil
	}
	defer file.Close()
	pairs := []*pair{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		from, err := strconv.Atoi(line[0])
		if err != nil {
			fmt.Printf("Error: %v", err)
			return nil
		}
		to, err := strconv.Atoi(line[1])
		if err != nil {
			fmt.Printf("Error: %v", err)
			return nil
		}
		weight, err := strconv.Atoi(line[2])
		if err != nil {
			fmt.Printf("Error: %v", err)
			return nil
		}
		p := pair{from: from, to: to, weight: weight}
		pairs = append(pairs, &p)
	}
	return pairs
}

func readBinaryPointsFromFile(fileName string) []int64 {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil
	}
	defer file.Close()
	result := []int64{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Join(strings.Fields(scanner.Text()), "")
		number, err := strconv.ParseInt(line, 2, 32)
		if err != nil {
			fmt.Printf("Error: %v", err)
			return nil
		}
		result = append(result, number)
	}
	return result
}

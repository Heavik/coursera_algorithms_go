package graphs

import (
	"coursera_algorithms/datastructs"
	"fmt"
	"strconv"
	"strings"
)

type graphNode struct {
	label   string
	adjList []int
}

type graph struct {
	vertexNum int
	vertices  map[int]graphNode
}

// NewGraph creates new graph from adj list
func NewGraph(adj [][]int) graph {
	gr := graph{vertices: make(map[int]graphNode)}
	for _, nodes := range adj {
		node := graphNode{label: strconv.Itoa(nodes[0])}
		for i := 1; i < len(nodes); i++ {
			node.adjList = append(node.adjList, nodes[i])
		}
		gr.vertices[nodes[0]] = node
		gr.vertexNum++
	}
	return gr
}

func (g graph) Bfs(start int) ([]int, error) {
	q := datastructs.NewQueue(g.vertexNum)
	q.Enqueue(start)
	visited := map[int]bool{
		start: true,
	}

	var result []int

	for !q.IsEmpty() {
		node, err := q.Dequeue()
		result = append(result, node)
		if err != nil {
			return nil, fmt.Errorf("Error: %v", err)
		}
		for _, val := range g.vertices[node].adjList {
			if !visited[val] {
				q.Enqueue(val)
				visited[val] = true
			}
		}
	}
	return result, nil
}

func (g graph) PrintGraph() {
	for _, node := range g.vertices {
		nodes := []string{}
		for _, n := range node.adjList {
			nodes = append(nodes, g.vertices[n].label)
		}
		fmt.Println(node.label, "is connected to [", strings.Join(nodes, ","), "]")
	}
}

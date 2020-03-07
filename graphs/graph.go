package graphs

import (
	"coursera_algorithms/datastructs"
	"fmt"
	"strconv"
	"strings"
)

type graphNode struct {
	label   string
	value   int
	adjList []int
}

type graph struct {
	vertexNum int
	vertices  map[int]graphNode
}

// Graph interface
type Graph interface {
	Bfs(start int) ([]int, error)
	Dfs(start int) []int
	TopoSort() []int
	PrintGraph()
}

// NewGraph creates new graph from adj list
func NewGraph(adj [][]int) Graph {
	gr := graph{vertices: make(map[int]graphNode)}
	for _, nodes := range adj {
		node := graphNode{label: strconv.Itoa(nodes[0]), value: nodes[0]}
		for i := 1; i < len(nodes); i++ {
			node.adjList = append(node.adjList, nodes[i])
		}
		gr.vertices[nodes[0]] = node
		gr.vertexNum++
	}
	return &gr
}

func (g *graph) Bfs(start int) ([]int, error) {
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

func dfs(g *graph, node int, visited map[int]bool, result []int) []int {
	visited[node] = true
	result = append(result, node)

	for _, val := range g.vertices[node].adjList {
		if !visited[val] {
			result = dfs(g, val, visited, result)
		}
	}
	return result
}

func (g *graph) Dfs(start int) []int {
	visited := make(map[int]bool)
	result := []int{}

	result = dfs(g, start, visited, result)
	return result
}

func dfsTopo(g *graph, node int, visited map[int]bool, result []int) []int {
	visited[node] = true

	for _, v := range g.vertices[node].adjList {
		if !visited[v] {
			result = dfsTopo(g, v, visited, result)
		}
	}
	result = append(result, node)
	return result
}

func (g *graph) TopoSort() []int {
	visited := make(map[int]bool)
	result := []int{}

	for _, v := range g.vertices {
		if !visited[v.value] {
			result = dfsTopo(g, v.value, visited, result)
		}
	}

	rLen := len(result)
	for i, j := 0, rLen-1; i < rLen/2; i++ {
		result[i], result[j] = result[j], result[i]
		j--
	}

	return result
}

func (g *graph) PrintGraph() {
	for _, node := range g.vertices {
		nodes := []string{}
		for _, n := range node.adjList {
			nodes = append(nodes, g.vertices[n].label)
		}
		fmt.Println(node.label, "is connected to [", strings.Join(nodes, ","), "]")
	}
}

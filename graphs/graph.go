package graphs

import (
	"coursera_algorithms/datastructs"
	"fmt"
	"sort"
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
	vertices  map[int]*graphNode
}

// Graph interface
type Graph interface {
	Bfs(start int) ([]int, error)
	Dfs(start int) []int
	TopoSort() []int
	PrintGraph()
	ComputeScc(topNum int) (int, []int)
}

// NewGraph creates new graph from adj list
func NewGraph(adj [][]int) Graph {
	gr := graph{vertices: make(map[int]*graphNode)}
	for _, nodes := range adj {
		node := graphNode{label: strconv.Itoa(nodes[0]), value: nodes[0]}
		for i := 1; i < len(nodes); i++ {
			node.adjList = append(node.adjList, nodes[i])
		}
		gr.vertices[nodes[0]] = &node
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

func reverseGraph(g *graph) *graph {
	reversed := graph{vertices: make(map[int]*graphNode)}
	for _, node := range g.vertices {
		for _, adjNode := range node.adjList {
			if n, ok := reversed.vertices[adjNode]; ok {
				n.adjList = append(n.adjList, node.value)
			} else {
				reversed.vertices[adjNode] = &graphNode{
					label: strconv.Itoa(adjNode),
					value: adjNode, adjList: []int{node.value},
				}
				reversed.vertexNum++
			}
		}
	}
	for _, node := range g.vertices {
		if _, ok := reversed.vertices[node.value]; !ok {
			reversed.vertices[node.value] = &graphNode{
				label: strconv.Itoa(node.value),
				value: node.value,
			}
			reversed.vertexNum++
		}
	}
	return &reversed
}

func dfsScc(g *graph, node int, visited map[int]bool) int {
	nodeCount := 1
	visited[node] = true
	for _, v := range g.vertices[node].adjList {
		if !visited[v] {
			nodeCount += dfsScc(g, v, visited)
		}
	}
	return nodeCount
}

func updateTopScc(topScc []int, value int) {
	topScc[len(topScc)-1] = value
	sort.Slice(topScc, func(i, j int) bool {
		return topScc[i] > topScc[j]
	})
}

// problem answer is 434821,968,459,313,211
func (g *graph) ComputeScc(topNum int) (int, []int) {
	reversedGr := reverseGraph(g)
	topo := reversedGr.TopoSort()

	numScc := 0
	visited := make(map[int]bool)
	topScc := make([]int, topNum+1)
	for _, node := range topo {
		if !visited[node] {
			numScc++
			count := dfsScc(g, node, visited)
			updateTopScc(topScc, count)
		}
	}
	return numScc, topScc[:topNum]
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

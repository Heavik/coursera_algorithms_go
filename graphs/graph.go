package graphs

import (
	"coursera_algorithms/datastructs"
	"coursera_algorithms/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const MAX_DISTANCE = 1000000

type edge struct {
	to     int
	length int
}

type graphNode struct {
	label   string
	value   int
	edges   map[int]*edge
	adjList []int
	key     int
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
	ShortestPath(start int) map[int]int
	PrimMst() int
	ShortestAllPairs() int
}

// NewGraph creates new graph from adj and length lists
func NewGraph(adj [][]int, lengths [][]int) Graph {
	gr := graph{vertices: make(map[int]*graphNode)}
	for index, nodes := range adj {
		node := graphNode{label: strconv.Itoa(nodes[0]), value: nodes[0]}
		if lengths != nil {
			node.edges = make(map[int]*edge)
		}
		for i := 1; i < len(nodes); i++ {
			node.adjList = append(node.adjList, nodes[i])
			if lengths != nil {
				edge := edge{to: nodes[i], length: lengths[index][i]}
				node.edges[nodes[i]] = &edge
			}
		}
		gr.vertices[nodes[0]] = &node
		gr.vertexNum++
	}
	return &gr
}

func (node *graphNode) GetVal() int {
	return node.value
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

func getMinLenVert(g *graph, dist map[int]int, processed map[int]bool) (int, int) {
	minLength := MAX_DISTANCE
	minVertex := -1
	for _, node := range g.vertices {
		if processed[node.value] {
			for to, edge := range node.edges {
				if !processed[to] && minLength > dist[node.value]+edge.length {
					minLength = dist[node.value] + edge.length
					minVertex = edge.to
				}
			}
		}
	}
	return minVertex, minLength
}

// problem answer is 2599,2610,2947,2052,2367,2399,2029,2442,2505,3068
func (g *graph) ShortestPath(start int) map[int]int {
	dist := make(map[int]int)
	processed := make(map[int]bool)

	for _, node := range g.vertices {
		dist[node.value] = MAX_DISTANCE
	}
	dist[start] = 0
	processed[start] = true
	for vert, length := getMinLenVert(g, dist, processed); vert != -1; {
		processed[vert] = true
		dist[vert] = length
		vert, length = getMinLenVert(g, dist, processed)
	}
	return dist
}

// problem answer is -3612829
func (g *graph) PrimMst() int {
	pq := datastructs.NewPQ(func(a, b interface{}) int {
		return a.(*graphNode).key - b.(*graphNode).key
	})
	startVertex := 1
	processed := make(map[int]bool)
	winner := make(map[int]*edge)
	result := make([]*edge, 0)
	for _, vertex := range g.vertices {
		if edge, ok := vertex.edges[startVertex]; ok {
			vertex.key = edge.length
			winner[vertex.value] = edge
		} else {
			vertex.key = MAX_DISTANCE
			winner[vertex.value] = nil
		}
		pq.Enqueue(vertex)
	}
	processed[startVertex] = true

	for !pq.IsEmpty() {
		minVertex := pq.Dequeue().(*graphNode)
		processed[minVertex.value] = true
		result = append(result, winner[minVertex.value])

		for _, edge := range minVertex.edges {
			if _, ok := processed[edge.to]; !ok {
				if edge.length < g.vertices[edge.to].key {
					pq.Remove(g.vertices[edge.to])
					g.vertices[edge.to].key = edge.length
					winner[edge.to] = edge
					pq.Enqueue(g.vertices[edge.to])
				}
			}
		}
	}
	lengthSum := 0
	for _, edge := range result {
		if edge != nil {
			lengthSum += edge.length
		}
	}
	return lengthSum
}

// problem answer is -19
func (g *graph) ShortestAllPairs() int {
	dist := make([][]int, g.vertexNum)
	for i := range dist {
		dist[i] = make([]int, g.vertexNum)
	}

	for i := 0; i < g.vertexNum; i++ {
		for j := 0; j < g.vertexNum; j++ {
			dist[i][j] = MAX_DISTANCE
		}
	}

	for _, vertex := range g.vertices {
		dist[vertex.value-1][vertex.value-1] = 0
		for _, edge := range vertex.edges {
			dist[vertex.value-1][edge.to-1] = edge.length
		}
	}

	for k := 0; k < g.vertexNum; k++ {
		for i := 0; i < g.vertexNum; i++ {
			for j := 0; j < g.vertexNum; j++ {
				if dist[i][k]+dist[k][j] < dist[i][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}

	for i := 0; i < g.vertexNum; i++ {
		if dist[i][i] < 0 {
			return -MAX_DISTANCE
		}
	}

	min := MAX_DISTANCE

	for i := 0; i < g.vertexNum; i++ {
		for j := 0; j < g.vertexNum; j++ {
			if i != j {
				min = utils.GetMin(min, dist[i][j])
			}
		}
	}

	return min
}

func (g *graph) PrintGraph() {
	for _, node := range g.vertices {
		nodes := []string{}
		for _, n := range node.adjList {
			nodes = append(nodes, g.vertices[n].label)
		}
		fmt.Println(node.label, "is connected to [", strings.Join(nodes, ","), "]")
		for _, edge := range node.edges {
			fmt.Println("Length from", node.label, "to", edge.to, "is", edge.length)
		}
	}
}

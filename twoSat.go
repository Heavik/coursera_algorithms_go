package main

import "coursera_algorithms/datastructs"

// problem answer is 101100
func isTwoSatSatisfy(a, b []int) bool {
	numOfVars := len(a)
	numOfClauses := numOfVars
	numOfNodes := 2 * numOfVars

	visited := make([]bool, numOfNodes+1)
	visitedInversed := make([]bool, numOfNodes+1)
	adjList := make([][]int, numOfNodes+1)
	scc := make([]int, numOfNodes+1)
	adjListInv := make([][]int, numOfNodes+1)
	stack := datastructs.NewStack()

	addEdge := func(from, to int, addTo [][]int) {
		addTo[from] = append(addTo[from], to)
	}

	var dfs func(i int)
	var dfsInverse func(i, counter int)

	dfs = func(i int) {
		if visited[i] {
			return
		}

		visited[i] = true
		for j := 0; j < len(adjList[i]); j++ {
			dfs(adjList[i][j])
		}

		stack.Push(i)
	}

	dfsInverse = func(i, counter int) {
		if visitedInversed[i] {
			return
		}

		visitedInversed[i] = true
		for j := 0; j < len(adjListInv[i]); j++ {
			dfsInverse(adjListInv[i][j], counter)
		}

		scc[i] = counter
	}

	for i := 0; i < numOfClauses; i++ {
		if a[i] > 0 && b[i] > 0 {
			addEdge(a[i]+numOfVars, b[i], adjList)
			addEdge(b[i], a[i]+numOfVars, adjListInv)
			addEdge(b[i]+numOfVars, a[i], adjList)
			addEdge(a[i], b[i]+numOfVars, adjListInv)
		} else if a[i] > 0 && b[i] < 0 {
			addEdge(a[i]+numOfVars, numOfVars-b[i], adjList)
			addEdge(numOfVars-b[i], a[i]+numOfVars, adjListInv)
			addEdge(-b[i], a[i], adjList)
			addEdge(a[i], -b[i], adjListInv)
		} else if a[i] < 0 && b[i] > 0 {
			addEdge(-a[i], b[i], adjList)
			addEdge(b[i], -a[i], adjListInv)
			addEdge(b[i]+numOfVars, numOfVars-a[i], adjList)
			addEdge(numOfVars-a[i], b[i]+numOfVars, adjListInv)
		} else {
			addEdge(-a[i], numOfVars-b[i], adjList)
			addEdge(numOfVars-b[i], -a[i], adjListInv)
			addEdge(-b[i], numOfVars-a[i], adjList)
			addEdge(numOfVars-a[i], -b[i], adjListInv)
		}
	}

	for i := 1; i <= numOfNodes; i++ {
		if !visited[i] {
			dfs(i)
		}
	}

	counter := 1
	for !stack.IsEmpty() {
		node, _ := stack.Pop()
		if !visitedInversed[node] {
			dfsInverse(node, counter)
			counter++
		}
	}

	for i := 1; i <= numOfVars; i++ {
		if scc[i] == scc[i+numOfVars] {
			return false
		}
	}
	return true
}

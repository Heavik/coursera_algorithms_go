package main

import (
	"coursera_algorithms/utils"
	"math"
)

type knapsack struct {
	capacity int
	items    int
	values   []int
	weights  []int
}

// problem answer is 10100110
func maxWeightSet(weights []int) []int {
	maxVals := make([]int, len(weights)+1)
	maxVals[0] = 0
	maxVals[1] = weights[0]

	for i := 2; i < len(maxVals); i++ {
		maxVals[i] = utils.GetMax(maxVals[i-1], maxVals[i-2]+weights[i-1])
	}

	result := []int{}

	for i := len(maxVals) - 1; i >= 1; {
		v := 0
		if i > 1 {
			v = maxVals[i-2]
		}
		if maxVals[i-1] >= v+weights[i-1] {
			i--
		} else {
			result = append(result, i)
			i -= 2
		}
	}

	return result
}

// problem answer is 2493893
func knapsackOptimal(k *knapsack) int {
	arr := make([][]int, k.items+1)
	for i := range arr {
		arr[i] = make([]int, k.capacity+1)
	}

	for i := 1; i < len(arr); i++ {
		for x := 0; x <= k.capacity; x++ {
			if k.weights[i-1] > x {
				arr[i][x] = arr[i-1][x]
			} else {
				arr[i][x] = utils.GetMax(arr[i-1][x], arr[i-1][x-k.weights[i-1]]+k.values[i-1])
			}
		}
	}

	return arr[k.items][k.capacity]
}

// problem answer is 4243395
func bigKnapsackOptimal(k *knapsack) int {
	prev := make([]int, k.capacity+1)
	current := make([]int, k.capacity+1)

	for i := 0; i < k.items; i++ {
		for x := 0; x <= k.capacity; x++ {
			if k.weights[i] > x {
				current[x] = prev[x]
			} else {
				current[x] = utils.GetMax(prev[x], prev[x-k.weights[i]]+k.values[i])
			}
		}
		for i, val := range current {
			prev[i] = val
		}
	}

	return current[k.capacity]
}

func calcDistance(x, y, z, w float64) float64 {
	return math.Sqrt(math.Pow(x-z, 2) + math.Pow(y-w, 2))
}

func getDistances(points [][]float64) [][]float64 {
	dist := make([][]float64, len(points))
	for i := range dist {
		dist[i] = make([]float64, len(points))
	}

	for i := 0; i < len(dist); i++ {
		for j := 0; j < len(dist[i]); j++ {
			if i == j {
				dist[i][j] = 0.0
			} else {
				dist[i][j] = calcDistance(points[i][0], points[i][1], points[j][0], points[j][1])
			}
		}
	}

	return dist
}

//problem answer is 26442
func tsp(dist [][]float64) float64 {
	dp := make([][]float64, int(math.Pow(2, float64(len(dist)))))
	for i := range dp {
		dp[i] = make([]float64, len(dist))
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	visitedAll := (1 << len(dist)) - 1

	var tspHelper func(mask int, pos int) float64

	tspHelper = func(mask int, pos int) float64 {
		if mask == visitedAll {
			return dist[pos][0]
		}
		if dp[mask][pos] != -1 {
			return dp[mask][pos]
		}

		ans := math.MaxFloat64

		for city := 0; city < len(dist); city++ {
			if (mask & (1 << city)) == 0 {
				newAns := dist[pos][city] + tspHelper(mask|(1<<city), city)
				ans = math.Min(ans, newAns)
			}
		}

		dp[mask][pos] = ans
		return dp[mask][pos]
	}
	return tspHelper(1, 0)
}

// problem answer is 1203406
func tspNearestNeighbor(coords [][]float64) float64 {
	result := 0.0
	lastVisited := 0

	visited := make(map[int]bool, len(coords))
	visited[lastVisited] = true

	for i := 0; i < len(coords)-1; i++ {
		minDist := math.MaxFloat64
		nextCity := -1
		for j := 0; j < len(coords); j++ {
			if !visited[j] {
				dist := calcDistance(coords[lastVisited][0], coords[lastVisited][1], coords[j][0], coords[j][1])
				if dist < minDist {
					minDist = dist
					nextCity = j
				}
			}
		}
		if nextCity != -1 {
			result += minDist
			visited[nextCity] = true
			lastVisited = nextCity
		}
	}

	result += calcDistance(coords[lastVisited][0], coords[lastVisited][1], coords[0][0], coords[0][1])
	return result
}

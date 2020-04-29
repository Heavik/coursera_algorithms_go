package main

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
		maxVals[i] = getMax(maxVals[i-1], maxVals[i-2]+weights[i-1])
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
				arr[i][x] = getMax(arr[i-1][x], arr[i-1][x-k.weights[i-1]]+k.values[i-1])
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
				current[x] = getMax(prev[x], prev[x-k.weights[i]]+k.values[i])
			}
		}
		for i, val := range current {
			prev[i] = val
		}
	}

	return current[k.capacity]
}

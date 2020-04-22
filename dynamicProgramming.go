package main

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

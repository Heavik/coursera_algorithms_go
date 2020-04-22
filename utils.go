package main

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func contains(arr []int, val int) bool {
	for _, v := range arr {
		if val == v {
			return true
		}
	}
	return false
}

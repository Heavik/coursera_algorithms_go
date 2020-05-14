package utils

func GetMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func GetMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Contains(arr []int, val int) bool {
	for _, v := range arr {
		if val == v {
			return true
		}
	}
	return false
}

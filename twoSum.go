package main

import "sort"

// assigment answer is 427
func twoSumCount(arr []int, loBound int64, highBound int64) int {
	sort.Ints(arr)
	sums := make(map[int64]bool)

	lo := 0
	high := len(arr) - 1
	for high > lo {
		sum := int64(arr[high]) + int64(arr[lo])
		if sum > highBound {
			high--
			continue
		}
		if sum < loBound {
			lo++
			continue
		}
		innerLo := lo
		for sum >= loBound && sum <= highBound {
			sums[sum] = true
			innerLo++
			sum = int64(arr[innerLo]) + int64(arr[high])
		}
		high--
	}
	return len(sums)
}

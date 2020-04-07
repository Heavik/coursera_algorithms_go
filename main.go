package main

import (
	"fmt"
)

func main() {
	// arr := readNumbersFromFile("median.txt")
	// median := datastructs.FindMedian(arr)
	//arr := readNumbersFromFile("numbers.txt")
	//total := twoSumCount(arr, -10000, 10000)
	//fmt.Println(total)
	jobs := []job{
		{weight: 5, length: 3},
		{weight: 2, length: 1},
	}
	jobs = readJobsFromFile("jobs.txt")
	fmt.Println(jobScheduleDiff(jobs))
	fmt.Println(jobScheduleRatio(jobs))
}

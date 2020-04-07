package main

import "sort"

type job struct {
	weight int
	length int
}

func calcRunningTime(jobs []job, less func(i, j int) bool) int {
	sort.Slice(jobs, less)

	length := 0
	sum := 0
	for _, job := range jobs {
		length += job.length
		sum += length * job.weight
	}
	return sum
}

// assigment answer is 69119377652
func jobScheduleDiff(jobs []job) int {
	less := func(i, j int) bool {
		diff1 := jobs[i].weight - jobs[i].length
		diff2 := jobs[j].weight - jobs[j].length

		if diff1 == diff2 {
			return jobs[i].weight > jobs[j].weight
		}
		return diff1 > diff2
	}
	return calcRunningTime(jobs, less)
}

// assigment answer is 67311454237
func jobScheduleRatio(jobs []job) int {
	less := func(i, j int) bool {
		ratio1 := float32(jobs[i].weight) / float32(jobs[i].length)
		ratio2 := float32(jobs[j].weight) / float32(jobs[j].length)

		return ratio1 > ratio2
	}
	return calcRunningTime(jobs, less)
}

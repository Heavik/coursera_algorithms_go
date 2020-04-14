package main

import (
	"coursera_algorithms/datastructs"
	"sort"
)

type pair struct {
	from   int
	to     int
	weight int
}

// problem answer is 106
func singleLink(pairs []*pair, nodes int, k int) int {
	sort.Slice(pairs, func(p1, p2 int) bool { return pairs[p1].weight < pairs[p2].weight })
	uf := datastructs.InitUnionFind(nodes)
	clusters := nodes
	for i := 0; clusters > k; i++ {
		p := pairs[i]
		if !uf.Connected(int64(p.from), int64(p.to)) {
			uf.Union(int64(p.from), int64(p.to))
			clusters--
		}
	}
	result := 0
	for _, p := range pairs {
		if !uf.Connected(int64(p.from), int64(p.to)) {
			result = p.weight
			break
		}
	}

	return result
}

// problem answer is 6118
func clusteringBig(numbers []int64, bits int) int {
	numMap := make(map[int64][]int)
	for i, num := range numbers {
		numMap[num] = append(numMap[num], i)
	}
	uf := datastructs.UnionFindFromSlice(numbers)

	for i, num := range numbers {
		for i1 := 0; i1 < bits; i1++ {
			flip := num ^ (1 << i1)
			unionNodes(i, flip, uf, numMap)
		}
		for i1 := 0; i1 < bits; i1++ {
			for i2 := 0; i2 < bits; i2++ {
				flip := num ^ (1 << i1) ^ (1 << i2)
				unionNodes(i, flip, uf, numMap)
			}
		}
	}
	clusters := 0
	for i := range numbers {
		find := uf.Find(int64(i))
		if find == int64(i) {
			clusters++
		}
	}
	return clusters
}

func unionNodes(n1 int, n2 int64, uf datastructs.UnionFind, numMap map[int64][]int) {
	if indexes, ok := numMap[n2]; ok {
		for _, idx := range indexes {
			uf.Union(int64(n1), int64(idx))
		}
	}
}

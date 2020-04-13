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
		if !uf.Connected(p.from, p.to) {
			uf.Union(p.from, p.to)
			clusters--
		}
	}
	result := 0
	for _, p := range pairs {
		if !uf.Connected(p.from, p.to) {
			result = p.weight
			break
		}
	}

	return result
}

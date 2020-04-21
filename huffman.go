package main

import (
	"coursera_algorithms/datastructs"
	"math"
)

type node struct {
	freq  int
	left  *node
	right *node
}

func (n *node) isLeaf() bool {
	return n.left == nil && n.right == nil
}

func huffman(freuqs []int) *node {
	pq := datastructs.NewPQ(func(a, b interface{}) int { return a.(*node).freq - b.(*node).freq })
	for _, fr := range freuqs {
		pq.Enqueue(&node{freq: fr})
	}
	for pq.Size() > 1 {
		left := pq.Dequeue().(*node)
		right := pq.Dequeue().(*node)

		parent := node{
			freq:  left.freq + right.freq,
			left:  left,
			right: right,
		}
		pq.Enqueue(&parent)
	}
	root := pq.Dequeue().(*node)
	return root
}

// problem answer is 9, 19
func getDepth(tree *node) (int, int) {
	min := math.MaxInt32
	max := math.MinInt32

	getMin := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	getMax := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var getDepthRec func(node *node, depth int)
	getDepthRec = func(node *node, depth int) {
		if node.isLeaf() {
			min = getMin(depth, min)
			max = getMax(depth, max)
		} else {
			if node.left != nil {
				getDepthRec(node.left, depth+1)
			}
			if node.right != nil {
				getDepthRec(node.right, depth+1)
			}
		}
	}
	getDepthRec(tree, 0)
	return min, max
}

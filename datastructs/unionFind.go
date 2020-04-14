package datastructs

// UnionFind datastructure interface
type UnionFind interface {
	Union(s1, s2 int64)
	Find(val int64) int64
	Connected(v1, v2 int64) bool
}

type disjointSet struct {
	elements []int64
	size     []int
}

// InitUnionFind creates union find datastruct from the given number [1..nums]
func InitUnionFind(nums int) UnionFind {
	ds := disjointSet{elements: make([]int64, nums+1), size: make([]int, nums+1)}
	for i := 1; i <= nums; i++ {
		ds.elements[i] = int64(i)
		ds.size[i] = 1
	}
	return &ds
}

// UnionFindFromSlice creates union find datastruct from slice
func UnionFindFromSlice(slice []int64) UnionFind {
	ds := disjointSet{elements: make([]int64, len(slice)), size: make([]int, len(slice))}
	for i := 0; i < len(slice); i++ {
		ds.elements[i] = int64(i)
		ds.size[i] = 1
	}
	return &ds
}

func (s *disjointSet) Union(s1, s2 int64) {
	s1, s2 = s.Find(s1), s.Find(s2)
	if s1 == s2 {
		return
	}
	if s.size[s1] > s.size[s2] {
		s1, s2 = s2, s1
	}
	s.elements[s1] = s2
	s.size[s2] += s.size[s1]
}

func (s *disjointSet) Find(val int64) int64 {
	if val != s.elements[val] {
		s.elements[val] = s.Find(s.elements[val])
	}
	return s.elements[val]
}

func (s *disjointSet) Connected(v1, v2 int64) bool {
	return s.Find(v1) == s.Find(v2)
}

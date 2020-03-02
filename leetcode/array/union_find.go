package main

type unionFind struct {
	id   []int
	size []int
}

func newUF(n int) *unionFind {
	uf := &unionFind{
		id:   make([]int, n),
		size: make([]int, n),
	}

	for i := 0; i < n; i++ {
		uf.id[i] = i
		uf.size[i] = 1
	}

	return uf
}

func (uf *unionFind) findRoot(p int) int {
	for p != uf.id[p] {
		uf.id[p] = uf.id[uf.id[p]]
		p = uf.id[p]
	}
	return p
}

func (uf *unionFind) union(a, b int) {
	rootA := uf.findRoot(a)
	rootB := uf.findRoot(b)

	if rootA == rootB {
		return
	}

	if uf.size[rootA] < uf.size[rootB] {
		uf.id[rootA] = rootB
		uf.size[rootB] += uf.size[rootA]
	} else {
		uf.id[rootB] = rootA
		uf.size[rootA] += uf.size[rootB]
	}
}

func (uf *unionFind) maxSize() int {
	max := 0
	for _, s := range uf.size {
		if s > max {
			max = s
		}
	}
	return max
}

func longestConsecutiveSequence(arr []int) int {
	n := len(arr)
	uf := newUF(n)
	hash := make(map[int]int)

	for i := 0; i < n; i++ {
		hash[arr[i]] = i
	}

	for _, num := range arr {
		if _, ok := hash[num+1]; ok {
			uf.union(hash[num], hash[num+1])
		}
	}

	return uf.maxSize()
}

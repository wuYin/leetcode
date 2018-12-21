package main

type MinHeap []int

// heapify
func NewMinHeap(nums []int) *MinHeap {
	hs := MinHeap(nums)
	n := len(hs)
	heap := &hs
	for i := n/2 - 1; i >= 0; i-- {
		heap.down(i, n)
	}
	return heap
}

func (h *MinHeap) Push(v int) {
	*h = append(*h, v)
	h.up(len(*h) - 1)
}

func (h *MinHeap) Pop() int {
	hs := *h
	min := hs[0]

	n := len(hs)
	h.swap(0, n-1)
	h.down(0, n-1) // 全体下滤
	*h = hs[:n-1]

	return min
}

// 下滤
func (h *MinHeap) down(mid, n int) {
	for {
		l := 2*mid + 1
		if l >= n {
			break
		}

		min := l
		if r := l + 1; r < n && h.less(r, l) {
			min = r
		}
		if !h.less(min, mid) {
			break
		}

		h.swap(mid, min) // ok
		mid = min
	}
}

// 上滤
func (h *MinHeap) up(i int) {
	for {
		parent := (i - 1) / 2
		if h.less(parent, i) || parent == i {
			break
		}
		h.swap(parent, i)
		i = parent
	}
}

func (h *MinHeap) less(i, j int) bool {
	hs := *h
	return hs[i] < hs[j]
}

func (h *MinHeap) swap(i, j int) {
	hs := *h
	// fmt.Println("swap", hs, hs[i], hs[j]) // debug
	hs[i], hs[j] = hs[j], hs[i]
}

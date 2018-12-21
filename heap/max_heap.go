package main

type MaxHeap []int

// heapify
func NewMaxHeap(nums []int) *MaxHeap {
	hs := MaxHeap(nums)
	n := len(hs)
	h := &hs
	for i := n/2 - 1; i >= 0; i-- {
		h.down(i, n)
	}
	return h
}

func (h *MaxHeap) Push(v int) {
	*h = append(*h, v)
	h.up(len(*h) - 1)
}

func (h *MaxHeap) Pop() int {
	hs := *h
	max := hs[0]
	n := len(hs)

	h.swap(0, n-1)
	h.down(0, n-1) // 除最后一个元素外全体下滤
	*h = hs[:n-1]

	return max
}

// 上滤
func (h *MaxHeap) up(i int) {
	for {
		parent := (i - 1) / 2
		if h.more(parent, i) || parent == i {
			break
		}
		h.swap(parent, i)
		i = parent
	}
}

// 下滤
func (h *MaxHeap) down(mid, n int) {
	for {
		l := 2*mid + 1
		if l >= n || l < 0 {
			break
		}
		max := l
		if r := l + 1; r < n && h.more(r, max) {
			max = r
		}
		if !h.more(max, mid) {
			break
		}

		h.swap(mid, max)
		mid = max
	}
}

func (h *MaxHeap) swap(i, j int) {
	hs := *h
	hs[i], hs[j] = hs[j], hs[i]
}

func (h *MaxHeap) more(i, j int) bool {
	hs := *h
	return hs[i] > hs[j]
}

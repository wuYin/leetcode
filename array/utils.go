package main

func min(vs ...int) int {
	if len(vs) == 0 {
		panic("func min args empty")
	}
	m := vs[0]
	for _, n := range vs[1:] {
		if n < m {
			m = n
		}
	}
	return m
}

func max(vs ...int) int {
	if len(vs) == 0 {
		panic("func min args empty")
	}
	m := vs[0]
	for _, n := range vs[1:] {
		if n > m {
			m = n
		}
	}
	return m
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

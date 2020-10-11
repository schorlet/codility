package main

func Solution(A []int) int {
	m := make(map[int]bool, (len(A)/2)+1)
	for _, i := range A {
		if m[i] {
			delete(m, i)
		} else {
			m[i] = true
		}
	}
	for k := range m {
		return k
	}
	panic(m)
}

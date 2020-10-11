package main

func Solution(A []int) int {
	lenB := len(A) + 1
	sum := ((lenB + 1) * (lenB)) / 2
	for _, a := range A {
		sum -= a
	}
	return sum
}

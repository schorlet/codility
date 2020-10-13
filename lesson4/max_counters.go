package main

import (
	"fmt"
)

func main() {
	A := []int{3, 4, 4, 6, 1, 4, 4}
	fmt.Println(Solution(5, A))
}

func Solution(N int, A []int) []int {
	B := make([]int, N+1)
	start, max := 0, 0
	for _, a := range A {
		if a == N+1 {
			B = make([]int, N+1)
			start = max
			continue
		}
		B[a] += 1
		if B[a]+start > max {
			max = B[a]+start
		}
	}
	for i := range B {
		B[i] += start
	}
	return B[1:]
}

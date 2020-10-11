package main

import (
	"fmt"
)

func main() {
	A := []int{1, 3, 1, 4, 2, 3, 5, 4}
	fmt.Println(6 == Solution(5, A))
	fmt.Println(-1 == Solution(6, A))

	B := []int{1, 3, 1, 3, 5, 4, 5, 9, 1, 5, 4, 2, 5, 6}
	fmt.Println(11 == Solution(5, B))

	D := []int{1, 3, 1, 3, 2, 1, 3}
	fmt.Println(4 == Solution(3, D))
}

func Solution(X int, A []int) int {
	sum := ((X + 1) * X) / 2
	cur := 0
	seen := make([]int, X+1)
	for i, a := range A {
		if a > X {
			continue
		}
		if seen[a] == 0 {
			seen[a] = a
			cur += a
			if cur == sum {
				return i
			}
		}
	}
	return -1
}

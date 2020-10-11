package solution

import (
	"fmt"
)

func main() {
	fmt.Println(1 == Solution([]int{3, 1, 2, 4, 3}))
	fmt.Println(0 == Solution([]int{-2, -3, -4, -1}))
	fmt.Println(0 == Solution([]int{1, 2, 3, 4, 2}))
	fmt.Println(1 == Solution([]int{1, 2}))
	fmt.Println(2000 == Solution([]int{-1000, 1000}))
}

func Solution(A []int) int {
	total := 0
	for _, a := range A {
		total += a
	}
	sum1, sum2 := 0, total
	best := 1<<31 - 1
	for _, a := range A[:len(A)-1] {
		sum1 += a
		sum2 -= a
		if d := diff(sum1, sum2); d < best {
			best = d
		}
	}
	return best
}

func diff(x, y int) int {
	return abs(x - y)
}

func abs(z int) int {
	if z < 0 {
		return -z
	}
	return z
}

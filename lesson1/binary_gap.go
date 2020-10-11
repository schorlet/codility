package solution

import (
	"fmt"
)

func Solution(N int) int {
	runes := []rune(fmt.Sprintf("%b", N))
	zeros, max := 0, 0
	for i := 1; i < len(runes); i++ {
		rn := runes[i]
		if rn == '0' {
			zeros += 1
			continue
		}
		if zeros > max {
			max = zeros
		}
		zeros = 0
	}
	return max
}

package main

func Solution(A []int, K int) []int {
	lenA := len(A)
	if lenA == 0 || K == 0 {
		return A
	}
	if lenA == K || K%lenA == 0 {
		return A
	}
	B := make([]int, K+lenA)
	copy(B[K:], A)
	for i := 0; i < K; i++ {
		B[K-1-i] = B[lenA+K-1-i]
	}
	return B[:lenA]
}

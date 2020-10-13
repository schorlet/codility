package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	A := []int{3, 4, 4, 6, 1, 4, 4}
	fmt.Println("[3 2 2 4 2]" == fmt.Sprintf("%d", Solution(5, A)))

	B := []int{1, 1, 2, 1, 1}
	fmt.Println("[4]" == fmt.Sprintf("%d", Solution(1, B)))

	C := []int{1, 1, 3, 2, 2}
	fmt.Println("[2 4]" == fmt.Sprintf("%d", Solution(2, C)))

	perfTest()
}

func perfTest() {
	rand.Seed(time.Now().UnixNano())
	D := make([]int, 100_000)
	for i := range D {
		if rand.Intn(10)%2 == 0 {
			D[i] = 10_000
		} else {
			D[i] = rand.Intn(10_000)
		}
	}

	var start time.Time
	done := make(chan []int)
	go func() {
		start = time.Now()
		done <- Solution(9_999, D)
	}()

	ticker := time.NewTicker(100 * time.Millisecond)
	select {
	case <-done:
		ticker.Stop()
		fmt.Printf("ok: %s\n", time.Since(start))
	case <-ticker.C:
		<-done
		fmt.Printf("timeout: %s\n", time.Since(start))
	}
}

func Solution(N int, A []int) []int {
	B := make([]int, N+1)
	start, max := 0, 0
	for _, a := range A {
		if a == N+1 {
			start = max
			continue
		}
		if B[a] < start {
			B[a] = start
		}
		B[a] += 1
		if B[a] > max {
			max = B[a]
		}
	}
	for i := range B {
		if B[i] < start {
			B[i] = start
		}
	}
	return B[1:]
}

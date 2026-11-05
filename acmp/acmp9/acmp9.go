package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)

	a := make([]int, n)

	for i := range a {
		fmt.Fscan(in, &a[i])
	}

	sum, mul := sliceMaxMin(a)

	fmt.Print(sum, mul)
}

func sliceMaxMin(slice []int) (int, int) {

	sum := 0
	maxI := slice[0]
	minI := slice[0]
	for i, v := range slice[1:] {
		if v > 0 {
			sum += v
		}
		if v > maxI {
			maxI = i
		}
		if v < minI {
			minI = i
		}
	}

	mul := 1
	for i := min(maxI, minI) + 1; i < max(maxI, minI); i++ {
		mul *= slice[i]
	}

	return sum, mul
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

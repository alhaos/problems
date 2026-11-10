package main

import (
	"fmt"
	"slices"
)

func main() {

	var a [][]int

	b := make([]int, 3)

	for i := range len(b) {
		b[i] = i
	}

	a = append(a, slices.Clone(b))

	b[0], b[1] = b[1], b[0]

	a = append(a, b)

	fmt.Println(a)
}

package main

import (
	"fmt"

	"github.com/alhaos/problems/misc/brutForce"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}
	result := brutForce.CombinationSum(nums, 5)
	for _, nums := range result {
		for i, n := range nums {
			if i == 0 {
				fmt.Printf("%d", n)
			} else {
				fmt.Printf(", %d", n)
			}
		}
		fmt.Println()
	}
}

package acmp45

import (
	"sort"
)

func Solve(n int) int {

	if n == 0 {
		return 10
	}

	if n == 1 {
		return 1
	}

	var divisors []int

	for n > 1 {
		for i := 9; i >= 1; i-- {
			if n%i == 0 && i != 1 {
				n /= i
				divisors = append(divisors, i)
				break
			}
			if i == 1 {
				return -1
			}
		}
	}

	sort.Ints(divisors)

	result := digitsToNumber(divisors)

	return result
}

func digitsToNumber(digits []int) int {
	n := 0
	for _, d := range digits {
		n = n*10 + d
	}
	return n
}

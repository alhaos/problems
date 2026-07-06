package main

import (
	"context"
	"fmt"
)

// piDigits — бесконечный генератор десятичных цифр пи (включая первую "3").
// Основан на spigot-алгоритме Рабиновица-Вагона, но без фиксированного n:
// когда текущий блок цифр исчерпан, алгоритм пересчитывает всё с нуля
// с увеличенным размером блока. Это просто, но неоптимально по CPU для очень
// больших n (квадратичный рост работы) — для учебных целей и пары тысяч
// цифр более чем достаточно.
func piDigits(ctx context.Context) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		block := 100
		for {
			digits := computeDigits(block)
			for _, d := range digits {
				select {
				case out <- d:
				case <-ctx.Done():
					return
				}
			}
			block *= 2 // на следующей итерации считаем вдвое больше цифр
		}
	}()

	return out
}

// computeDigits вычисляет первые n десятичных цифр пи (вкл. "3") как срез.
func computeDigits(n int) []int {
	size := n*10/3 + 1
	A := make([]int, size)
	for i := range A {
		A[i] = 2
	}

	result := make([]int, 0, n)
	nines := 0
	predigit := -1

	for j := 0; j < n; j++ {
		carry := 0
		for i := size - 1; i >= 0; i-- {
			x := 10*A[i] + carry*(i+1)
			A[i] = x % (2*i + 1)
			carry = x / (2*i + 1)
		}
		A[0] = carry % 10
		q := carry / 10

		switch {
		case q == 9:
			nines++
		case q == 10:
			if predigit >= 0 {
				result = append(result, predigit+1)
			}
			for k := 0; k < nines; k++ {
				result = append(result, 0)
			}
			predigit = 0
			nines = 0
		default:
			if predigit >= 0 {
				result = append(result, predigit)
			}
			predigit = q
			for k := 0; k < nines; k++ {
				result = append(result, 9)
			}
			nines = 0
		}
	}
	if predigit >= 0 {
		result = append(result, predigit)
	}
	return result
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	const want = 100

	fmt.Print("π = ")
	i := 0
	for d := range piDigits(ctx) {
		if i == 1 {
			fmt.Print(".")
		}
		fmt.Print(d)
		i++
		if i >= want {
			cancel()
			break
		}
	}
	fmt.Println()
}

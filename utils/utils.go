package utils

import "fmt"

func NextCombination(x uint8) uint8 {
	c := x & -x             // шаг 1: младший установленный бит
	r := x + c              // шаг 2: сдвигаем правый блок единиц влево
	changed := r ^ x        // шаг 3: биты которые изменились
	shifted := changed >> 2 // шаг 4: убираем два лишних бита
	tail := shifted / c     // шаг 5: нормализуем к нулевой позиции
	next := tail | r        // шаг 6: собираем результат
	return next
}

func test() {
	for i := range 5 {
		switch {
		case i%4 == 0:
			fmt.Println(i)
			fallthrough
		case i%2 == 0:
			fmt.Println(i)
		}
	}
}

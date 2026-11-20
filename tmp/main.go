package main

import "fmt"

type s []int

func (h s) ChangeFirst()   { h[0] = 100 }
func (h s) Cut()           { h = h[:1] }
func (h *s) CutForReal()   { *h = (*h)[:1] }

func main() {
	a := s{1, 2, 3}

	a.ChangeFirst()
	fmt.Println(a)

	a.Cut()
	fmt.Println(a)

	a.CutForReal()
	fmt.Println(a)
}

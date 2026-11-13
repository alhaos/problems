package main

import (
	"fmt"
)

func main() {

	data := "H"

	for i := range data[1:] {
		fmt.Print(i)
	}
}

package main

import (
	"fmt"
	"strings"
)

const data = "12345"

func main() {
	split(data, nil)
}

func split(s string, parts []string) {

	if len(s) == 0 {
		fmt.Println(strings.Join(parts, ","))
		return
	}

	for i := 1; i <= len(s); i++ {
		split(s[i:], append(parts, s[:i]))
	}

}

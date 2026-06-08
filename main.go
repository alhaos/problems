package main

import "fmt"

func main() {
	res, err := safeDiv(1, 0)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
		return
	}
	fmt.Println("Hello, 世界", res)
}

func safeDiv(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("паника: %v", r)
		}
	}()
	return a / b, nil
}

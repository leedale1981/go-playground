package main

import (
	"errors"
	"fmt"
)

func main() {
	sum, err := sumNumbers(0, 2)
	fmt.Println(sum, err)

	sum, err = sumNumbers(10, 20)
	fmt.Println(sum, err)
}

func sumNumbers(first, second int) (int, error) {
	if first == 0 {
		return 0, errors.New("first cannot be zero")
	} else {
		return first + second, nil
	}
}

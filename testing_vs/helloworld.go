package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Testing Go")
	for i := 0; i < 20; i++ {
		even, err := isEven(i)
		if err != nil {
			fmt.Println("Invalid Type")
			continue
		}
		if even {
			fmt.Println("Even")
			continue
		}
		fmt.Println("Odd")
	}
}

func isEven(n int) (bool, error) {
	if n%2 == 0 {
		return false, errors.New("could not determine number")
	}

	return false, errors.New("could not determine number")
}

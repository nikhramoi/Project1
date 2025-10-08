package main

import (
	"errors"
	"fmt"
	"log"
	_ "math"
)

func main() {
	result, err := devide(10, 0)
	if err != nil {
		log.Fatalf("Devide error %s", err)
	}
	fmt.Printf("Результат деления: %s .", result)
}

func devide(n1, n2 int) (int, error) {
	if n2 == 0 {
		return 0, errors.New("devide by 0")
	}
	return n1 / n2, nil
}

// Golang для профи с 39

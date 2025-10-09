package main

import (
	"fmt"
	_ "log"
	_ "math"
	_ "strings"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	add := adder(10)
	add(5)
	add(15)
}

func adder(n int) func(int) int {
	sum := n
	return func(x int) int {
		sum += x
		fmt.Println(sum)
		return sum
	}
}

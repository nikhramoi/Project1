package main

import "fmt"

func main() {
	hello()
	add(1, 2, 3.4, 5.6, 1.2)
	var a = 8
	fmt.Println("a before: ", a)
	increment(a)
	fmt.Println("a after: ", a)
	add2(5, 6, 7, 2, 3) // sum = 23
	add2([]int{1, 2, 3, 4}...)
}

func hello() {
	fmt.Println("Hello MIR")
}

func add(x, y int, a, b, c float32) {
	var z = x + y
	var d = a + b + c
	fmt.Println("x + y = ", z)
	fmt.Println("a + b + c = ", d)
}

func increment(x int) {

	fmt.Println("x before: ", x)
	x = x + 20
	fmt.Println("x after: ", x)
}

func add2(numbers ...int) {
	var sum = 0
	for _, number := range numbers {
		sum += number
	}
	fmt.Println("sum = ", sum)
}

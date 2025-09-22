package main

import "fmt"

func main() {
	fmt.Println(sum(4)) // 10
	//fmt.Println(sum(5))   // 15
	//fmt.Println(sum(6))   // 21
	//fmt.Println(sum(100)) // 5050

	//Факториал числа
	fmt.Println(factorial(4)) // 24
	fmt.Println(factorial(5)) // 120
	fmt.Println(factorial(6)) // 720

	// Фибоначчи
	fmt.Println(fibbonachi(4)) // 3
	fmt.Println(fibbonachi(5)) // 5
	fmt.Println(fibbonachi(6)) // 8
}

func sum(n uint) uint {

	if n == 1 {
		return n
	}
	fmt.Println(n)
	return n + sum(n-1)
}

//Факториал числа
func factorial(n uint) uint {

	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

//Фибоначчи
func fibbonachi(n uint) uint {

	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibbonachi(n-1) + fibbonachi(n-2)
}

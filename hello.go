package main

import "fmt"

func main() {
	f := func(x, y int) int { return x + y }
	fmt.Println(f(3, 4)) // 7
	fmt.Println(f(6, 7)) // 13
	//Анонимная функция как аргумент функции
	action(10, 25, func(x int, y int) int { return x + y }) // 35
	action(5, 6, func(x int, y int) int { return x * y })   // 30
	//Анонимная функция как результат функции
	f2 := selectFn(1)
	fmt.Println(f2(2, 3)) // 5
	fmt.Println(f2(4, 5)) // 9
	fmt.Println(f2(7, 6)) // 13

}

//Анонимная функция как аргумент функции
func action(n1 int, n2 int, operation func(int, int) int) {
	result := operation(n1, n2)
	fmt.Println(result)
}

//Анонимная функция как результат функции
func selectFn(n int) func(int, int) int {
	if n == 1 {
		return func(x int, y int) int { return x + y }
	} else if n == 2 {
		return func(x int, y int) int { return x - y }
	} else {
		return func(x int, y int) int { return x * y }
	}
}

package main

import "fmt"

func main() {
	var f func(int, int) int = add //Присваиваем f функцию add()
	fmt.Println(f(3, 4))
	var x = f(4, 5) // 9
	fmt.Println(x)

	f2 := add2            //или так var f func(int, int) int = add
	fmt.Println(f2(3, 4)) // 7

	f = multiply         // теперь переменная f указывает на функцию multiply
	fmt.Println(f(3, 4)) // 12

	// f = display      // ошибка, так как функция display имеет тип func(string)

	var f1 func(string) = display // норм
	f1("hello")

	action(10, 25, add3)    // 35
	action(5, 6, multiply3) // 30
}

func add(x int, y int) int {
	return x + y
}

func add2(x int, y int) int     { return x + y }
func multiply(x int, y int) int { return x * y }

func display(message string) {
	fmt.Println(message)
}

//Функции как параметры других функций
func add3(x int, y int) int {

	return x + y
}
func multiply3(x int, y int) int {

	return x * y
}
func action(n1 int, n2 int, operation func(int, int) int) {

	result := operation(n1, n2)
	fmt.Println(result)
}

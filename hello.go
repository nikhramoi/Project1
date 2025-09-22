package main

import "fmt"

func main() {
	/* func имя_функции (список_параметров) тип_возвращаемого_значения {
	    выполняемые_операторы
	    return возвращаемое_значение
		} */
	var a = add(4, 5)  // 9
	var b = add(20, 6) // 26
	fmt.Println(a)
	fmt.Println(b)
	var age, name = add3(4, 5, "Tom", "Simpson")
	fmt.Println(age)  // 9
	fmt.Println(name) // Tom Simpson
}

func add(x, y int) int {
	return x + y
}

// Именованные возвращаемые результаты
func add2(x, y int) (z int) {
	z = x + y
	return
}

func add3(x, y int, firstName, lastName string) (int, string) {
	var z int = x + y
	var fullName = firstName + " " + lastName
	return z, fullName
}

/* func add4(x, y int, firstName, lastName string) (z int, fullName string) {
    z = x + y
    fullName = firstName + " " + lastName
    return
} */

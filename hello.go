package main

import "fmt"

func main() {
	a := 8
	b := 8
	if a < b {
		fmt.Println("a меньше b")
	} else if a > b {
		fmt.Println("a больше b")
	} else {
		fmt.Println("a равно b")
	}

	a1 := 9
	switch a1 {
	case 9:
		fmt.Println("a = 9")
		fallthrough
	case 8:
		fmt.Println("a = 8")
	case 7:
		fmt.Println("a = 7")
	case 6, 5, 4:
		fmt.Println("a = 6 или 5 или 4, но это не точно")
	default:
		fmt.Println("значение переменной a не определено")
	}

}

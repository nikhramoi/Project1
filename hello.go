package main

import "fmt"

func main() {
	var a int = 2
	fmt.Printf("a = %b\n", a) // a = 10
	a = a << 2                // сдвиг влево на два разряда
	fmt.Printf("a = %b\n", a) // a = 1000
}

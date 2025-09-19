package main

import "fmt"

func main() {
	for i := 1; i < 10; i++ {
		fmt.Println(i * i)
	}

	var i = 1
	for i < 10 {
		fmt.Println(i * i)
		i++
	}

	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			fmt.Print(i*j, "\t")
		}
		fmt.Println()
	}

	string1 := "Hello"
	for index, value := range string1 {
		fmt.Printf("Indexes: %d, Values: %c\n", index, value)
	}

	str := "Hello"
	for _, value := range str {
		fmt.Printf("%c ", value)
	}
	fmt.Printf("%c ", 10)

	// Операторы break и continue работают так же, как и в других языках программирования
	//break OuterLoop - выход из внешнего цикла
}

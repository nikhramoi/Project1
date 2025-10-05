package main

import (
	"fmt"
)

func main() {
	x := 10
	var p *int = &x
	fmt.Println("X=", x, "P=", p)
	*p++
	fmt.Println("X=", x, "P=", *p)

	ukaz := new(int)
	*ukaz = 100
	ukaz2 := new(int)
	ukaz2 = &x
	fmt.Println(*ukaz, ukaz, ukaz2, *ukaz2)
}

// Golang для профи с 39

package main

import "fmt"

func main() {
	var a int = 10 // простое присвоение
	b := 5

	a += b
	fmt.Println("a += b:", a) // a += b: 15

	a -= b
	fmt.Println("a -= b:", a) // a -= b: 10

	a *= 10
	fmt.Println("a *= 10:", a) // a *= 10: 100

	a /= b
	fmt.Println("a /= b:", a) // a /= b: 20

	a %= 15
	fmt.Println("a %= 15:", a) // a %= 15: 5

	a <<= 2
	fmt.Println("a <<= 2:", a) // a <<= 2: 20

	a >>= 1                    // 10100 >>= 1 = 1010
	fmt.Println("a >>= 1:", a) // a >>= 1: 10

	a &= 8                    // 10100 &=1 = 1000
	fmt.Println("a &= 8:", a) // a &= 8: 8

	a ^= 10                    // 1000 ^= 1010 = 0010
	fmt.Println("a ^= 10:", a) // a ^= 10: 2

	a |= 5                    // 010 |= 101 = 111
	fmt.Println("a |= 5:", a) // a |= 5: 7
}

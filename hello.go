package main

import (
	"fmt"
	_ "log"
	_ "math"
	_ "strings"
)

type Day int

const (
	_ Day = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func main() {
	fmt.Println(isWeekend(Wednesday))
	fmt.Println(isWeekend(Sunday))
}

func isWeekend(i Day) bool {
	if i == 6 || i == 7 {
		return true
	}
	return false
}

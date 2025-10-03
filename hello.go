package main

import "fmt"

func main() {
	m := map[string]int{
		"hello": 5,
		"world": 0,
	}
	var ebok bool
	v, ok := m["hello"]
	fmt.Println(v, ok)
	v, ok = m["world"]
	fmt.Println(v, ok)
	v, ebok = m["goodbye"]
	fmt.Println(v, ebok)
}

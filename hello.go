package main

import (
	"errors"
	"fmt"
	"log"
	_ "math"
	"strings"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	message, err := UserProfileToString("Никита", -5)
	if err != nil {
		log.Fatalf("Devide error %s", err)
		//fmt.Errorf("Devide error %s", err)
	}
	fmt.Println(message)
}

func UserProfileToString(name string, age int) (string, error) {

	if name == "" {
		return "", errors.New("empty name")
	}
	if age < 0 {
		return "", errors.New("negative age")
	}
	if strings.TrimSpace(name) == "" {
		return "", errors.New("name cannot contain only spaces")
	}
	return fmt.Sprintf("Имя человека: %s, возраст: %d.", name, age), nil
}

func calculate(var1, var2 float64, operation string) (float64, error) {
	if operation == "add" {
		return var1 + var2, nil
	}
	if operation == "subtract" {
		return var1 - var2, nil
	}

	if operation == "multiply" {
		return var1 * var2, nil
	}

	if operation == "divide" {
		if var2 == 0 {
			return 0, errors.New("division by zero")
		}
		return var1 / var2, nil
	}

	return 0, errors.New("unknown operation")
}

// Golang для профи с 39

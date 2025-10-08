package main

import (
	"fmt"
	_ "math"
	"reflect"
)

type Secret struct {
	Username string
	Password string
}
type Record struct {
	Field1 string
	Field2 float64
	Field3 Secret
}

func main() {
	A := Record{"String value", -12.123, Secret{"Mihalis", "Tsoukalos"}}
	r := reflect.ValueOf(A)
	fmt.Println("String value:", r.String())
	iType := r.Type()
	fmt.Printf("i Type: %s\n", iType)
	fmt.Printf("The %d fields of %s are\n", r.NumField(), iType)
	for i := 0; i < r.NumField(); i++ {
		fmt.Printf("\t%s ", iType.Field(i).Name)
		fmt.Printf("\twith type: %s ", r.Field(i).Type())
		fmt.Printf("\tand value _%v_\n", r.Field(i).Interface())
		// проверьте, есть ли другие структуры, встроенные в запись
		k := reflect.TypeOf(r.Field(i).Interface()).Kind()
		// нужно преобразование в строку, чтобы можно было сравнить
		if k.String() == "struct" {
			//Чтобы проверить тип данных переменной с помощью строки, нам нужно сначала преобразовать тип данных в string.
			fmt.Println(r.Field(i).Type())
		}
		// то же, что и раньше, но с использованием внутреннего значения
		if k == reflect.Struct {
			//Во время проверки вы также можете использовать внутреннее представление типа данных. Однако смысла в этом меньше, чем в использовании значения string.
			fmt.Println(r.Field(i).Type())
		}
	}
}

// Golang для профи с 39

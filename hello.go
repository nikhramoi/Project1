package main

import "fmt"

func main() {
	var tom person = person{"Tom", 23} // переменная структуры person
	var alice person = person{age: 23, name: "Alice"}
	var tomohawk = person{name: "Tomohawk", age: 24}
	bob := person{name: "Bob", age: 31}
	fmt.Println(tom) // { 0}
	fmt.Println(alice)
	fmt.Println(tomohawk)
	fmt.Println(bob)
	//Анонимные структуры
	tom1.name = "Tom"
	tom1.age = 23
	fmt.Println(tom1.name) // Tom
	fmt.Println(tom1.age)  // 23
	tom2 := struct {       // объявление переменной анонимной структуры
		name string
		age  int
	}{ // инициализация переменной
		name: "Tom",
		age:  41,
	}
	fmt.Println(tom2.name) // Tom
	fmt.Println(tom2.age)  // 41
	victor := person{"Victor", 30}
	fmt.Println(victor) // {Tom 41}

	// анонимная структура с new
	bob1 := new(struct {
		name, company string
		age           int
	})
	bob1.name = "Bob"
	bob1.company = "Google"
	bob1.age = 46
	fmt.Println(bob1)
	// копирование структуры
	tomas := tom // копирование структуры tom
	tomas.age = 18

	// изменение одной структуры никак не влияет на другую
	fmt.Println(tom.age)   // 41
	fmt.Println(tomas.age) // 18

	//Передача структуры в функцию
	// в функцию передается адрес структуры tom
	increment_age(&tom)

	// изменение структуры внутри функции по указателю приведет к изменению оригинальной структуры
	fmt.Println(tom) // {Tom 42}

	//При сравнении структур они считаются равными, если их поля и соответствующие значения одинаковы:
}

type person struct {
	name string
	age  int
}

// Анонимные структуры
var tom1 struct {
	name string
	age  int
}

//Анонимные поля структур

type person2 struct {
	string
	int
}

//Передача структур в функцию
func increment_age(user *person) {
	user.age += 1
	fmt.Println(*user) // {Tom 42}
}

package main

import "fmt"

func main() {
	fn := outer() // fn = inner, так как функция outer возвращает функцию inner
	// вызываем внутреннюю функцию inner
	fn() // 6
	fn() // 7
	fn() // 8
	//Пример 2
	fn2 := multiply(5)
	result1 := fn2(6)    // 30
	fmt.Println(result1) // 30

	result2 := fn2(5)    // 25
	fmt.Println(result2) // 25

	result3 := multiply(7)(6) // 42
	fmt.Println(result3)      // 42
}

func outer() func() { // внешняя функция
	var n int = 5     // некоторая переменная - лексическое окружение функции inner
	inner := func() { // вложенная функция
		// действия с переменной n
		n += 1
		fmt.Println(n)
	}
	return inner
}

/* Пример с анонимной функцией
    Function outer(){
    var n = 5;
    return (){
        n++;
        print(n);
    };
} */

//Пример 2
func multiply(n int) func(int) int {

	return func(m int) int { return n * m }
}

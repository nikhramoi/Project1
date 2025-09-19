package main

import "fmt"

func main() {
	var numbers [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)
	var numbers2 [5]int = [5]int{1, 2}
	fmt.Println(numbers2) // [1 2 0 0 0]
	// numbers := [5]int{1,2,3,4,5} Краткая форма объявления массива
	var numbers3 = [...]int{1, 2, 3, 4, 5} // длина массива 5
	numbers4 := [...]int{1, 2, 3}          // длина массива 3
	numbers4[2] = 100                      // изменение значения элемента
	fmt.Println(numbers3)                  // [1 2 3 4 5]
	fmt.Println(numbers4)                  // [1 2 3]
	numbers5 := [3][2]int{
		{1, 2},
		{4, 5},
		{7, 8},
	}
	numbers5[0] = [2]int{5, 5}
	numbers5[1] = [2]int{6, 6}
	numbers5[2] = [2]int{7, 78}
	numbers5[0][0] = 1
	numbers5[0][1] = 2

	numbers5[1][0] = 4
	numbers5[1][1] = 5
	numbers5[2][0] = 7
	numbers5[2][1] = 8
	fmt.Println(numbers5)

	nums1 := [4]int{3, 4, 5, 6}
	nums2 := nums1 // копирование массива nums1 в nums2

	nums2[1] = 11 // меняем значение в nums2, nums1 не меняется

	fmt.Println("nums1:", nums1) // nums1: [3 4 5 6]

	fmt.Println("nums2:", nums2) // nums2: [3 11 5 6]

}

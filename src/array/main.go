package main

import (
	"fmt"
)

func changeElement(arr *[3]int) {
	// arr[0] = 1000
	(*arr)[0] = 1000
}

func main() {
	var arr1 [5]int
	arr1[2] = 10
	fmt.Println(arr1)

	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)

	arr3 := [...]int{4, 5, 6, 7, 8}
	fmt.Println(arr3)

	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}

	for i, num := range arr3 {
		fmt.Println(i, num)
	}

	for _, num := range arr3 {
		fmt.Println(num)
	}

	for i := range arr3 {
		fmt.Println(i)
	}

	var arr4 = [3][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 0, 11, 12},
	}
	fmt.Println(arr4)

	arr5 := [3]int{1, 2, 3}
	changeElement(&arr5)
	fmt.Println(arr5)
}

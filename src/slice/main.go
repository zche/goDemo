package main

import "fmt"

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := arr[2:6]
	fmt.Println(s1, len(s1), cap(s1), arr)
	// s1 = append(s1, 10)
	// fmt.Println(s1, len(s1), cap(s1), arr)
	// s1 = append(s1, 20, 30, 40)
	// fmt.Println(s1, len(s1), cap(s1), arr)
	s1 = append(s1, 10, 20, 30, 40, 50, 60)
	fmt.Println(s1, len(s1), cap(s1), arr)

	var s2 []int
	fmt.Println(s2, s2 == nil, len(s2), cap(s2))

	s6 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s6 = s6[:len(s6)-1]
	fmt.Println(s6)
}

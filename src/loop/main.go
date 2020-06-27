package main

import (
	"fmt"
)

func main() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	a := 15
	b := 10
	for a >= b {
		b++
		if b == 14 {
			break
		}
		if b == 11 {
			continue
		}
		fmt.Println(b)
	}
}

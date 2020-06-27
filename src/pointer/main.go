package main

import (
	"fmt"
)

func pointerTest() {
	var p = 20
	var ip *int
	ip = &p
	*ip = 30
	fmt.Println(p)
}

func funcValRef(a int) int {
	a = 1000
	return a
}

func funcValRefPtr(a *int) {
	*a = 1000
}

func main() {
	pointerTest()
	p := 12
	funcValRef(p)
	fmt.Println(p)
	fmt.Println("===================")
	funcValRefPtr(&p)
	fmt.Println(p)
}

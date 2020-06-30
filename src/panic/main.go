package main

import (
	"fmt"
)

func tryPanic() {
	// defer func() {
	// 	if e := recover(); e != nil {
	// 		fmt.Printf("catch panic in recover function: %v\n", e)
	// 	}
	// }()
	fmt.Println("first line in tryPanic func")
	panic("call a panic")
}

func protect(g panicFunc) {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("catch panic in recover function: %v\n", e)
		}
	}()
	g()
}

type panicFunc func()

func main() {
	fmt.Println("start call tryPanic")
	// tryPanic()
	protect(tryPanic)
	fmt.Println("end call tryPanic")
}

package main

import (
	"fmt"
	"goTesting/calc"
)

func main() {
	for i:=0;i<=100;i++ {
		fmt.Printf("%d is even? %v\n",i, calc.Even(i))
	}
}

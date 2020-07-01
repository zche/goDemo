package main

import (
	"fmt"
	"time"
)

func printGroutine(i int) {
	fmt.Println("I am from goroutine %d",i)
}

func main() {
	var a [12]int
	for i:=1; i< 10; i++ {
		go func(i int) {
			a[i]++
			//fmt.Printf("I am from goroutine %d\n",i)
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println(a)
}

package main

import (
	"fmt"
	"runtime"
	"strconv"
)

func formatVal() {
	var a = 1
	var b = "hello"
	fmt.Printf("%T %T\n", a, b)
	var c int
	var d string
	fmt.Printf("%d %q\n", c, d)
	fmt.Println(a)
}
func systemVal() {
	cpuArch := runtime.GOARCH
	intSize := strconv.IntSize
	fmt.Println(cpuArch, intSize)
}
func pointVal() {
	var p = 1
	ptr := &p
	fmt.Printf("%T %T\n", p, ptr)
}
func enumDemo() {
	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)
}
func typeDefAndAlias() {
	type MyInt1 int
	type MyInt2 = int
	var i = 1
	var i1 = MyInt1(i)
	fmt.Println(i1)
	var i2 = i
	fmt.Println(i2)
}

func main() {
	formatVal()
	systemVal()
	pointVal()
	enumDemo()
	typeDefAndAlias()
}

package main

import (
	"fmt"
)

func testIfelse() {
	var a = 20
	if a < 20 {
		fmt.Println("小于20")
	} else if a == 20 {
		fmt.Println("等于20")
	} else {
		fmt.Println("大于20")
	}
}

func testSwitch() {
	var a = 20
	switch a {
	case 20:
		fmt.Println("a=20")
		fallthrough
	case 21:
		fmt.Println("a=21")
	case 22:
		fmt.Println("a=22")
	default:
		fmt.Println("unkown")

	}
}

func main() {
	testIfelse()
	fmt.Println("---------")
	testSwitch()
}

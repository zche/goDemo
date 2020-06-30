package main

import (
	"defer/file"
	"fmt"
)

func main() {
	if content, err := file.ReadFile("abc1.txt"); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(content)
	}
}

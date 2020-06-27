package main

import (
	"fmt"
	"math"
	"os"
)

func operate(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a / b, nil
	case "/":
		return a + b, nil
	default:
		return 0, fmt.Errorf("not support operate: %s", op)
		// panic(fmt.Sprintf("not support operate: %s", op))
	}
}

func swap1(a, b int) (int, int) {
	return b, a
}

func swap2(a, b int) (x, y int) {
	x, y = b, a
	return
}

func compute(op func(int, int) int, a, b int) int {
	return op(a, b)
}

func powInt(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

type greeting func(name string) string

func say(g greeting, word string) {
	fmt.Println(g(word))
}

func english(name string) string {
	return "Hello," + name
}

func chinese(name string) string {
	return "你好," + name
}

func sum(nums ...int) int {
	s := 0
	for i := 0; i < len(nums); i++ {
		s += nums[i]
	}
	return s
}

func main() {
	// var res, err = operate(2, 3, "+")
	if res, err := operate(2, 3, "ee"); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(res)
	}
	// fmt.Println(res)
	fmt.Println(swap1(6, 8))
	x, y := swap2(6, 8)
	fmt.Println(x, y)
	// file, err := os.Open("ab.txt")
	if file, err := os.Open("ab.txt"); err != nil {
		fmt.Println("打开文件有误")
	} else {
		fmt.Println(file)
	}
	fmt.Println(compute(powInt, 3, 4))
	say(english, "word!")
	say(chinese, "word!")
	fmt.Println(sum(1, 2, 3))
}

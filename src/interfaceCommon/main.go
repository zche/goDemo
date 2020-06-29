package main

import (
	"fmt"
	"interfaceCommon/example"
)

type Appender interface {
	Append(int)
}

type Lener interface {
	Len() int
}

type List []int

func (ls List) Len() int {
	return len(ls)
}

func (ls *List) Append(val int) {
	*ls = append(*ls, val)
}

func Counter(a Appender, start, end int) {
	for i := start; i <= end; i++ {
		a.Append(i)
	}
}

func IsLarge(l Lener) bool {
	return l.Len() > 50
}

func main() {
	// var lst List
	// Counter(&lst, 1, 10)
	// IsLarge(lst)

	pList := new(List)
	Counter(pList, 1, 10)
	fmt.Println(pList)
	// 指针类型的参数是可以自动解引用的
	IsLarge(pList)

	course := new(example.Course)
	course.Title = "go"
	course.SubTitle = "golang实战"

	fmt.Println(course)

	data := []int{23, 50, 79, 11, 19, 60}
	ia := example.IntArray(data)
	example.Sort(ia)
	fmt.Println(ia)
}

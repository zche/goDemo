package example

import "sort"

func Sort(data sort.Interface) {
	for i := 1; i < data.Len(); i++ {
		for j := 0; j < data.Len()-i; j++ {
			if data.Less(j+1, j) {
				data.Swap(j, j+1)
			}
		}
	}
}

type IntArray []int

func (ia IntArray) Len() int {
	return len(ia)
}

func (ia IntArray) Less(i, j int) bool {
	return ia[i] < ia[j]
}

func (ia IntArray) Swap(i, j int) {
	ia[i], ia[j] = ia[j], ia[i]
}

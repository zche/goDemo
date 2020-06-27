package main

import "fmt"

func main() {
	var m map[string]int
	fmt.Println(m, m == nil)
	var m1 map[string]int
	m1 = make(map[string]int)
	m1["a"] = 10
	fmt.Println(m1, m1 == nil)
	m2 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
	}
	fmt.Println(m2, m2["b"])
	if val, ok := m2["e"]; ok {
		fmt.Printf("key e in map and value is %d", val)
	}
	for key := range m2 {
		fmt.Printf("key=%s val=%d\n", key, m2[key])
	}

	for key, val := range m2 {
		fmt.Printf("key=%s val=%d\n", key, val)
	}
	delete(m2, "a")
}

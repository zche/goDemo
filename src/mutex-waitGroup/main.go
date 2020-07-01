package main

import (
	"fmt"
	"sync"
)

type safeInt struct {
	a int
	lock sync.Mutex
}

func (si *safeInt) increase() {
	si.lock.Lock()
	defer si.lock.Unlock()
	si.a++
}

func (si *safeInt) get() int {
	si.lock.Lock()
	defer si.lock.Unlock()
	return  si.a
}

func main() {
	//var a int
	//a++
	//go func() {
	//	a++
	//}()
	var si safeInt
	si.increase()
	var wg sync.WaitGroup
	for i:=0;i<50000;i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			si.increase()
		}()
	}
	//time.Sleep(time.Millisecond)
	wg.Wait()
	fmt.Println(si.get())
}


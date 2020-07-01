package main

import (
	"fmt"
	"time"
)

func getData(i int,c chan int) {
	//for {
	//	//n := <- c
	//	//fmt.Println(n)
	//	if n, ok := <- c; ok {
	//		fmt.Printf("get data %d from channel %d\n",n,i)
	//	} else {
	//		break
	//	}
	//}
	for n := range c {
		fmt.Printf("get data %d from channel %d\n",n,i)
	}
}

func createChannel(i int) chan<- int {
	c := make(chan int)
	go getData(i,c)
	return  c
}

func channelDemo() {
	//c := make(chan int)
	//go func() {
	//	for {
	//		n := <- c
	//		fmt.Println(n)
	//	}
	//}()
	//c <- 10 // 发送者
	//c <- 20
	var channels [10]chan<- int
	for i:=0;i<10;i++ {
		channels[i] = createChannel(i)
	}

	//for i:=0;i<10;i++ {
	//	close(channels[i])
	//}

	for i:=0;i<10;i++ {
		channels[i] <- i
	}

	time.Sleep(time.Millisecond)
}

func main() {
	channelDemo()
}

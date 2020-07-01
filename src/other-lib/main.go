package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"regexp"
	"time"
)

func timeDemo() {
	now := time.Now()
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(now.Year(),now.Month(),now.Day())
}

func randDemo() {
	fmt.Println(rand.Intn(100))
}

type BaseResponse struct {
	Code int
	Data ResponseData
}

type ResponseData struct {
	Name string `json:"code"`
	Age int `json:"age"`
}

func jsonDemo() {
	br := BaseResponse{
		Code: 1,
		Data:ResponseData{
			Name: "check",
			Age: 27,
		},
	}
	jsonBytes, _ := json.Marshal(br)
	fmt.Println(string(jsonBytes))

	var br2 BaseResponse
	_ = json.Unmarshal(jsonBytes,&br2)
	fmt.Println(br2)
}

func regexDemo() {
	input := "My email is lala@outlook.com check@gmail.com yyy@test.com.cn"
	//exp, _ := regexp.Compile("check@gmail.com")
	exp, _ := regexp.Compile("[a-zA-Z0-9]+@[a-zA-Z0-9]+.[a-zA-Z0-9]+.*[a-zA-Z0-9]*")
	fmt.Println(exp.FindString(input))
	fmt.Println(exp.FindAllString(input, -1))
}

// init函数式最先执行的
func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	timeDemo()

	randDemo()

	jsonDemo()

	regexDemo()
}

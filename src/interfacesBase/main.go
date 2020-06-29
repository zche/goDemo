package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type Speaker interface {
	Say(string)
}

func SpeakSomething(speaker Speaker) {
	speaker.Say("abcdeft,lala....")
}

type Person struct {
	name string
}

func (p *Person) Say(message string) {
	fmt.Println(p.name, ":", message)
}

type SpeakWriter struct {
	w io.Writer
}

func (w *SpeakWriter) Say(message string) {
	io.WriteString(w.w, message)
}

type FileWriter struct {
	filename string
}

func (fw *FileWriter) Write(p []byte) (n int, err error) {
	err = ioutil.WriteFile(fw.filename, p, 0644)
	n = 0
	return
}

type Phone interface {
	Call(string)
}

type Camera interface {
	Take() string
}

type SmartPhone interface {
	Phone
	Camera
	Downlad(string) string
}

type MiPhone struct {
	Logo string
}

func (mp *MiPhone) Call(phone string) {
	fmt.Println("call the phone:", phone)
}

func (mp *MiPhone) Take() string {
	return "logo.png"
}

func (mp *MiPhone) Downlad(url string) string {
	return fmt.Sprintf("visit the url:%s", url)
}

func ListSmartPhoneFeature(sp SmartPhone) {
	if v, ok := sp.(*MiPhone); ok {
		v.Call("10010")
		fmt.Println("take photos:", v.Take())
		fmt.Println("download url is:", v.Downlad("www.baidu.com"))
	} else {
		fmt.Println("Not MiPhone Pointer type")
	}
}

type Any interface{}

func GetAnyType(any interface{}) {
	switch t := any.(type) {
	case int:
		fmt.Println("any is int type")
	case string:
		fmt.Println("any is string type")
	case *MiPhone:
		fmt.Println("any is MiPhone point type")
	default:
		fmt.Printf("unexpected type %T\n", t)
	}
}

type Queue []interface{}

func (q *Queue) Push(n interface{}) {
	*q = append(*q, n)
}

func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func main() {
	check := new(Person)
	check.name = "check"
	SpeakSomething(check)

	console := new(SpeakWriter)
	console.w = os.Stdout
	SpeakSomething(console)

	fileWriter := new(FileWriter)
	fileWriter.filename = "abc.txt"
	fmt.Println("=======================")
	fileSpeakWriter := new(SpeakWriter)
	fileSpeakWriter.w = fileWriter
	SpeakSomething(fileSpeakWriter)

	var sp SmartPhone
	mp := new(MiPhone)
	mp.Logo = "xiaomi"
	sp = mp
	ListSmartPhoneFeature(sp)

	q := Queue{1, 2, 3}
	q.Push(4)
	q.Push(5)
	fmt.Println(q)
	q.Pop()
	fmt.Println(q)

	q.Push("ab")
	fmt.Println(q)
}

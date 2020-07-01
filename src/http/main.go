package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"path"
)

func sendHttpRequestDemo() {
	//User-Agent: Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1
	req, err := http.NewRequest(http.MethodGet,"http://www.baidu.com",nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("User-Agent","Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1")
	resp, err := http.DefaultClient.Do(req)
	//resp, err :=http.Get("http://www.baidu.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err!= nil {
		panic(err)
	}
	fmt.Println(string(body))
}

type User struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func httpServerDemo() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer,"hello golang \n")
	})

	http.HandleFunc("/json", func(writer http.ResponseWriter, request *http.Request) {
		user := User{
			Name: "check",
			Age: 20,
		}
		userJson, err := json.Marshal(user)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		writer.Header().Set("Content-Type","application/json")
		writer.Write(userJson)
	})

	http.HandleFunc("/image", func(writer http.ResponseWriter, request *http.Request) {
		image := path.Join("images","golang.png")
		http.ServeFile(writer,request,image)
	})

	http.HandleFunc("/html", func(writer http.ResponseWriter, request *http.Request) {
		user := User{
			Name: "check",
			Age: 30,
		}
		htmlFile := path.Join("templates", "index.html")
		temp, err := template.ParseFiles(htmlFile)
		if err != nil {
			http.Error(writer,err.Error(),http.StatusInternalServerError)
			return
		}
		if err := temp.Execute(writer,user); err != nil {
			http.Error(writer,err.Error(),http.StatusInternalServerError)
		}
	})

	log.Fatal(http.ListenAndServe(":1234",nil))
}

func main() {
	//sendHttpRequestDemo()
	httpServerDemo()
}

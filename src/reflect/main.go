package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type User struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

func (u User) String(user User) {
	fmt.Printf("Name=%s,Age=%d\n",user.Name,user.Age)
}

func main() {
	//reflectBase()
	//reflectStruct()
	reflectChangeFieldValue()
	//reflectDynamicCallMethod()
	//reflectStructTag()
}

func reflectBase() {
	u := User{"check",33}
	t := reflect.TypeOf(u)
	fmt.Println(t)

	v := reflect.ValueOf(u)
	fmt.Println(v)

	fmt.Printf("%T,%v\n",u,u)

	fmt.Println("=================")

	// reflect.Value 转成原始数据
	u1 := v.Interface().(User)
	fmt.Println(u1,reflect.TypeOf(u1),v.Type())

	fmt.Println(v.Type().Kind())
}

func reflectStruct() {
	u := User{"张三",20}
	t := reflect.TypeOf(u)
	for i:=0;i<t.NumField() ;i++ {
		f := t.Field(i)
		fmt.Printf("fieldName: %s\n", f.Name)
	}


	for i:=0;i<t.NumMethod() ;i++ {
		f := t.Method(i)
		fmt.Printf("fieldName: %s\n", f.Name)
	}
}

func reflectChangeFieldValue() {
	x := 2
	v := reflect.ValueOf(&x)
	v.Elem().SetInt(100)
	fmt.Println(x)

	u := User{"mac", 3}
	fmt.Println(u)
	vMac := reflect.ValueOf(&u).Elem()
	vName := vMac.FieldByName("Name")
	vName.SetString("iphone")
	fmt.Println(u)
}

func reflectDynamicCallMethod() {
	u := User{"check",33}
	v := reflect.ValueOf(u)
	med := v.MethodByName("String")
	if med.IsValid() {
		args := []reflect.Value{v}
		fmt.Println(med.Call(args))
	}
}

func reflectStructTag() {
	var u User
	h := `{"name":"check","age":29}`
	err := json.Unmarshal([]byte(h), &u)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(u,u.Name,u.Age)
	}

	t := reflect.TypeOf(u)
	for i:=0; i<t.NumField(); i++ {
		f := t.Field(i)
		fmt.Println(f.Tag, f.Tag.Get("json"))
	}
}
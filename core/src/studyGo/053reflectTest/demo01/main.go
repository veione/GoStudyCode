package main

import (
	"fmt"
	"reflect"
)

type BaseInfo struct {
	Height float32 `json:"height"`
	Weight float32 `json:"weight"`
}

type Student struct {
	BaseInfo
	Name string `json:"name"` // 是 ` ` (tab键上的~按键) ,不是 ' '
	Sex  string `json:"sex"`
	Age  int    `json:"age"`
}

func testReflect(in interface{}) {
	t := reflect.TypeOf(in)
	v := reflect.ValueOf(in)

	fmt.Printf("type:%v  value:%v\n", t, v)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()
		fmt.Printf("name:%s type:%v value:%v tag:%v\n", field.Name, field.Type, value, field.Tag)
	}
}

func main() {
	PersonA := Student{
		Name: "张三",
		Sex:  "男",
		Age:  24,
	}
	PersonA.Height = 178.2
	PersonA.Weight = 58.2

	testReflect(PersonA)
}

package main

import (
	"fmt"
	"reflect"
)

//使用反射来遍历结构体的字段，调用结构体的方法，修改结构体字段的值，并获取结构体标签的值

//定义结构体
type Student struct {
	Name string	`json:"name"`  // 是 ` ` (tab键上的~按键) ,不是 ' '
	Sex string `json:"sex"`
	Age int `json:"age"`
	Sal float64 `json:"sal"`
}

func (s Student) GetName() string  {  //第0个方法
	fmt.Println("该结构体Name字段值为：",s.Name)
	return s.Name
}

func (s *Student) Set(newName string,newAge int,newSal float64){  //第2个方法
	s.Name = newName
	s.Age = newAge
	s.Sal = newSal
	s.Print()
}

func (s Student) Print()   { //第1个方法
	fmt.Println("调用 Print 函数输出结构体：",s)
}

//反射获取结构体字段、方法，并调用
func testReflect(b interface{})  {
	rVal := reflect.ValueOf(b).Elem()
	rType := reflect.TypeOf(b).Elem()

	//判断是否是结构体在进行下一步操作
	if rType.Kind() != reflect.Struct{
		fmt.Println("该类型不是结构体。所以无法获取字段及其方法。")
	}

	//获取字段数量
	numField := rVal.NumField()
	fmt.Printf("该结构体有%d个字段\n",numField)
	//遍历字段
	for i := 0; i < numField; i++ {
		//获取字段值、标签值
		rFieldTag := rType.Field(i).Tag.Get("json")
		if rFieldTag != "" {
			fmt.Printf("结构体第 %v 个字段值为：%v ," +
				"Tag‘json’名为：%v\n",i,rVal.Field(i),rFieldTag)
		}
	}

	//获取方法数量
	numMethod := rVal.NumMethod()   //用指针可以获取到指针接收的方法
	fmt.Printf("该结构体有%d个方法\n",numMethod)

	//调用方法（方法顺序 按照ACSII码排序）
	rVal.Method(0).Call(nil)
	rVal.Method(1).Call(nil)

	//参数也需要以 Value 的切片 传入
	params  := make([]reflect.Value ,3)
	params[0] = reflect.ValueOf("hhhh")
	params[1] = reflect.ValueOf(28)
	params[2] = reflect.ValueOf(99.9)
	rVal.Method(2).Call(params)

	rVal.Method(1).Call(nil)
}

func main() {
	stu := Student{
		Name: "莉莉安",
		Sex: "f",
		Age: 19,
		Sal: 98.5,
	}

	//调用编写的函数并输出
	testReflect(&stu)
	fmt.Println("主函数输出结构体 Student ：",stu)
}
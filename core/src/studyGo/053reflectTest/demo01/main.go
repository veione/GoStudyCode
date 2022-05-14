package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"strings"
)

type BaseInfo struct {
	Height float32 `json:"height"`
	Weight float32 `json:"weight"`
	Flag   string  `json:"flag"`
}

type Student struct {
	BaseInfo
	Name string `json:"name"` // 是 ` ` (tab键上的~按键) ,不是 ' '
	Sex  string `json:"sex"`
	Age  int    `json:"age"`
}

type jsonkeyVal struct {
	Key   string
	Value string
}

func SortMapToURL(mReq map[string]interface{}, join1, join2 string) string {
	//STEP 1, 对key进行升序排序.
	sorted_keys := make([]string, 0)
	for k, _ := range mReq {
		sorted_keys = append(sorted_keys, k)
	}
	sort.Strings(sorted_keys)

	//STEP2, 对key=value的键值对用join1,join2连接起来，略过空值  如keyjoin1valuejoin2 a=1&b=2
	var signStrings string
	for i, k := range sorted_keys {
		value := fmt.Sprintf("%v", mReq[k])
		if value != "" {
			if i != (len(sorted_keys) - 1) {
				signStrings = signStrings + k + join1 + value + join2
			} else {
				signStrings = signStrings + k + join1 + value //最后一个不加此符号
			}
		} else {
			signStrings = signStrings + k + join1 + join2
		}
	}
	return strings.Trim(signStrings, join2)
}

func testReflect(in interface{}) {
	mReq := make(map[string]interface{})
	var readFieldToKV func(fv *reflect.Value, ft *reflect.StructField)
	readFieldToKV = func(fv *reflect.Value, ft *reflect.StructField) {
		if ft.Anonymous {
			numField := fv.NumField()
			for i := 0; i < numField; i++ {
				ffv := fv.Field(i)
				fft := ft.Type.Field(i)
				readFieldToKV(&ffv, &fft)
			}
			return
		}
		if !fv.CanInterface() {
			return
		}
		key := ft.Tag.Get("json")
		if key != "flag" {
			mReq[key] = fv.Interface()
		}
	}
	t := reflect.TypeOf(in)
	v := reflect.ValueOf(in)

	fmt.Printf("type:%v  value:%v\n", t, v)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		readFieldToKV(&value, &field)
	}
	data := SortMapToURL(mReq, "", "")
	fmt.Printf("排序后：%v", data)
}

func testMarshal(st *Student) {
	b, err := json.Marshal(st)
	if err != nil {
		fmt.Printf("json Marshal error %v \n", err)
	} else {
		fmt.Printf("json: %v \n", string(b))
	}
	newSt := &Student{}
	err = json.Unmarshal(b, newSt)
	if err != nil {
		fmt.Printf("json UnMarshal error %v \n", err)
	} else {
		fmt.Printf("student: %v ", *newSt)

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
	PersonA.Flag = "wtq"

	testReflect(PersonA)
	//testMarshal(&PersonA)
}

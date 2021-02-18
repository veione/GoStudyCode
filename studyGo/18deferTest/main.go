package main

import (
	"fmt"
	"runtime"
)

func GetFunctionName() string {

	pc := make([]uintptr, 10)
	runtime.Callers(2,pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}

func testwhendefertorun() (rc int) {

	defer func() {
		fmt.Printf("i am defer , rc: %d, at: %s\n", rc, GetFunctionName())
		rc = 2


	}()

	//return 1
	return func() int {

		fmt.Printf("i am return, rc: %d, at: %s\n", rc, GetFunctionName())

		return 1
	}()

}
func main()  {

	x := testwhendefertorun()
	fmt.Printf("i am main, rc: %d, at: %s\n", x, GetFunctionName())
}



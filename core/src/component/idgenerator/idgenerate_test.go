package idgenerator

import (
	"fmt"
	"testing"
)

func TestFetchUUID(t *testing.T) {
	for i := 0; i <= 100; i++ {
		uid := FetchUUID()
		fmt.Printf("uid:%v len:%v \n", uid, len(uid))
	}
}

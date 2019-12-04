package basic

import (
	"fmt"
	"runtime/debug"
	"testing"
)

func TestStack(t *testing.T) {
	defer func() {
		fmt.Println(string(debug.Stack()))
	}()
	call1()
}

func call1()  {
	call2()
}

func call2()  {
	call3()
}

func call3()  {
	panic("errorrrrrrrrr")
}

package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	defer func() {
		r := recover()
		fmt.Println(r)
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

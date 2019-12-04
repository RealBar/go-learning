package basic

import (
	"fmt"
	"testing"
	"time"
)

type T struct {
	A int
	B int
	C int
}
var b int
//go:noinline
func TestStruct(t *testing.T) {
	times := 1000_000_000
	a := &T{A:13,B:14,C:55}
	start := time.Now()
	for i:=0;i < times;i++{
		b = a.A
	}
	l1 := time.Since(start)
	fmt.Println(l1)
	start = time.Now()
	for i:=0;i < times;i++{
		b = a.B
	}
	l1 = time.Since(start)
	fmt.Println(l1)
	start = time.Now()
	for i:=0;i < times;i++{
		b = a.C
	}
	l1 = time.Since(start)
	fmt.Println(l1)
	fmt.Println(b)
}

package basic

import (
	"fmt"
	"reflect"
	"testing"
)

type T struct {
	A int
}
var T1 struct{
	T
	B string
}
//go:noinline
func TestStruct(t *testing.T) {
	a := T{}
	t1 := reflect.ValueOf(a)
	println(fmt.Sprint(t1.Field(0)))
}

package basic

import (
	"reflect"
	"testing"
)

type Person struct {
	Name string
	Age int
}

func (p Person) Speak() {
	t := reflect.TypeOf(p)
	println(t.Name())
}

type Worker struct {
	Person
	Work string
}

func TestReflect1(t *testing.T) {
}

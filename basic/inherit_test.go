package basic

import (
	"fmt"
	"reflect"
	"testing"
)

func TestStudent(t *testing.T) {
	s := Student{
		Human:      Human{"aaa", 123},
		Class:      3,
		Speciality: "science",
	}
	println(s.Greet())
	Speak(&s)

	var a interface{}
	a = "dsa"
	if a == nil{
		fmt.Println("a is nil")
	}else{
		fmt.Println(reflect.TypeOf(a))
	}

}

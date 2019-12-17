package basic

import (
	"fmt"
	"reflect"
	"testing"
)

func TestStudent1(t *testing.T) {
	a := Student1{Name: "aaa", Age: 12}
	var b Talkable = &a
	of := reflect.TypeOf(b)
	if of.Kind() == reflect.Ptr {
		of = of.Elem()
	}
	fmt.Println(of.Name())
	v := reflect.ValueOf(b)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	fmt.Println(v.Field(2).Int())
}

func TestStudent1_TalkEnglish(t *testing.T) {
	type Student struct {
		Name string
	}
	a := &Student{Name: "aaa"}
	var b interface{} = a
	a.Name = "bbb"
	fmt.Println(b.(*Student).Name)

}

package basic

import (
	"fmt"
	"testing"
	"unsafe"
)

type test struct{
	arr [1000]int
	name string
}

type Empty struct {}

func TestSize(t *testing.T) {
	a := [1000]int{1234}
	fmt.Printf("int array 1000:%v\n", unsafe.Sizeof(a))
	var b interface{}=[1000]int{0}
	fmt.Printf("interface array 1000:%v\n", unsafe.Sizeof(b))
	fmt.Printf("struct:%v\n", unsafe.Sizeof(test{}))
	c := make(map[string]int)
	c["asd"]=23
	c["dasf"]=432
	d := make([]int,1000)
	e := make(chan int,1000)
	fmt.Printf("map:%v\n", unsafe.Sizeof(c))
	fmt.Printf("slice:%v\n", unsafe.Sizeof(d))
	fmt.Printf("chan 1000:%v\n", unsafe.Sizeof(e))
	fmt.Printf("empty struct:%v\n",unsafe.Sizeof(Empty{}))
	type Person struct {
		Name string
		Age int
	}
	g := Person{}
	fmt.Printf("Person struct:%v\n",unsafe.Sizeof(g))

}

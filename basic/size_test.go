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

func TestSize(t *testing.T) {
	a := [1000]int{1234}
	fmt.Printf("int array 1000:%v\n", unsafe.Sizeof(a))
	var b interface{}=[1000]int{0}
	fmt.Printf("interface array 1000:%v\n", unsafe.Sizeof(b))
	fmt.Printf("struct:%v\n", unsafe.Sizeof(test{}))
	c := make(map[string]int)
	c["asd"]=23
	c["dasf"]=432
	fmt.Printf("map:%v\n", unsafe.Sizeof(c))
}

package basic

import (
	"fmt"
	"testing"
)

func TestSliceCap(t *testing.T) {
	a := make([]interface{}, 3, 4)
	fmt.Println(cap(a))
	a = append(a,"da",int64(44))
	fmt.Println(cap(a))
}

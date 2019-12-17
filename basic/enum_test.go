package basic

import (
	"fmt"
	"testing"
)

type EnumType uint8

const (
	AAA EnumType = iota
	BBB
)

func TestEnum(t *testing.T) {
	a := BBB
	fmt.Println(a)
}
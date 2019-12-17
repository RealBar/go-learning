package basic

import (
	"fmt"
	"strings"
	"testing"
)

func TestReplace(t *testing.T) {
	a := "asd fdsa fsd"
	fmt.Println(strings.Replace(a," ","",-1))
}


package basic

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestReplace(t *testing.T) {
	a := "asd fdsa fsd"
	fmt.Println(strings.Replace(a," ","",-1))
}

func TestRegex(t *testing.T) {
	println(time.Now().Format(time.RFC3339))
}
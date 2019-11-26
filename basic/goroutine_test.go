package basic

import (
	"fmt"
	"testing"
)

func TestGo(t *testing.T) {
	m := 0
	f := func() {
		m += 1000
	}
	f()
	fmt.Println(m)
}

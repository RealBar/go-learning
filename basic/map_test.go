package basic

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	a := make(map[int64]bool)
	a[4134] = false
	a[21] = true
	b := a[235234]
	fmt.Printf("%v\n",b)

	c := make(map[string]string)
	d := c["dsaf"]
	fmt.Printf("%v\n",d)
}

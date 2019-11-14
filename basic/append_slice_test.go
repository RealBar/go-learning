package basic

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestSliceAppend(t *testing.T) {
	println("preparing data...")
	ran := rand.New(rand.NewSource(time.Now().Unix()))
	l := 10 * 1024 * 1024
	a := make([]int, l)
	for i := 0; i < l; i++ {
		a[i] = ran.Int()
	}
	b := make([]int, l)
	for i := 0; i < l; i++ {
		b[i] = ran.Int()
	}
	println("preparing data finished")
	t1 := time.Now()
	c := append(a, b...)
	d1 := time.Since(t1)
	fmt.Printf("append uses:%vms,%d\n", d1.Milliseconds(), c[1024])

	t2 := time.Now()
	d := make([]int, 2*l)
	copy(d, a)
	copy(d[l:], b)
	d2 := time.Since(t2)
	fmt.Printf("copy uses:%vms,%d\n", d2.Milliseconds(), d[1024])
}

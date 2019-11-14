package basic

import (
	"testing"
)

// basic example
func TestNewUser(t *testing.T) {
	_ = new(User)
	_ = &User{}
	_ = NewUser()
}

// example1
func TestCall1(t *testing.T) {
	u := &User{}
	Call1(u)
}

// example2
func TestCall2(t *testing.T) {
	u := &User{}
	Call2(u)
}

// example3
func TestCall3(t *testing.T) {
	u := &User{}
	Call3(u)
}

// example4
func TestCall6(t *testing.T) {
	u := &User{}
	Call6(u)
}

func TestCall7(t *testing.T) {
	u := &User{}
	Call7(u)
}

// example5
func TestCall8(t *testing.T) {
	u := &User{}
	Call8(u)
}

// example6
func TestCall9(t *testing.T) {
	u1 := &User{}
	u2 := &User{}
	u3 := &User{}
	Call9(u1, u2, u3)
}

// example0
func TestCall0(t *testing.T) {
	u := &User{}
	Call0(u)
}

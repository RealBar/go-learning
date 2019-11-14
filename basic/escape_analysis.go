package basic

import "fmt"

type User struct {
	Name string
	Age  int
}

type System struct {
	U *User
}

type FullSystem struct {
	S *System
}

// example1
func Call1(u *User) {
	fmt.Printf("%v", u)
}

//example2
func Call2(u *User) int {
	return u.Age * 2
}

// example3
func Call3(u *User) int {
	return Call4(u)
}

func Call4(u *User) int {
	return Call5(u)
}

func Call5(u *User) int {
	return u.Age * 2
}

// example4
func Call6(u *User) {
	s := System{}
	s.U = u
}

func Call7(u *User) {
	s := &System{}
	s.U = u
}

func (s *System) SetUser(u *User) {
	s.U = u
}
func (s System) SetUser1(u *User) {
	s.U = u
}

func Call8(u *User) {
	s := System{}
	// make &s to escape
	f := &FullSystem{}
	f.S = &s
	s.U = u
}

func NewUser() *User {
	return &User{}
}

func Call9(u1 *User,u2 *User,u3 *User) {
	s := make([]*User, 1)
	m := make(map[string]*User)
	c := make(chan *User, 1)
	s[0] = u1
	m["aa"] = u2
	c <- u3
}

// example0
func Call0(u ...interface{}) {
	user := u[0].(*User)
	print(user.Name)
}

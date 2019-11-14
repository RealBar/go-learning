package basic

import "testing"

func TestStudent(t *testing.T) {
	s := Student{
		Human:      Human{"aaa", 123},
		Class:      3,
		Speciality: "science",
	}
	println(s.Greet())
	Speak(&s)
}

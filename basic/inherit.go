package basic

import "fmt"

type Speakable interface {
	Greet() string
}

type Human struct {
	Name string
	Age  int
}

func (h *Human) Greet() string {
	return fmt.Sprintf("Hello, I'm %s, %d yeats old.", h.Name, h.Age)
}

type Student struct {
	Human
	Class      int
	Speciality string
}

func Speak(s Speakable)string{
	return s.Greet()
}
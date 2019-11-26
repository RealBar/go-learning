package basic

import "fmt"

type Talkable interface {
	TalkEnglish(string)
	TalkChinese(string)
}

type Student1 struct {
	Talkable
	Name string
	Age  int
}

func (s *Student1) TalkEnglish(s1 string) {
	fmt.Printf("I'm %s,%d years old,%s", s.Name, s.Age, s1)
}

func (s *Student1) TalkChinese(s1 string) {
	fmt.Printf("我是 %s, 今年%d岁,%s", s.Name, s.Age, s1)
}

package basic

import (
	"fmt"
	"testing"
)

func TestStudent1(t *testing.T) {
	a := Student1{Name: "aaa", Age: 12}
	var b Talkable = &a
	fmt.Println(b)
	b.TalkEnglish("I love you\n")
	b.TalkChinese("我爱你\n")
}

func TestStudent1_TalkEnglish(t *testing.T) {
	type Student struct {
		Name string
	}
	a := &Student{Name:"aaa"}
	var b interface{} = a
	a.Name = "bbb"
	fmt.Println(b.(*Student).Name)


}

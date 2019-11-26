package basic

import (
	"fmt"
	"io"
	"testing"
)

func test1(a int) Speakable {
	return test2(a)
}

func test2(a int) *Student {
	if a < 0 {
		return &Student{
			Human: Human{
				Name: "asd",
				Age:  124,
			},
			Class:      12,
			Speciality: "math",
		}
	}
	return nil
}

type MyReader struct {
	id int64
}

func (r *MyReader) Read(p []byte) (n int, err error) {
	return int(r.id), nil
}

func GetReader(id int64) io.Reader {
	var r *MyReader = nil
	if id > 0 && id < 10000 {
		r = openReader(id)
	}
	return r
}

func openReader(id int64) *MyReader {
	return &MyReader{id: id}
}

func TestSystem_SetUser(t *testing.T) {
	r := GetReader(-2)
	if r == nil {
		fmt.Println("bad reader")
	} else {
		fmt.Println("valid reader")
	}
}

func TestNewUser2(t *testing.T) {

	var c interface{} = nil
	fmt.Println(c == nil)
	var a *Student = nil
	var b interface{} = a
	fmt.Println(b == nil)

}

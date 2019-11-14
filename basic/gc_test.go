package basic

import (
	"fmt"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"testing"
	"time"
)

type Buffer struct {
	b []byte
}

func TestGC(t *testing.T) {
	go func() {
		// 启动一个 http server，注意 pprof 相关的 handler 已经自动注册过了
		if err := http.ListenAndServe(":8081", nil); err != nil {
			t.Fatal(err)
		}
	}()

	goroutineNum := 100
	maxAllocTimes := 1000
	c := make(chan int, goroutineNum)
	for i := 0; i < goroutineNum; i++ {
		fmt.Printf("starting goroutine No.%d\n", i)
		go func() {
			buf := new(Buffer)
			times := rand.Intn(maxAllocTimes)
			for j := 0; j < times; j++ {
				b := make([]byte, 1000_1001)
				b[1000_1000] = byte(rand.Intn(100))
				buf.b = b
				time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
			}
			c <- 1
		}()
	}
	var received = 0
	for {
		received += <-c
		fmt.Printf("%d goroutine finished\n", received)
		if received == goroutineNum {
			break
		}
	}
	fmt.Println("exit successfully!")
}

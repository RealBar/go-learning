package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"sync/atomic"
)

func main() {
	var count int32 = 0
	server := &http.Server{
		Addr: ":4444",
		Handler: http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.Header().Set("Connection", "close")
		}),
		ConnContext: func(ctx context.Context, c net.Conn) context.Context {
			atomic.AddInt32(&count, 1)
			if c2 := ctx.Value("count"); c2 != nil {
				fmt.Printf("发现了遗留数据: %d\n", c2.(int32))
			}
			fmt.Printf("本次数据: %d\n", count)
			return context.WithValue(ctx, "count", count)
		},
	}
	go func() {
		panic(server.ListenAndServe())
	}()

	var err error

	fmt.Println("第一次请求")
	_, err = http.Get("http://localhost:4444/")
	if err != nil {
		panic(err)
	}
	fmt.Println("\n第二次请求")

	_, err = http.Get("http://localhost:4444/")
	if err != nil {
		panic(err)
	}
}

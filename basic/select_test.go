package basic

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestCtx(t *testing.T) {
	c := context.Background()
	ctx,cancel := context.WithCancel(c)
	w := sync.WaitGroup{}
	w.Add(1)
	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("timeout")
		default:
			for{
				fmt.Println("waiting")
				time.Sleep(1000 * time.Second)
				fmt.Println("wake up!")
			}
		}
		w.Done()
	}()
	time.Sleep(2 * time.Second)
	cancel()
	fmt.Println("called cancel")
	w.Wait()
}

func TestSelectDefault(t *testing.T) {
	select {
	case <- time.After(time.Second):
		fmt.Println("time out")
	default:
		for{
			fmt.Println("aaa")
			time.Sleep(time.Second)
		}
	}
}

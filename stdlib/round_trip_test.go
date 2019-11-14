package stdlib

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"
	"time"
)

type TransportWrapper struct {
	http.Transport
	logger *log.Logger
}

func (t *TransportWrapper) RoundTrip(req *http.Request) (*http.Response, error) {
	t.logger.Printf("sending to %s...", req.URL)
	start := time.Now()
	response, err := t.Transport.RoundTrip(req)
	latency := time.Since(start)
	if err != nil {
		t.logger.Printf("error %s", err.Error())
	} else {
		t.logger.Printf("get resp len %d...", response.ContentLength)
	}
	t.logger.Printf("req takes %dms", latency.Milliseconds())
	return response, err
}

func Register() *http.Client {
	l := log.New(os.Stdout, "", 0755)
	dt := http.DefaultTransport.(*http.Transport)
	t := TransportWrapper{
		Transport: *dt,
		logger:    l,
	}
	c := http.Client{
		Transport: &t,
		Timeout:   3000 * time.Millisecond,
	}
	return &c
}

func TestRoundTripper(t *testing.T) {
	req, err := http.NewRequest("GET", "http://baidu.com", nil)
	checkError(err, t)
	client := Register()
	for i := 0; i < 10; i++ {
		fmt.Printf("No.%d req", i)
		_, err := client.Do(req)
		checkError(err, t)
		time.Sleep(time.Second)
	}
}

func checkError(err error, t *testing.T) {
	if err != nil {
		t.Error(err)
	}
}

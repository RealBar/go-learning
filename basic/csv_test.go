package basic

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"testing"
)

func TestCsv(t *testing.T) {
	b := &bytes.Buffer{}
	w := csv.NewWriter(b)
	err := w.Write([]string{"aaa", "bbb", "ccc"})
	if err != nil {
		t.Fatal(err)
	}
	w.Flush()
	for _, b := range b.Bytes() {
		fmt.Printf("%02X,",b)
	}
}

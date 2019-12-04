package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	b, err := ioutil.ReadFile("/Users/yalouwang/Downloads/template.csv")
	checkErr(err)
	reader := csv.NewReader(bytes.NewBuffer(b))
	records, err := reader.ReadAll()
	checkErr(err)
	for i, record := range records {
		fmt.Println(i, strings.Join(record, ","))
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

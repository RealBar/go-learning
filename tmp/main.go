package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Println(ValidateDomain("https://alibaba.com17"))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
const (
	httpPrefix  = "http://"
	httpsPrefix = "https://"
)

var domainPattern = regexp.MustCompile(`^[a-zA-Z0-9-]+\.[a-zA-Z0-9]+(\.[a-zA-Z0-9]+)*$`)
var urlPattern = regexp.MustCompile(`^http(s)?://([a-zA-Z0-9-.]+)/.*$`)

func ValidateDomain(domain string) string {
	if strings.HasPrefix(domain, httpPrefix) {
		domain = strings.TrimLeft(domain, httpPrefix)
	} else if strings.HasPrefix(domain, httpsPrefix) {
		domain = strings.TrimLeft(domain, httpsPrefix)
	}
	if domainPattern.MatchString(domain) {
		return domain
	}
	return ""
}

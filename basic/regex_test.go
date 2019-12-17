package basic

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

const (
	httpPrefix  = "http://"
	httpsPrefix = "https://"
)
var urlPattern = regexp.MustCompile(`^(https://|http://)?([a-zA-Z0-9-.]+)/.*$`)
var domainPattern = regexp.MustCompile(`^[a-zA-Z0-9-]+\.[a-zA-Z0-9]+(\.[a-zA-Z0-9]+)*/.*$`)

func TestRegex(t *testing.T) {
	domain := "http://dwa.gsdag.fdasg.gdasg.gdsag/fdsa/gsadg?fasf=esaf%gsag"
	if strings.HasPrefix(domain, httpPrefix) {
		domain = strings.TrimPrefix(domain, httpPrefix)
	} else if strings.HasPrefix(domain, httpsPrefix) {
		domain = strings.TrimPrefix(domain, httpsPrefix)
	}
	fmt.Println(domainPattern.MatchString(domain))
}
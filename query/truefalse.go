package query

import (
	"fmt"
	"net/url"
)

func isURLValid(input string) bool {
	_, err := url.ParseRequestURI(input)
	return err == nil
}

func Trufalse() {
	testURLs := []string{
		"https://erp.student.najottalim.uz/my-groups/985/staff/985",
		"ftp://example.com/resource",
		"youtube.com",
		"http://192.168.0.1:8080",
		"https://:80",
	}

	for _, url := range testURLs {
		valid := isURLValid(url)
		fmt.Printf("URL: '%s' is valid: %v\n", url, valid)
	}
}

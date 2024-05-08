package main

import (
    "fmt"
    "net/url"
    "os"
	query "31.GO/query"
)

func main() {
    inputURL := "https://erp.student.najottalim.uz/my-groups/985/staff/985"
    
    parsedURL, err := url.Parse(inputURL)
    if err != nil {
        fmt.Println("Error parsing URL:", err)
        os.Exit(1)
    }

	fmt.Println("------1-shart------")

    fmt.Println("Schema: ", parsedURL.Scheme)
    fmt.Println("Host: ", parsedURL.Host)
    fmt.Println("Path: ", parsedURL.Path)
    fmt.Println("Query Params: ", parsedURL.RawQuery)
    fmt.Println("Fragment: ", parsedURL.Fragment)


	fmt.Println("------2-shart------")
	query.QueryParametrs()

	fmt.Println("------3-shart------")
	query.Trufalse()
	
}

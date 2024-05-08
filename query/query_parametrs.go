package query

import (
	"fmt"
	"net/url"
	"os"
)

func QueryParametrs() {
	inputURL := "https://erp.student.najottalim.uz/my-groups/985/staff/985"

	parsedURL, err := url.Parse(inputURL)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		os.Exit(1)
	}
	queryParameters := parsedURL.Query()
	if len(queryParameters) == 0 {
		fmt.Println("No query parameters found.")
	} else {
		fmt.Println("Query Parameters:")
		for key, values := range queryParameters {
			fmt.Printf("%s: %v\n", key, values)
		}
	}
}

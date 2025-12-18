package main

import (
	"fmt"
	"net/url"
)

func main() {

	// [scheme://][userinfo@]host[:port][/path][?querty][#fragment]

	rawURL := "https://example.com:8080/path?querty=param#fragment"

	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Schame:", parsedURL.Scheme)
	fmt.Println("Host:", parsedURL.Host)
	fmt.Println("Port:", parsedURL.Port())
	fmt.Println("Path:", parsedURL.Path)
	fmt.Println("Raw Query:", parsedURL.RawQuery)
	fmt.Println("Fragment:", parsedURL.Fragment)

	rawURL1 := "https://example.com/path?name=John&age=30"
	parsedURL1, err := url.Parse(rawURL1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	queryParams := parsedURL1.Query()
	fmt.Println(queryParams)
	fmt.Println("Name:", queryParams.Get("name"))
	fmt.Println("Age:", queryParams.Get("age"))

	baseURL := &url.URL{
		Scheme: "https",
		Host:   "example.com",
		Path:   "/path",
	}

	query := baseURL.Query()
	query.Set("name", "John")
	query.Set("age", "22")
	baseURL.RawQuery = query.Encode()

	fmt.Println("Built URL:", baseURL.String())

	values := url.Values{}

	values.Add("name", "Jane")
	values.Add("age", "24")
	values.Add("city", "Minsk")
	values.Add("country", "Belarus")

	encodedQuery := values.Encode()
	fmt.Println(encodedQuery)

	baseURL1 := "https://example.com/search"
	fullURL := baseURL1 + "?" + encodedQuery

	fmt.Println(fullURL)
}

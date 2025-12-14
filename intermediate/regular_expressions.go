package main

import (
	"fmt"
	"regexp"
)

func main() {

	fmt.Println("He said, \"I am great\"")
	fmt.Println(`He said, "I am great"`)

	// Compile a regex pattern to match email address
	re := regexp.MustCompile(`[a-zA-Z0-9._+%-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)

	// Test strings
	email1 := "user@email.com"
	email2 := "invalid_email"

	// Match
	fmt.Println("Email:", re.MatchString(email1))
	fmt.Println("Email:", re.MatchString(email2))

	// Capturing Groups
	// Compile a regex pattern to capture data compopnents
	re = regexp.MustCompile(`(\d{2})-(\d{2})-(\d{4})`)

	// Test string
	date := "14-12-2025"

	// Find all submatches
	submatch := re.FindStringSubmatch(date)
	fmt.Println(submatch)
	fmt.Println(submatch[0])
	fmt.Println(submatch[1])
	fmt.Println(submatch[2])
	fmt.Println(submatch[3])

	// Source string
	str := "Hello World"
	re = regexp.MustCompile(`[aeioul]`)

	result := re.ReplaceAllString(str, "*")
	fmt.Println(result)

	// i - case insensitive
	// m - milti line model
	// s - dot matches all

	re = regexp.MustCompile(`(?i)go`)

	// Test string
	text := "Golang is going great"

	// Match
	fmt.Println("Match:", re.MatchString(text))
}

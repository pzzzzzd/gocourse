package main

import (
	"fmt"
	"strings"
)

func main() {

	// str := "Hello Go!"

	// fmt.Println(len(str))
	// str1 := "Hello"
	// str2 := "World"
	// result := str1 + " " + str2
	// fmt.Println(result)

	// fmt.Println(str[0])
	// fmt.Println(str[1:7])

	// // String Conversion
	// num := 19
	// str3 := strconv.Itoa(num)
	// fmt.Println(len(str3))

	// // Strings splitting
	// fruints := "apple, orange, banana"
	// fmt.Println(fruints)
	// parts := strings.Split(fruints, ",")
	// fmt.Println(parts)

	// countries := []string{"Germany", "France", "Italy"}
	// joined := strings.Join(countries, ", ")
	// fmt.Println(joined)

	// fmt.Println(strings.Contains(str, "Go"))
	// fmt.Println(strings.Contains(str, "Go222"))

	// replaced := strings.Replace(str, "Go", "Univ", 1)
	// fmt.Println(replaced)

	// strwspace := " Hello Everyone "
	// fmt.Println(strwspace)
	// fmt.Println(strings.TrimSpace(strwspace))

	// fmt.Println(strings.ToLower(strwspace))
	// fmt.Println(strings.ToUpper(strwspace))

	// fmt.Println(strings.Repeat("foo ", 3))

	// fmt.Println(strings.Count("Hello", "He"))
	// fmt.Println(strings.Count("Hello", "l"))

	// fmt.Println(strings.HasPrefix("Hello", "He"))
	// fmt.Println(strings.HasPrefix("Hello", "he"))

	// str5 := "Hello, 1244_3333 Go 6643"
	// re := regexp.MustCompile(`\d+`)
	// mathes := re.FindAllString(str5, -1)
	// fmt.Println(mathes)

	// str6 := "Hello, ну типо да"
	// fmt.Println(utf8.RuneCountInString(str6))

	// STRING BUILDER
	var builder strings.Builder

	// Write some strings
	builder.WriteString("Hello")
	builder.WriteString(", ")
	builder.WriteString("World")

	// Convert builder to a string
	result := builder.String()
	fmt.Println(result)

	// Using Writerune to add a character
	builder.WriteRune(' ')
	builder.WriteString("How are u")

	result = builder.String()
	fmt.Println(result)

	// Reset the builder
	builder.Reset()
	builder.WriteString("Starting fresh")
	result = builder.String()
	fmt.Println(result)

}

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	message := "Hello, \nGo!"
	message1 := "Hello, \tGo!"
	message2 := "Hello, \rGo!" // Go!lo
	rawMessage := `Hello \nGo`

	fmt.Println(message)
	fmt.Println(message1)
	fmt.Println(message2)
	fmt.Println(rawMessage)

	fmt.Println("Length of message variable is", len(message))

	fmt.Println("The first character in message var is", message[0]) // ASCII

	greeting := "Hello "
	name := "Alice"
	fmt.Println(greeting + name)

	str1 := "Apple"  // A has an ASCII value 65
	str := "apple"   // A has an ASCII value 65
	str2 := "banana" // b has an ASCII value 98
	str3 := "app"    // a has an ASCII value 97

	fmt.Println(str1 < str2)
	fmt.Println(str3 < str1)
	fmt.Println(str > str1)
	fmt.Println(str > str3)

	for _, char := range message {
		// fmt.Printf("Character at index %d is %c\n", i, char)
		fmt.Printf("%v\n", char)
	}

	fmt.Println("Rune count:", utf8.RuneCountInString(greeting))

	greetingWithName := greeting + name
	fmt.Println(greetingWithName)

	var ch rune = 'a'
	rch := 'и'
	fmt.Println(ch)
	fmt.Println(rch)

	fmt.Printf("%c\n", ch)
	fmt.Printf("%c\n", rch)

	cstr := string(ch)
	fmt.Println(cstr)

	fmt.Printf("Type of cstr if %T\n", cstr)

	const NIHONGO = "ну да"
	fmt.Println(NIHONGO)

	rhello := "привет"
	for _, runeValue := range rhello {
		fmt.Printf("%c\n", runeValue)
	}

	r := '😎'
	fmt.Printf("%v\n", r)
	fmt.Printf("%c\n", r)

}

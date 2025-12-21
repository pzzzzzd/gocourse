package main

import (
	"encoding/base64"
	"fmt"
)

func main() {

	data := []byte("hello, Base64")

	encoded := base64.StdEncoding.EncodeToString(data)
	fmt.Println(encoded)

	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("Decoded:", decoded)
	fmt.Println("Decoded:", string(decoded))

	urlSafeEncoded := base64.URLEncoding.EncodeToString(data)
	fmt.Println("Url Safe Encoded:", urlSafeEncoded)

}

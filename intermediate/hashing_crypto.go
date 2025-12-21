package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
)

func main() {

	password := "qwerty"

	// hash := sha256.Sum256([]byte(password))
	// hash512 := sha512.Sum512([]byte(password))

	// fmt.Println(password)
	// fmt.Println(hash)

	// fmt.Printf("SHA-256 Hash hex val: %x\n", hash)
	// fmt.Printf("SHA-512 Hash hex val: %x\n", hash512)
	salt, err := generateSalt()
	fmt.Println("Original Salt:", salt)
	fmt.Printf("Original Salt: %x\n", salt)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	singUpHash := hashPassword(password, salt)

	saltStr := base64.StdEncoding.EncodeToString(salt)
	fmt.Println("Salt:", saltStr)
	fmt.Println("Hash:", singUpHash)
	hashOriginalPassword := sha256.Sum256([]byte(password))
	fmt.Println("Hash of just the password string without salt:", base64.StdEncoding.EncodeToString(hashOriginalPassword[:]))

	decodeSalt, err := base64.StdEncoding.DecodeString(saltStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	loginHash := hashPassword("password123", decodeSalt)

	if singUpHash == loginHash {
		fmt.Println("Password is correct")
	} else {
		fmt.Println("Login failde")
	}

}

func generateSalt() ([]byte, error) {
	salt := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

func hashPassword(password string, salt []byte) string {
	saltedPassword := append(salt, []byte(password)...)
	hash := sha256.Sum256(saltedPassword)
	return base64.StdEncoding.EncodeToString(hash[:])
}

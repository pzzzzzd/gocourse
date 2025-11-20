<<<<<<< HEAD
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	sourse := rand.NewSource(time.Now().UnixNano())
	random := rand.New(sourse)

	target := random.Intn(100) + 1
	fmt.Println("Welcom to guessing game")
	fmt.Println("I have chosen a number between 1 to 100")
	fmt.Println(target)

	var guess int
	for {
		fmt.Println("Enter your gess: ")
		fmt.Scanln(&guess)
		if guess < target {
			fmt.Println("Number is more")
			continue
		}
		if guess > target {
			fmt.Println("Number is less")
			continue
		}
		if guess == target {
			fmt.Println("Congratulations your answer is correct")
			break
		}
	}
}
=======
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	sourse := rand.NewSource(time.Now().UnixNano())
	random := rand.New(sourse)

	target := random.Intn(100) + 1
	fmt.Println("Welcom to guessing game")
	fmt.Println("I have chosen a number between 1 to 100")
	fmt.Println(target)

	var guess int
	for {
		fmt.Println("Enter your gess: ")
		fmt.Scanln(&guess)
		if guess < target {
			fmt.Println("Number is more")
			continue
		}
		if guess > target {
			fmt.Println("Number is less")
			continue
		}
		if guess == target {
			fmt.Println("Congratulations your answer is correct")
			break
		}
	}
}
>>>>>>> fb36f3b (wls_push)

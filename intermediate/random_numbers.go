package main

import (
	"fmt"
	"math/rand"
)

func main() {

	// val := rand.New(rand.NewSource(time.Now().Unix()))

	// fmt.Println(rand.Intn(6) + 5)
	// fmt.Println(val.Intn(101))
	// fmt.Println(rand.Float64())

	for {
		fmt.Println("Dice Game")
		fmt.Println("1. Roll the dice")
		fmt.Println("2. Exit")
		fmt.Print("Enter ur choice 1 or 2: ")

		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil || (choice != 1 && choice != 2) {
			fmt.Println("Invalid choice")
			continue
		}
		if choice == 2 {
			fmt.Println("Bye")
			break
		}

		die1 := rand.Intn(6) + 1
		die2 := rand.Intn(6) + 1

		fmt.Printf("U rolled a %d and a %d\n", die1, die2)
		fmt.Println("Total:", die1+die2)

		fmt.Print("Do u want to roll again?(y/n): ")
		var rollAgain string
		_, err = fmt.Scan(&rollAgain)
		if err != nil || (rollAgain != "y" && rollAgain != "n") {
			fmt.Println("Invalid input")
			continue
		}
		if rollAgain == "n" {
			fmt.Println("Bye")
			break
		}
	}

}

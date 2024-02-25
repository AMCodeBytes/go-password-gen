package main

import (
	"fmt"
	"os"
)

const characters string = "0123456789 abcdefghijklmnopqrstuvwxyz#,.!?:;@$%~&^_-"

func main() {
	var choice int

	fmt.Println("Welcome to the password generator")

	for {
		displayMenu()

		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			// Generate a password
		case 2:
			// Customise the password length
		case 3:
			os.Exit(0)
			return
		default:
			panic("Invalid input entered.")
		}
	}
}

func displayMenu() {
	fmt.Println("Please select what you'd like to do")
	fmt.Println("1. Generate password")
	fmt.Println("2. Password length")
	fmt.Println("3. Quit")
}

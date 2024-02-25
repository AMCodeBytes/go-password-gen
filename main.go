package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

var characters string = "0123456789 ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz#,.!?:;@$%~&^_-"
var length int = 12

func main() {
	var password string
	var choice int

	fmt.Println("Welcome to the password generator")

	for {
		password = ""
		displayMenu()

		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			password = generatePassword()
			fmt.Println("-----  Password  -----")
			fmt.Println(password)
		case 2:
			length = updatePasswordLength()
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

func generatePassword() string {
	var pass string
	char := strings.Split(characters, "")

	for i := 0; i < length; i++ {
		pass += char[rand.Intn(len(char))]
	}

	return pass
}

func updatePasswordLength() int {
	var newLength int

	fmt.Print("New password length: ")
	fmt.Scan(&newLength)

	return newLength
}

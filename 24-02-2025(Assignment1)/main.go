package main

import (
	"fmt"
	"24-02-2025/question1"
	"24-02-2025/question2"
)

func main() {
	for {
		fmt.Println("Select an option:")
		fmt.Println("1. Run Question1")
		fmt.Println("2. Run Question2")
		fmt.Println("3. Exit")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			question1.Run()
		case 2:
			question2.Run()
		case 3:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option. Please choose again.")
		}
	}
}

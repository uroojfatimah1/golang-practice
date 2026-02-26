package main

import "fmt"

func main() {
	var username string = "John Doe"
	fmt.Printf("Hello, %s! \n", username)
	fmt.Printf("The type for respected variable is: %T \n", username)

	var isLoggedIn bool = true
	fmt.Printf("Is the user logged in: %t \n", isLoggedIn)
	fmt.Printf("The type for respected variable is: %T \n", isLoggedIn)

	var index int = 212
	fmt.Printf("Index is: %d\n", index)
	fmt.Printf("The type for respected variable is: %T \n", index)
}

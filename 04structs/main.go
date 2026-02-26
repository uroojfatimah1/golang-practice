package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	user := User{"name", 12, "7", 1, true}
	fmt.Printf("User details: %+v\n", user)
	user.GetStatus()

	loginCount := 7
	if loginCount > 10 {
		fmt.Println("User eligible for exam")
	} else {
		fmt.Println("User not eligible for exam")
	}

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Printf("Dice number: %d\n", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice 1")
	case 2:
		fmt.Println("Dice 2")
	case 3:
		fmt.Println("Dice 3")
	case 4:
		fmt.Println("Dice 4")
	case 5:
		fmt.Println("Dice 5")
	case 6:
		fmt.Println("Dice 6")
	default:
		fmt.Println("Dice default")
	}
}

type User struct {
	Name       string
	Age        int
	Class      string
	rollNo     int
	isLoggedIn bool
}

func (u User) GetStatus() {
	fmt.Println("Is user logged in: ", u.isLoggedIn)
}

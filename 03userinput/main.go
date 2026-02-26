package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	welcome := "Welcome"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter username: ")
	username, _ := reader.ReadString('\n')

	fmt.Print("Enter rating: ")
	rating, _ := reader.ReadString('\n')

	numRating, err := strconv.ParseFloat(strings.TrimSpace(rating), 64)

	fmt.Print("\n")
	if err != nil {
		fmt.Println("Error parsing rating ", err)
	} else {
		fmt.Println("Updated rating is: ", numRating+1)
	}
	fmt.Print("Username is: ", username)
	fmt.Print("Provided Rating is: ", rating)
	fmt.Println("Error encountered: ", err)

}

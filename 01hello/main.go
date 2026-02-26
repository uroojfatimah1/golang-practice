package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello!")

	presentTime := time.Now()
	println(presentTime.Format("2006-01-02 15:04:05 Monday"))

	createdDate := time.Date(2025, time.December, 28, 0, 0, 0, 0, time.UTC)
	println(createdDate.Format("2006-01-02 15:04:05 Monday"))

	myNumber := 23
	var ptr = &myNumber
	println("The number is", *ptr)
	println("The address is", ptr)

	var fruitList = []string{"Apple", "Orange", "Pear"}

	fmt.Println("The fruits in list are", fruitList)

	fruitList = append(fruitList, "Mango", "Orange")
	fmt.Println("Updated list is: ", fruitList)

	numberList := make([]byte, 5)
	numberList[0] = 7
	numberList[1] = 9
	numberList[2] = 3

	fmt.Println("The number are", numberList)

	fmt.Append(numberList, 8, 1)
	fmt.Println("The numbers are", numberList)

	languages := make(map[string]string)
	languages["GO"] = "GoLang"
	languages["JS"] = "JavaScript"
	languages["PY"] = "Python"
	languages["RU"] = "Ruby"

	println(languages)
	println(languages["RU"])

	for key, value := range languages {
		fmt.Println("Key:", key, "Value:", value)
	}
}

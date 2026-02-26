package main

import (
	"fmt"
	"os"
)

func main() {
	content := `Appended Text here
				Day 4: Redis Introduction
				Day 5: Redis Practice
				Day 6: Integrate redis into a go-lang project
				Focus: Building backend fundamentals
				`

	file, err := os.Create("./test.txt")
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("File created: %s\n", file.Name())
		file, err := os.OpenFile("notes.txt", os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			checkError(err)
		}
		defer file.Close()
		_, err = file.WriteString(content)
		checkError(err)
		//file.WriteString("2nd line\n")
		//file.WriteString("3rd line\n")
		checkError(err)
		os.Rename("notes.txt", "notes.txt")
		fmt.Printf("File renamed: %s\n", file.Name())
		os.Remove("test.txt")
		fmt.Printf("File deleted: %s\n", file.Name())
		//data, _ := os.ReadFile("./test.txt")
		//fmt.Printf("File content: %s\n", string(data))
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

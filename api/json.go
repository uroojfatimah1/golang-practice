package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Course struct {
	Name     string `json:"CourseName"`
	Price    int
	Password string
	Platform string
	Tags     []string
}

func main() {
	EncodeJson()
}

func EncodeJson() {
	lcoCourses := []Course{
		{"React", 99, "123asd", "desktop", []string{"web-dev", "js"}},
		{"JS", 99, "123asd", "desktop", []string{"web-dev", "js"}},
		{"React Native", 99, "123asd", "mobile", []string{"mobile-dev", "js"}},
	}

	EncodedJson, _ := json.MarshalIndent(lcoCourses, "", "\t")
	fmt.Printf("%s\n", EncodedJson)
}

func DecodeJson() {
	fetchedJsonData, _ := os.ReadFile("course.json")

	checkValid := json.Valid(fetchedJsonData)
	if !checkValid {
		fmt.Println("Invalid JSON")
	}
	json.Unmarshal(fetchedJsonData, &Course{})
	fmt.Printf("%+v\n", Course{})
}

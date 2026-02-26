package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	GetRequest()
	PostRequest()
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func GetRequest() {
	resp, err := http.Get("https://www.google.com")
	checkError(err)
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Content Length:", resp.ContentLength)

	content, _ := io.ReadAll(resp.Body)
	fmt.Println("Response Body:", string(content))
}

func PostRequest() {
	resp, err := http.Post("https://www.google.com", "text/html; charset=utf-8", nil)
	checkError(err)
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Content Length:", resp.ContentLength)
}

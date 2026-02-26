package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

const exampleUrl = "https://go.dev/ref/spec"

func main() {
	response, err := http.Get(exampleUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("File content: %s\n", string(data))

	urlResult, _ := url.Parse(exampleUrl)

	fmt.Println(urlResult.Scheme + "://" + urlResult.Host + urlResult.Path + urlResult.RawQuery + urlResult.Fragment + urlResult.RawQuery + urlResult.Port())
	fmt.Println(urlResult.Scheme)
	fmt.Println(urlResult.Host)
	fmt.Println(urlResult.Path)
	fmt.Println(urlResult.RawQuery)
	fmt.Println(urlResult.Fragment)
	fmt.Println(urlResult.Port())
}

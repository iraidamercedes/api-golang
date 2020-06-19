// substring_matching.go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func getTitle() {
	// Make HTTP GET request
	response, err := http.Get("https://www.sibcolombia.net")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Get the response body as a string
	dataInBytes, err := ioutil.ReadAll(response.Body)
	pageContent := string(dataInBytes)

	// Find a substr
	titleStartIndex := strings.Index(pageContent, "<title>")
	if titleStartIndex == -1 {
		fmt.Println("No title element found")
		os.Exit(0)
	}

	titleStartIndex += 7

	// Find the index of the closing tag
	titleEndIndex := strings.Index(pageContent, "</title>")
	if titleEndIndex == -1 {
		fmt.Println("No closing tag for title found.")
		os.Exit(0)
	}

	pageTitle := []byte(pageContent[titleStartIndex:titleEndIndex])

	// Print out the result
	fmt.Printf("Page title: %s\n", pageTitle)
}

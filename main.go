package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const dataURL = "http://sfdbi.org/building-permits-filed-and-issued"

func main() {
	linksData, err := ioutil.ReadFile("links.txt")
	if err != nil {
		fmt.Printf("Trouble reading file: %v", err)
	}

	links := strings.Split(string(linksData), ", ")

	// if the link ends with .xlsx, .xls
	for _, url := range links {
		if strings.HasSuffix(url, ".xlsx") {
			resp, err := http.Get(url)
			if err != nil {
				fmt.Printf("Could not fetch %v because %v", url, err)
			}
			defer resp.Body.Close()

			// save the file
			// get some ideas from https://github.com/thbar/golang-playground/blob/master/download-files.go
		} else {
			fmt.Printf("Skipping: %v", url)
		}
	}
	// or redirects to an .xlsx or .xls

}

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	links := parseLinks()

	for _, url := range links {
		if isExcelDocument(url) {
			// parallelize me!
			downloadFromURL(url)
		} else {
			// follow redirects
			// if then ends with .xlsx or .xls fetch it
			fmt.Printf("Skipping: %v", url)
		}
	}
}

func downloadFromURL(url string) error {
	tokens := strings.Split(url, "/")
	fileName := tokens[len(tokens)-1]
	fmt.Printf("Downloading %v to %v", url, fileName)

	content, err := os.Create("temp_docs/" + fileName)
	if err != nil {
		fmt.Printf("Error while creating %v because of %v", fileName, err)
		return err
	}

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Could not fetch %v because %v", url, err)
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(content, resp.Body)
	if err != nil {
		fmt.Printf("Error while saving %v from %v", fileName, url)
		return err
	}

	return nil
}

func isExcelDocument(url string) bool {
	return strings.HasSuffix(url, ".xlsx") || strings.HasSuffix(url, ".xls")
}

func parseLinks() []string {
	linksData, err := ioutil.ReadFile("links.txt")
	if err != nil {
		fmt.Printf("Trouble reading file: %v", err)
	}

	links := strings.Split(string(linksData), ", ")

	return links
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type user struct {
	name     string
	email    string
	url      string
	username string
}

func main() {
	file, err := os.Open("urls.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		url := scanner.Text()
		processURL(url)
	}
}

func processURL(url string) {
	doc := downloadURL(url)
	if strings.Contains(url, "/orgs/") {
		fmt.Println("Scraping organization: " + url)
		scrapeOrganization(doc, url)
	} else if strings.Contains(url, "/search?") {
		fmt.Println("Scraping search results: " + url)
		scrapeSearch(doc, url)
	} else if strings.Contains(url, "/stargazers") {
		fmt.Println("Scraping stargazers: " + url)
		scrapeStarGazers(doc, url)
	} else {
		fmt.Println("Scraping Github profile: " + url)
		scrapeProfile(doc)
	}
	fmt.Println("====================")
}

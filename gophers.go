package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type user struct {
	name     string
	email    string
	url      string
	username string
}

func main() {
	url := flag.String("github_url", "", "github url you want to scrape")
	flag.Parse()
	githubURL := *url
	doc, err := goquery.NewDocument(githubURL)
	if err != nil {
		log.Fatal(err)
	}

	if strings.Contains(githubURL, "/orgs/") {
		fmt.Println("Organization URL. Beginning to scrape.")
		scrapeOrganization(doc, githubURL)
	} else if strings.Contains(githubURL, "/search?") {
		fmt.Println("Search URL. Beginning to scrape.")
		scrapeSearch(doc, githubURL)
	} else if strings.Contains(githubURL, "/stargazers") {
		fmt.Println("Stargazer URL. Beginning to scrape.")
		scrapeStarGazers(doc, githubURL)
	} else {
		fmt.Println("Single profile URL. Beginning to scrape.")
		scrapeProfile(doc)
	}
}

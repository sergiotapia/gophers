package main

import (
	"flag"
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
		scrapeOrganization(doc, githubURL)
	} else if strings.Contains(githubURL, "/search?") {
		scrapeSearch(doc, githubURL)
	} else if strings.Contains(githubURL, "/stargazers") {
		scrapeStarGazers(doc, githubURL)
	} else {
		scrapeProfile(doc)
	}
}

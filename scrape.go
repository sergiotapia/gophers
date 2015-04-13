package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

func scrapeUserFromSearch(element *goquery.Selection) (user, error) {
	email := element.Find("a.email").Text()
	url, _ := element.Find("a").Attr("href")
	username := element.Find("em").Text()
	name := ""

	// info := element.Find(".user-list-info")
	// info = info.Remove(".fooa")
	// info = info.Remove("a")
	// name := info.Text()

	user := user{name: name, email: email, url: url, username: username}
	return user, nil
}

func scrapeProfile(document *goquery.Document) {
	vcard := document.Find(".vcard")

	email := vcard.Find("a.email").Text()
	url := "http://github.com/" + vcard.Find(".vcard-username").Text()
	username := vcard.Find(".vcard-username").Text()
	name := vcard.Find(".vcard-fullname").Text()

	fmt.Println("Parsed user profile:" + username)
	user := user{name: name, email: email, url: url, username: username}
	dumpToCSV(user)
}

func scrapeOrganization(document *goquery.Document, url string) {
	pagesStr := document.Find("a.next_page").Prev().Text()
	pages, _ := strconv.Atoi(pagesStr)
	page := 1
	for page <= pages {
		doc, err := goquery.NewDocument(url + "?page=" + strconv.Itoa(page))
		if err != nil {
			log.Fatal(err)
		}

		doc.Find(".member-list-item").Each(func(i int, s *goquery.Selection) {
			scrapeOrganizationItem(s)
		})

		page = page + 1
	}
}

func scrapeOrganizationItem(element *goquery.Selection) {
	url := "http://github.com/" + element.Find(".member-username").Text()
	username := element.Find(".member-username").Text()
	name := element.Find(".member-fullname").Text()

	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}
	email := doc.Find(".vcard").Find("a.email").Text()
	fmt.Println("Parsed organization user:" + username)
	user := user{name: name, email: email, url: url, username: username}
	dumpToCSV(user)
}

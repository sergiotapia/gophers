package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func scrapeSearch(document *goquery.Document, url string) {
	pagesStr := document.Find("a.next_page").Prev().Text()
	pages, _ := strconv.Atoi(pagesStr)
	page := 1
	for page <= pages {
		pageURL := url + "&p=" + strconv.Itoa(page)
		fmt.Println("Analyzing page: " + pageURL)
		doc := downloadURL(pageURL)
		doc.Find(".user-list-item").Each(func(i int, s *goquery.Selection) {
			email := s.Find("a.email").Text()
			profileURL, _ := s.Find("a").Eq(1).Attr("href")
			username := profileURL[1:len(profileURL)]
			profileURL = "http://github.com" + profileURL
			info := s.Find(".user-list-info")
			_ = info.Find("ul.user-list-meta").Remove()
			_ = info.Find("a").Remove()
			name := strings.TrimSpace(info.Text())
			fmt.Println("Parsed user: " + username)
			user := user{name: name, email: email, url: profileURL, username: username}
			dumpToCSV(user)
		})

		page = page + 1
	}
}

func scrapeProfile(document *goquery.Document) {
	vcard := document.Find(".vcard")

	email := vcard.Find("a.email").Text()
	url := "http://github.com/" + vcard.Find(".vcard-username").Text()
	username := vcard.Find(".vcard-username").Text()
	name := vcard.Find(".vcard-fullname").Text()

	fmt.Println("Parsed user: " + username)
	user := user{name: name, email: email, url: url, username: username}
	dumpToCSV(user)
}

func scrapeOrganization(document *goquery.Document, url string) {
	nextPage := document.Find("a.next_page").Length()
	if nextPage == 0 {
		document.Find(".member-list-item").Each(func(i int, s *goquery.Selection) {
			scrapeOrganizationItem(s)
		})
	} else {
		pagesStr := document.Find("a.next_page").Prev().Text()
		pages, _ := strconv.Atoi(pagesStr)
		page := 1
		for page <= pages {
			pageURL := url + "?page=" + strconv.Itoa(page)
			doc := downloadURL(pageURL)
			doc.Find(".member-list-item").Each(func(i int, s *goquery.Selection) {
				scrapeOrganizationItem(s)
			})

			page = page + 1
		}
	}
}

func scrapeOrganizationItem(element *goquery.Selection) {
	url := "http://github.com/" + element.Find(".member-username").Text()
	username := element.Find(".member-username").Text()
	name := element.Find(".member-fullname").Text()

	doc := downloadURL(url)
	email := doc.Find(".vcard").Find("a.email").Text()
	fmt.Println("Parsed user: " + username)
	user := user{name: name, email: email, url: url, username: username}
	dumpToCSV(user)
}

func scrapeStarGazers(document *goquery.Document, url string) {
	totalStarGazers := 0
	document.Find("ul.pagehead-actions li").Each(func(i int, s *goquery.Selection) {
		if i == 1 {
			count := strings.TrimSpace(s.Find("a.social-count").Text())
			totalStarGazers, _ = strconv.Atoi(count)
		}
	})
	pages := totalStarGazers / 30
	page := 1
	for page <= pages {
		pageURL := url + "?page=" + strconv.Itoa(page)
		doc := downloadURL(pageURL)
		doc.Find(".follow-list-item span.css-truncate a").Each(func(i int, s *goquery.Selection) {
			profileURL, _ := s.Attr("href")
			profileURL = "http://github.com" + profileURL
			profile := downloadURL(profileURL)
			scrapeProfile(profile)
		})

		page = page + 1
	}
}

func downloadURL(url string) *goquery.Document {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(5 * time.Second)
	return doc
}

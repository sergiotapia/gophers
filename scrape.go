package main

import "github.com/PuerkitoBio/goquery"

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

	user := user{name: name, email: email, url: url, username: username}
	dumpToCSV(user)
}

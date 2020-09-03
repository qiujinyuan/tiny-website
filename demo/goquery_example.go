package demo

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func GoqueryExample() {
	// Request the HTML page.
	res, err := http.Get("http://www.yinwang.org/")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	listSelection := doc.Find(".list-group .list-group-item")
	fmt.Println(listSelection.Length())
	listSelection.Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		band := s.Find("a").Text()
		href, _ := s.Find("a").Attr("href")
		fmt.Printf("Review %d: %s - %s\n", i, band, href)
	})
}

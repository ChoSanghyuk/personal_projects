package gold

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// crwal data of https://goldprice.org/, whose path is #goldchange > div > div > div > div > div.tick-value-wrap.d-flex.align-items-center.justify-content-center.flex-wrap > div.tick-value.price-value > span
func Crawl() {
	fmt.Println(strings.Repeat("#", 50))
	// The URL you want to crawl
	url := "https://goldprice.org/"

	// The CSS path you want to extract data from
	cssPath := "#goldchange > div > div > div > div > div.tick-value-wrap.d-flex.align-items-center.justify-content-center.flex-wrap > div.tick-value.price-value > span"

	// Send a GET request to the URL
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	// Check the response status
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Create a goquery document from the response body
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(doc.Find("span").Text())
	// Find the elements by the CSS selector
	doc.Find(cssPath).Each(func(i int, s *goquery.Selection) {
		// Extract and print the data
		fmt.Println(s.Text())
	})
}

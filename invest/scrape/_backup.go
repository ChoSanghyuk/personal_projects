package crawl

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func Crawl_Backup() {
	fmt.Println(strings.Repeat("#", 50))
	// The URL you want to crawl
	// url := config.ConfigInfo.GoldConfig.Crawl.Url
	url := "https://search.naver.com/search.naver?where=nexearch&sm=top_hty&fbm=0&ie=utf8&query=%ED%95%9C%EA%B5%AD%EA%B1%B0%EB%9E%98%EC%86%8C+%EC%8B%A4%EC%8B%9C%EA%B0%84+%EA%B8%88+%EC%8B%9C%EC%84%B8"
	// cssPath := "#goldchange > div > div > div > div > div.tick-value-wrap.d-flex.align-items-center.justify-content-center.flex-wrap > div.tick-value.price-value > span"
	cssPath := "#main_pack > section.sc_new.pcs_gold_rate._cs_gold_rate > div > div.gold_price.up > a > strong"

	jar, _ := cookiejar.New(nil)
	client := &http.Client{
		Jar: jar,
	}

	// Send a GET request to the URL
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set a custom User-Agent header
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")

	// Send the request
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}

	defer res.Body.Close()

	// Check the response status
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// show me the content of body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Convert the response body to a string and print it
	fmt.Println(string(body))

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

func SampleCrawl() {
	url := "https://ncov.mohw.go.kr/"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// HTML 읽기
	html, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// 현황판 파싱
	wrapper := html.Find("ul.liveNum")
	items := wrapper.Find("li")

	// items 순회하면서 출력
	items.Each(func(idx int, sel *goquery.Selection) {
		title := sel.Find("strong.tit").Text()
		value := sel.Find("span.num").Text()
		before := sel.Find("span.before").Text()

		fmt.Println(title, value, before)
	})
}

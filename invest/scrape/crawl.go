package scrape

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/gofiber/fiber/v2/log"
)

func (s Scraper) Crawl(url string, cssPath string) (string, error) {

	// Send the request
	res, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("error making request\n%w", err)
	}

	defer res.Body.Close()

	// Check the response status
	if res.StatusCode != 200 {
		body, _ := io.ReadAll(res.Body)
		return "", fmt.Errorf("status code error: %d %s %s", res.StatusCode, res.Status, body)
	}

	// Create a goquery document from the response body
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", fmt.Errorf("error creating document\n%w", err)
	}

	// fmt.Println(doc.Text())

	var v string
	// Find the elements by the CSS selector
	doc.Find(cssPath).Each(func(i int, s *goquery.Selection) {
		// Extract and print the data
		v = s.Text()
	})

	return v, nil
}

func (s *Scraper) ExchageRate() float64 {

	if s.Exchange.Rate != 0 && s.Exchange.Date.Format("20060102") == time.Now().Format("20060102") {
		return s.Exchange.Rate
	}

	// Todo config화 시킬지 결정
	url := "https://search.naver.com/search.naver?where=nexearch&sm=top_hty&fbm=0&ie=utf8&query=%ED%99%98%EC%9C%A8"
	cssPath := "#main_pack > section.sc_new.cs_nexchangerate > div:nth-child(1) > div.exchange_bx._exchange_rate_calculator > div > div.inner > div:nth-child(3) > div.num > div > span"

	rtn, err := s.Crawl(url, cssPath)
	if err != nil {
		log.Error(err)
	}

	re := regexp.MustCompile(`[^\d.]`)
	exrate, err := strconv.ParseFloat(re.ReplaceAllString(rtn, ""), 64)
	if err != nil {
		return 0
	}

	return exrate // TODO 환율 크롤링
}

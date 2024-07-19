package crawl

import (
	"fmt"
	"io"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func Crawl(url string, cssPath string) (gp string, _ error) {

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

	fmt.Println(doc.Text())

	// Find the elements by the CSS selector
	doc.Find(cssPath).Each(func(i int, s *goquery.Selection) {
		// Extract and print the data
		gp = s.Text()
	})

	return gp, nil
}

// func CrawlChrome(url string, id string) (string, error) {
// 	ctx, cancel := chromedp.NewContext(context.Background())
// 	defer cancel()

// 	// 결과를 저장할 변수
// 	var res string

// 	// 크롤링 작업 수행
// 	err := chromedp.Run(ctx,
// 		chromedp.Navigate(url),
// 		chromedp.WaitVisible(id, chromedp.ByID), // 요소가 렌더링될 때까지 대기
// 		chromedp.Text(id, &res, chromedp.ByID),  // 요소의 텍스트 추출
// 	)
// 	if err != nil {
// 		return "", err
// 	}

// 	return res, nil
// }

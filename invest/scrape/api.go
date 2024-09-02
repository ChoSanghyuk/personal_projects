package scrape

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (s Scraper) CallApi(url string, header map[string]string) (string, error) {

	rtn, err := callApi(url, nil)
	if err != nil {
		return "", err
	}
	var d map[string]any

	err = json.Unmarshal([]byte(rtn[1:len(rtn)-1]), &d)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%f", d["trade_price"]), nil
}

func callApi(url string, header map[string]string) (string, error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("error making request\n%w", err)
	}

	// Add headers to the request
	for k, v := range header {
		req.Header.Add(k, v)
	}

	// Send the request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request\n%w", err)
	}
	defer res.Body.Close()

	// Read the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("error reading body\n%w", err)
	}

	// Print the response body
	return string(body), nil
}

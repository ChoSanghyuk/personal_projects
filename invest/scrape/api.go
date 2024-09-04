package scrape

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// TODO. callapi 합쳐
func (s Scraper) CallApi(url string, header map[string]string) (string, error) {

	var rtn []map[string]any
	err := sendRequest(url, nil, &rtn)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%f", rtn[0]["trade_price"]), nil
}

func sendRequest(url string, header map[string]string, response any) error {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("error making request\n%w", err)
	}

	// Add headers to the request
	for k, v := range header {
		req.Header.Add(k, v)
	}

	// Send the request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request\n%w", err)
	}
	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(response)

}

/*
// Read the response body
body, err := io.ReadAll(res.Body)
if err != nil {
	return "", fmt.Errorf("error reading body\n%w", err)
}

// Print the response body
return string(body), nil
*/

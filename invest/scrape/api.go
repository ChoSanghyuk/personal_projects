package scrape

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// TODO. callapi 합쳐
func (s Scraper) CallApi(url string, header map[string]string) (string, error) {

	var rtn []map[string]any
	err := sendRequest(url, http.MethodGet, nil, nil, &rtn)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%f", rtn[0]["trade_price"]), nil
}

/*
중요!
json.Marshal(nil) results in []byte("null"), which is equivalent to the string "null" in JSON.
This "null" will be sent as the body of the request when using http.NewRequest.
So, if the API expects a non-empty JSON body or no body at all, sending "null" may lead to unexpected behavior. If you want to send an empty body, it would be better to pass nil directly to http.NewRequest() instead of marshalling it.
*/
func sendRequest(url string, method string, header map[string]string, body map[string]string, response any) error {

	// bodyByte, err := json.Marshal(body)
	// if err != nil {
	// 	return fmt.Errorf("error request body marshaling \n%w", err)
	// }

	req, err := http.NewRequest(method, url, nil) // bytes.NewBuffer(bodyByte)
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

	b, _ := io.ReadAll(res.Body)

	// Print the response body
	fmt.Println(string(b))

	return json.NewDecoder(res.Body).Decode(response)

}

func sendRequest2(addr string, method string, header map[string]string, body map[string]string, response any) error {

	client := &http.Client{}

	// bodyByte, err := json.Marshal(body)
	// if err != nil {
	// 	return fmt.Errorf("error request body marshaling \n%w", err)
	// }

	req, err := http.NewRequest(method, addr, nil) // bytes.NewBuffer(bodyByte)
	if err != nil {
		return fmt.Errorf("error making request\n%w", err)
	}

	// Add headers to the request
	for k, v := range header {
		req.Header.Add(k, v)
	}

	// Send the request
	params := url.Values{}
	params.Add("fid_cond_mrkt_div_code", "J")
	params.Add("fid_input_iscd", "005930") // stock_no를 적절히 변경

	req.URL.RawQuery = params.Encode()

	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request\n%w", err)
	}
	defer res.Body.Close()

	b, _ := io.ReadAll(res.Body)

	// Print the response body
	fmt.Println(string(b))

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

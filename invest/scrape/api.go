package scrape

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
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

func (s Scraper) upbitApi(sym string) (float64, error) {

	url := "" // TODO. 여기 부분에 config를 Transmitter로 활용

	var rtn []map[string]string
	err := sendRequest(url, http.MethodGet, nil, nil, &rtn)
	if err != nil {
		return 0, err
	}

	cp, err := strconv.ParseFloat(rtn[0]["trade_price"], 64)
	if err != nil {
		return 0, err
	}

	return cp, nil
}

/*
중요!
request body에 nil을 바로 넣는다면, 빈 json 데이터가 들어감.
하지만 nil을 json.Marshal해서 넣는다면, "null"이라는 json 데이터가 형성.
이는 request body에 nil값을 넣는 것과 다른 결과 초래 할 수 있음
*/
func sendRequest(url string, method string, header map[string]string, body map[string]string, response any) error {

	var rb io.Reader
	if body == nil {
		rb = nil
	} else {
		bodyByte, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("error request body marshaling \n%w", err)
		}
		rb = bytes.NewBuffer(bodyByte)
	}

	req, err := http.NewRequest(method, url, rb)
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

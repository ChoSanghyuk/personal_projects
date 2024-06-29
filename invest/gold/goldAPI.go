package gold

import (
	"fmt"
	"invest/config"
	"io"
	"log"
	"net/http"
)

func CallGoldApi() {

	req, err := http.NewRequest("GET", config.ConfigInfo.GoldConfig.API.Url, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Add headers to the request
	req.Header.Add("x-access-token", config.ConfigInfo.GoldConfig.API.ApiKey)

	// Send the request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	// Read the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Print the response body
	fmt.Println(string(body))
}

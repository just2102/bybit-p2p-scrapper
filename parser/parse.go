package parser

import (
	"bytes"
	"fmt"
	"goScrapper/data"
	"goScrapper/utils"
	"io"
	"net/http"
)

const url = "https://api2.bybit.com/fiat/otc/item/online"

func ParseAndDecode(rawPayload data.Payload) data.AdsResult {
	payload, encodeErr := utils.Encode(rawPayload)
	if encodeErr != nil {
		fmt.Println("Failed to encode payload!", encodeErr)
		panic(encodeErr)
	}

	resp, requestErr := http.Post(
		url,
		"application/json",
		bytes.NewBuffer(payload),
	)
	if requestErr != nil {
		fmt.Println("Failed to send request!", requestErr)
		panic(requestErr)
	}

	defer resp.Body.Close()

	// contains 10 items (or less)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read response body!", err)
		panic(err)
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Request failed with status code", resp.StatusCode)
		fmt.Println("Response body:", string(body))
		panic("Request failed")
	}

	var response data.Response

	if decodeErr := utils.Decode(body, &response); decodeErr != nil {
		fmt.Println("Failed to decode response!", decodeErr)
		panic(decodeErr)
	}

	return response.Result
}

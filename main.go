package main

import (
	"fmt"
	"goScrapper/data"
	"goScrapper/parser"
	"goScrapper/utils"
)

var testPayload = data.Payload{
	CurrencyId: data.VND,
	TokenId:    data.USDT,
	Side:       data.BUY,
	Size:       "",
}

func main() {

	response := parser.ParseAndDecode(testPayload)

	if response.Count <= 10 {
		utils.WriteItemsToFile(response.Items)
		return
	}

	newPayload := data.Payload{
		CurrencyId: testPayload.CurrencyId,
		TokenId:    testPayload.TokenId,
		Side:       testPayload.Side,
		Size:       fmt.Sprint(response.Count),
	}

	response2 := parser.ParseAndDecode(newPayload)

	utils.WriteItemsToFile(response2.Items)
}

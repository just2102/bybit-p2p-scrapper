package utils

import (
	"encoding/json"
	"fmt"
	"goScrapper/data"
	"os"
)

func WriteItemsToFile(items []data.Item) {
	file, err := os.Create("items.json")
	if err != nil {
		fmt.Println("Failed to create file:", err)
		return
	}
	defer file.Close()

	data, err := json.MarshalIndent(items, "", "  ")
	if err != nil {
		fmt.Println("Failed to marshal items:", err)
		return
	}

	if _, err := file.Write(data); err != nil {
		fmt.Println("Failed to write items to file:", err)
	}
}

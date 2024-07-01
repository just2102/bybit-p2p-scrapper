package utils

import "encoding/json"

func Encode(body interface{}) ([]byte, error) {
	return json.Marshal(body)
}

package jsonparse

import (
	"encoding/json"

	"github.com/buger/jsonparser"
)

func Set(data, value interface{}, keys ...string) ([]byte, error) {
	dataBytes, _ := json.Marshal(data)
	valueBytes, _ := json.Marshal(value)
	return jsonparser.Set(dataBytes, valueBytes, keys...)
}
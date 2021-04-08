package jsonparse

import (
	"encoding/json"

	"github.com/buger/jsonparser"
)

func Set(data, value interface{}, keys ...string) error {
	dataBytes, _ := json.Marshal(data)
	valueBytes, _ := json.Marshal(value)
	_, err := jsonparser.Set(dataBytes, valueBytes, keys...)
	
	return err
}
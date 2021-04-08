package jsonparse

import (
	"encoding/json"

	"github.com/buger/jsonparser"
)

func Set(data interface{}, value interface{}, keys ...string) error {
	dataBytes, _ := json.Marshal(data)
	valueBytes, _ := json.Marshal(value)

	bytes, err := jsonparser.Set(dataBytes, valueBytes, keys...)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, data)
	return err
}
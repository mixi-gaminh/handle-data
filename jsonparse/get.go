package jsonparse

import (
	"encoding/json"

	"github.com/buger/jsonparser"
)

func GetString(data interface{}, keys ...string) (string, error) {
	bytes, _ := json.Marshal(data)
	return jsonparser.GetString(bytes, keys...)
}

func GetUnsafeString(data interface{}, keys ...string) (string, error) {
	bytes, _ := json.Marshal(data)
	return jsonparser.GetUnsafeString(bytes, keys...)
}

func GetInt(data interface{}, keys ...string) (int64, error) {
	bytes, _ := json.Marshal(data)
	return jsonparser.GetInt(bytes, keys...)
}

func GetFloat(data interface{}, keys ...string) (float64, error) {
	bytes, _ := json.Marshal(data)
	return jsonparser.GetFloat(bytes, keys...)
}

func GetBoolean(data interface{}, keys ...string) (bool, error) {
	bytes, _ := json.Marshal(data)
	return jsonparser.GetBoolean(bytes, keys...)
}
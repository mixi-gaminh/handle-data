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

func GetMapString(data interface{}, keys ...string) (map[string]interface{}, error) {
	bytes, _ := json.Marshal(data)
	val, err :=  jsonparser.GetUnsafeString(bytes, keys...)
	if err != nil {
		return nil, err
	}

	ret := make(map[string]interface{})
	err = json.Unmarshal([]byte(val), &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func GetSlice(data interface{}, keys ...string) ([]interface{}, error) {
	bytes, _ := json.Marshal(data)
	val, err :=  jsonparser.GetUnsafeString(bytes, keys...)
	if err != nil {
		return nil, err
	}

	var ret []interface{}
	err = json.Unmarshal([]byte(val), &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
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
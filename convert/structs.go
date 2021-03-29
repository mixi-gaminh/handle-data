package convert

import (
	"encoding/json"
	"errors"
)

func ToStruct(data interface{}, ret interface{}) error {
	bytes, _ := json.Marshal(data)
	err := json.Unmarshal(bytes, ret)
	if err != nil {
		return errors.New("Json To Struct failed: " + err.Error())
	}
	return nil
}

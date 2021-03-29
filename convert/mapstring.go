package convert

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

func ToMapString(data interface{}) (map[string]interface{}, error) {
	ret, ok := data.(map[string]interface{})
	if !ok {
		typeData := ""
		typeData = fmt.Sprintf("%v", reflect.TypeOf(data))
		return nil, errors.New(`data is ` + typeData + ` not map[string]interface{}`)
	}
	return ret, nil
}

func ToSliceMapString(data interface{}) ([]map[string]interface{}, error) {
	list, ok := data.([]interface{})
	if !ok {
		typeData := ""
		typeData = fmt.Sprintf("%v", reflect.TypeOf(data))
		return nil, errors.New(`data is ` + typeData + ` not []map[string]interface{}`)
	}

	var ret []map[string]interface{}
	for index, val := range list {
		t, ok := val.(map[string]interface{})
		if !ok {
			typeData := ""
			typeData = fmt.Sprintf("%v", reflect.TypeOf(t))
			return nil, errors.New(`data[` + strconv.Itoa(index) + `] is ` + typeData + ` not map[string]interface{}`)
		}
		ret = append(ret, t)
	}

	return ret, nil
}
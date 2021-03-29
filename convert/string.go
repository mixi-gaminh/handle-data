package convert

import (
	"errors"
	"fmt"
	"reflect"
)

func ToString(data interface{}, force bool) (string, error) {
	var ret string
	var ok bool
	if !force {
		ret, ok = data.(string)
		if !ok {
			typeData := fmt.Sprintf("%v", reflect.TypeOf(data))
			return "", errors.New(`data is ` + typeData + ` not string`)
		}
	} else {
		ret = fmt.Sprintf("%v", data)
	}

	return ret, nil
}

func ToSliceString(data interface{}) ([]string, error) {
	ret, ok := data.([]string)
	if !ok {
		typeData := fmt.Sprintf("%v", reflect.TypeOf(data))
		return nil, errors.New(`data is ` + typeData + ` not []string`)
	}

	return ret, nil
}

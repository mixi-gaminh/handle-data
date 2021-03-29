package convert

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

func ToInt(data interface{}) (int, error) {
	str := fmt.Sprintf("%v", data)
	ret, err := strconv.Atoi(str)
	if err != nil {
		typeData := ""
		typeData = fmt.Sprintf("%v", reflect.TypeOf(data))
		return 0, errors.New(`data is ` + typeData + ` not int`)
	}
	return ret, nil
}

func ToInt64(data interface{}) (int64, error) {
	str := fmt.Sprintf("%v", data)
	ret, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		typeData := fmt.Sprintf("%v", reflect.TypeOf(data))
		return 0, errors.New(`data is ` + typeData + ` not int64`)
	}
	return ret, nil
}

func ToFloat64(data interface{}) (float64, error) {
	str := fmt.Sprintf("%v", data)
	ret, err := strconv.ParseFloat(str, 10)
	if err != nil {
		typeData := ""
		typeData = fmt.Sprintf("%v", reflect.TypeOf(data))
		return 0, errors.New(`data is ` + typeData + ` not float64`)
	}
	return ret, nil
}

func ToSliceInt(data interface{}) ([]int, error) {
	ret, ok := data.([]int)
	if !ok {
		typeData := ""
		typeData = fmt.Sprintf("%v", reflect.TypeOf(data))
		return nil, errors.New(`data is ` + typeData + ` not []int`)
	}
	return ret, nil
}

func ToSliceInt64(data interface{}) ([]int64, error) {
	ret, ok := data.([]int64)
	if !ok {
		typeData := ""
		typeData = fmt.Sprintf("%v", reflect.TypeOf(data))
		return nil, errors.New(`data is ` + typeData + ` not []int64`)
	}
	return ret, nil
}

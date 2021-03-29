package mapstring

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/mixi-gaming/handle-data/convert"
)

func GetInt(data map[string]interface{}, keys ...string) (int, error) {
	past := ""
	var ok bool

	for i, query := range keys {
		past += query
		t := data[query]
		if i == len(keys)-1 {
			return convert.ToInt(t)
		}
		data, ok = t.(map[string]interface{})
		if !ok {
			typeData := fmt.Sprintf("%v", reflect.TypeOf(t))
			return 0, errors.New(`data["` + past + `"] is ` + typeData + ` not int`)
		}
		past += "."
	}

	return 0, errors.New(`data["` + past + `"] is not int`)
}

func GetInt64(data map[string]interface{}, keys ...string) (int64, error) {
	past := ""
	var ok bool

	for i, query := range keys {
		past += query
		t := data[query]
		if i == len(keys)-1 {
			return convert.ToInt64(t)
		}
		data, ok = t.(map[string]interface{})
		if !ok {
			typeData := fmt.Sprintf("%v", reflect.TypeOf(t))
			return 0, errors.New(`data["` + past + `"] is ` + typeData + ` not int64`)
		}
		past += "."
	}

	return 0, errors.New(`data["` + past + `"] is not int64`)
}

func GetString(data map[string]interface{}, keys ...string) (string, error) {
	past := ""
	var ok bool

	for i, query := range keys {
		past += query
		t := data[query]
		if i == len(keys)-1 {
			return convert.ToString(t, false)
		}
		data, ok = t.(map[string]interface{})
		if !ok {
			typeData := fmt.Sprintf("%v", reflect.TypeOf(t))
			return "", errors.New(`data["` + past + `"] is ` + typeData + ` not map[string]interface{}`)
		}
		past += "."
	}

	return "", errors.New(`data["` + past + `"] is not string`)
}

func GetUnsafeString(data map[string]interface{}, keys ...string) (string, error) {
	past := ""
	var ok bool

	for i, query := range keys {
		past += query
		t := data[query]
		if i == len(keys)-1 {
			return convert.ToString(t, true)
		}
		data, ok = t.(map[string]interface{})
		if !ok {
			typeData := fmt.Sprintf("%v", reflect.TypeOf(t))
			return "", errors.New(`data["` + past + `"] is ` + typeData + ` not map[string]interface{}`)
		}
		past += "."
	}

	return "", errors.New(`data["` + past + `"] is not string`)
}

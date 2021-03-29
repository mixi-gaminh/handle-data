package mapstring

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/mixi-gaming/handle-data/convert"
)

func GetInt(data map[string]interface{}, queryPath string) (int, error) {
	strArr := strings.Split(queryPath, ".")
	past := ""
	var ok bool

	for i, query := range strArr {
		past += query
		t := data[query]
		if i == len(strArr)-1 {
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

func GetInt64(data map[string]interface{}, queryPath string) (int64, error) {
	strArr := strings.Split(queryPath, ".")
	past := ""
	var ok bool

	for i, query := range strArr {
		past += query
		t := data[query]
		if i == len(strArr)-1 {
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

func GetString(data map[string]interface{}, queryPath string, force bool) (string, error) {
	strArr := strings.Split(queryPath, ".")
	past := ""
	var ok bool

	for i, query := range strArr {
		past += query
		t := data[query]
		if i == len(strArr)-1 {
			return convert.ToString(t, force)
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

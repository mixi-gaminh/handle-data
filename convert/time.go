package convert

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

func TimestampToTime(data interface{}) (time.Time, error) {
	str := fmt.Sprintf("%v", data)

	tm, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return time.Time{}, errors.New(`data is not time.Time`)
	}

	ret := time.Unix(tm, 0)
	return ret, nil
}

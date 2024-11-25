package util

import (
	"strconv"
	"time"
)

func GetCurrentTimeInMillis() string {
	return strconv.FormatInt(time.Now().UnixMilli(), 10)
}

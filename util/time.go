package util

import (
	"strconv"
	"time"
)

type utilTime struct {
}

var Time = utilTime{}

func(*utilTime) NowShortDateTime() string {
	n := time.Now()
	return strconv.Itoa(n.Year()) + "-" + strconv.Itoa(int(n.Month())) + "-" + strconv.Itoa(n.Day())

}
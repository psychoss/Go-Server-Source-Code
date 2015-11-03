package util

import (
	"strconv"
	"time"
)
 

func  NowShortDateTime() string {
	n := time.Now()
	return strconv.Itoa(n.Year()) + "-" + strconv.Itoa(int(n.Month())) + "-" + strconv.Itoa(n.Day())

}
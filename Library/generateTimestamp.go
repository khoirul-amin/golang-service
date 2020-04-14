package Library

import (
	"strconv"
	"time"
)

func TimeStamp() string {
	var now = time.Now()

	year := now.Year()
	month := int(now.Month())
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()

	var timeNow string = strconv.Itoa(year) + "-" + strconv.Itoa(month) + "-" + strconv.Itoa(day) + " " + strconv.Itoa(hour) + ":" + strconv.Itoa(minute) + ":" + strconv.Itoa(second)

	return timeNow

}

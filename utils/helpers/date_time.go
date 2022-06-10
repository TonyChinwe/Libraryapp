package helpers

import (
	"time"
)

// UnixMilli use to get milliseconds of given time
// @params - time
// return - milliseconds of the given time
func UnixMilli(t time.Time) int64 {
	return t.Round(time.Millisecond).UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
}

// CurrentTimeInMilli use to get current time in milliseconds
// This function will use when we need current timestamp
// This functions return current timestamp (in millisecods)
func CurrentTimeInMilli() int64 {
	return UnixMilli(time.Now())
}

func CurrentDateTimeInString() string {
	return time.Now().UTC().Format("2006-02-01 15:04:05")
}

func CurrentDate() time.Time {
	ctime := time.Now().UTC().Format("2006-02-01 15:04:05")
	currentTime, _ := time.Parse("2006-02-01 15:04:05", ctime)
	return currentTime
}

func ConvertToDate(date string) (time.Time, error) {
	return time.Parse("2006-01-02", date)

}

func CurrentTimeInString() string {
	return time.Now().UTC().Format("15:04:05")
}

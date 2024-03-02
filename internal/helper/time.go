package helper

import (
	"time"
)

const (
	SecondsPerMinute = 60
	MinutesPerHour   = 60
	HoursOffset      = 9
)

var (
	JST            = time.FixedZone("Asia/Tokyo", HoursOffset*MinutesPerHour*SecondsPerMinute)
	JSTDatetime    = "20060102-150405"
	TimeFormat     = "15:04"
	DateTimeFormat = "2006-01-02T15:04:05"
)

type Date struct {
	Year  int
	Month time.Month
	Day   int
}

func ParseDate(date string) time.Time {
	parsedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return time.Time{}
	}
	return parsedDate
}

func FormatDate(date time.Time) string {
	return date.Format("2006-01-02")
}

func FormatDateJP(date time.Time) string {
	return date.Format("2006年1月2日")
}

func FormatTime(date time.Time) string {
	return date.Format(TimeFormat)
}

package entity

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"time"
)

const timeFormat = "15:04"

type Time struct {
	time.Time
}

// Scan dbのTIME型をtime.Timeに変換するラッパー関数
func (t *Time) Scan(value interface{}) error {
	if value == nil {
		*t = Time{time.Time{}}
		return nil
	}
	if bv, err := driver.String.ConvertValue(value); err == nil {
		if v, ok := bv.(string); ok {
			newTime, err := time.Parse("15:04:05", v)
			if err != nil {
				return err
			}

			*t = Time{newTime}
			return nil
		}
	}
	return errors.New("failed to scan Time")
}

func (t Time) Value() (driver.Value, error) {
	if t.Time.IsZero() {
		return nil, errors.New("zero time")
	}
	return t.Time.Format(time.TimeOnly), nil
}

func (t *Time) IsValid() bool {
	return t.Time.Format(timeFormat) != "00:00"
}

func (t *Time) FormatTime() string {
	return t.Time.Format(timeFormat)
}

func ParseTime(timeStr string) (Time, error) {
	parsedTime, err := time.Parse(timeFormat, timeStr)
	if err != nil {
		return Time{}, err
	}

	return Time{parsedTime}, nil
}

func FromTime(t time.Time) Time {
	return Time{t}
}

func (t Time) ToTime() time.Time {
	return t.Time
}

// MakeDateSeries makes series of dates, which includes the Day 'to'.
func MakeDateSeries(from, to time.Time) ([]time.Time, error) {
	days := int(to.Sub(from).Hours()/24) + 1 // toの当日を含めるため+1
	if days <= 0 {
		err := fmt.Errorf("'to' must be later than or equal to 'from': from=%v, to=%v", from, to)
		return nil, err
	}

	dates := make([]time.Time, days)
	for i := range dates {
		dates[i] = from.Add(time.Hour * 24 * time.Duration(i))
	}
	return dates, nil
}

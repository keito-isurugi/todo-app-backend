package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimeMethods(t *testing.T) {
	a := assert.New(t)
	tm := &Time{}
	t.Run("scan", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			a = assert.New(t)
			err := tm.Scan("13:45:00")
			a.NoError(err)
		})

		t.Run("invalid", func(t *testing.T) {
			a = assert.New(t)
			err := tm.Scan("invalid")
			a.Error(err)
		})

		t.Run("nil", func(t *testing.T) {
			a = assert.New(t)
			err := tm.Scan(nil)
			a.NoError(err)
		})

		t.Run("failed to scan Time", func(t *testing.T) {
			a = assert.New(t)
			err := tm.Scan(1)
			a.Error(err)
		})
	})

	t.Run("Value", func(t *testing.T) {
		t.Run("zero time", func(t *testing.T) {
			a = assert.New(t)
			tm.Time = time.Time{}
			got, err := tm.Value()
			a.Error(err)
			a.Nil(got)
		})

		t.Run("success", func(t *testing.T) {
			a = assert.New(t)
			tm.Time = time.Date(2000, 1, 1, 13, 45, 45, 0, time.UTC)
			got, err := tm.Value()
			a.NoError(err)
			a.Equal("13:45:45", got)
		})
	})

	t.Run("IsValid", func(t *testing.T) {
		a = assert.New(t)
		t.Run("valid", func(t *testing.T) {
			a = assert.New(t)
			tm.Time = time.Date(2000, 1, 1, 13, 45, 45, 0, time.UTC)
			a.True(tm.IsValid())
		})

		t.Run("invalid", func(t *testing.T) {
			a = assert.New(t)
			tm.Time = time.Time{}
			a.False(tm.IsValid())
		})
	})

	t.Run("FormatTime", func(t *testing.T) {
		a = assert.New(t)
		tm.Time = time.Date(2000, 1, 1, 13, 45, 45, 0, time.UTC)
		expected := "13:45"
		a.Equal(expected, tm.FormatTime())
	})

	t.Run("ParseTime", func(t *testing.T) {
		a = assert.New(t)
		t.Run("success", func(t *testing.T) {
			a = assert.New(t)
			got, err := ParseTime("13:45")
			a.NoError(err)
			expected := Time{time.Date(0, 1, 1, 13, 45, 0, 0, time.UTC)}
			a.Equal(expected, got)
		})

		t.Run("invalid", func(t *testing.T) {
			a = assert.New(t)
			got, err := ParseTime("invalid")
			a.Error(err)
			a.Equal(Time{}, got)
		})
	})

	t.Run("MakeDateSeries", func(t *testing.T) {
		a = assert.New(t)
		t.Run("success", func(t *testing.T) {
			a = assert.New(t)
			from := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
			to := time.Date(2020, 1, 5, 0, 0, 0, 0, time.UTC)
			got, err := MakeDateSeries(from, to)
			a.NoError(err)
			expected := []time.Time{
				time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
				time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2020, 1, 3, 0, 0, 0, 0, time.UTC),
				time.Date(2020, 1, 4, 0, 0, 0, 0, time.UTC),
				time.Date(2020, 1, 5, 0, 0, 0, 0, time.UTC),
			}
			a.Equal(expected, got)
		})

		t.Run("invalid/fromがtoより前の日付", func(t *testing.T) {
			a = assert.New(t)
			from := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
			to := time.Date(2019, 12, 31, 0, 0, 0, 0, time.UTC)
			got, err := MakeDateSeries(from, to)
			a.Error(err)
			a.Nil(got)
		})
	})
}

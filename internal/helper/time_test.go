package helper_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/keito-isurugi/todo-app-backend/internal/helper"
)

func TestParseDate(t *testing.T) {
	testCases := []struct {
		name        string
		input       string
		expected    time.Time
		expectError bool
	}{
		{
			name:     "valid date",
			input:    "2023-07-02",
			expected: time.Date(2023, 7, 2, 0, 0, 0, 0, time.UTC),
		},
		{
			name:        "invalid date",
			input:       "invalid-date",
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := helper.ParseDate(tc.input)

			if tc.expectError {
				assert.Equal(t, time.Time{}, output)
			} else {
				assert.Equal(t, tc.expected, output)
			}
		})
	}
}

func TestFormatDate(t *testing.T) {
	testCases := []struct {
		name     string
		input    time.Time
		expected string
	}{
		{
			name:     "valid date",
			input:    time.Date(2023, 7, 2, 0, 0, 0, 0, time.UTC),
			expected: "2023-07-02",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output := helper.FormatDate(tc.input)
			assert.Equal(t, tc.expected, output)
		})
	}
}

func TestFormatDateJP(t *testing.T) {
	a := assert.New(t)
	helper.FormatDateJP(time.Date(2023, 7, 2, 0, 0, 0, 0, time.UTC))
	a.Equal("2023年7月2日", helper.FormatDateJP(time.Date(2023, 7, 2, 0, 0, 0, 0, time.UTC)))
}

func TestFormatTime(t *testing.T) {
	a := assert.New(t)
	helper.FormatTime(time.Date(2023, 7, 2, 0, 0, 0, 0, time.UTC))
	a.Equal("00:00", helper.FormatTime(time.Date(2023, 7, 2, 0, 0, 0, 0, time.UTC)))
}

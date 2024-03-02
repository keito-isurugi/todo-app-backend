package handler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	TestBool Bool   `json:"test_bool"`
	TestTime string `json:"test_time" validate:"isTime"`
}

func TestUnmarshalJSON(t *testing.T) {
	a := assert.New(t)

	testData := []struct {
		jsonString string
		expected   Bool
		shouldErr  bool
	}{
		{jsonString: `"true"`, expected: true, shouldErr: false},
		{jsonString: `"false"`, expected: false, shouldErr: false},
		{jsonString: `"0"`, expected: false, shouldErr: false},
		{jsonString: `"1"`, expected: true, shouldErr: false},
		{jsonString: `"invalid"`, expected: false, shouldErr: true},
	}

	for _, td := range testData {
		var b Bool
		err := b.UnmarshalJSON([]byte(td.jsonString))

		if td.shouldErr {
			a.Error(err)
		} else {
			a.NoError(err)
			a.Equal(td.expected, b)
		}
	}
}

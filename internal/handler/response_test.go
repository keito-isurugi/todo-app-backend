package handler

import (
	"reflect"
	"testing"
)

func TestEmptyResponse(t *testing.T) {
	empty := emptyResponse{}
	t.Run("emptyResponse", func(t *testing.T) {
		numFields := reflect.TypeOf(empty).NumField()
		if numFields != 0 {
			t.Errorf("emptyResponse has %v fields, want 0", numFields)
		}
	})
}

func TestCreatedResponse(t *testing.T) {
	id := "test_id"
	created := newCreatedResponse(id)
	t.Run("newCreatedResponse", func(t *testing.T) {
		if created.ID != id {
			t.Errorf("got %v for ID, want %v", created.ID, id)
		}
	})
}

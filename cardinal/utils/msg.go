package utils

import (
	"encoding/json"
	"errors"
	"net/http"
	"reflect"
)

func DecodeMsg[T any](r *http.Request, msg *T) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	// Try to decode the request body into struct T.
	// Errors if msg has unknown fields.
	err := decoder.Decode(&msg)
	if err != nil {
		return err
	}

	// Check that all fields are present
	fields := reflect.ValueOf(msg).Elem()
	for i := 0; i < fields.NumField(); i++ {
		if fields.Field(i).IsZero() {
			return errors.New("some msg field(s) are missing")
		}
	}

	return nil
}

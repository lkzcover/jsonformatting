package jsonformatting

import (
	"encoding/json"
	"errors"
)

var ErrJsonFormat = errors.New("jsonformatting: incorrect json format")

// ConvertToFormatJSON if the incoming data is JSON, the function returns a formatted slice byte.
// Otherwise it returns the original slice.
func ConvertToFormatJSON(data []byte) []byte {

	if !json.Valid(data) {
		return data
	}

	return handler(data)
}

// ConvertToFormatJSONWithError if the incoming data is JSON, the function returns a formatted slice byte.
// Otherwise it returns error.
func ConvertToFormatJSONWithError(data []byte) ([]byte, error) {
	if !json.Valid(data) {
		return nil, ErrJsonFormat
	}

	return handler(data), nil
}

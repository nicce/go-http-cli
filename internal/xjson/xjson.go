package xjson

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// PrettyFormat - format a byte array in to a json string. Returns an error if the byte array can't be formatted as json.
func PrettyFormat(b []byte) (string, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "  ")

	if err != nil {
		return "", fmt.Errorf("error formatting byte array: %w", err)
	}

	return out.String(), nil
}

package xjson

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// PrettyFormat - format a byte array in to a json string. Returns an error if the byte array can't be formatted as json.
func PrettyFormat(b []byte, compact bool) (string, error) {
	var out bytes.Buffer

	if compact {
		err := json.Compact(&out, b)
		if err != nil {
			return "", fmt.Errorf("error compacting byte array: %w", err)
		}

		return out.String(), nil
	}

	err := json.Indent(&out, b, "", "  ")
	if err != nil {
		return "", fmt.Errorf("error indenting byte array: %w", err)
	}

	return out.String(), nil
}

package xjson_test

import (
	"go-http-cli/internal/xjson"
	"testing"
)

// TestPrettyFormatCompactedJson - verifies the json string is compacted.
func TestPrettyFormatCompactedJson(t *testing.T) {
	t.Parallel()
	// arrange
	jsonString := "{\n\"a\": 1,\n\"b\": 2\n}"

	// act
	got, err := xjson.PrettyFormat([]byte(jsonString), true)
	if err != nil {
		t.Fatalf("got error: %v", err)
	}

	// arrange
	want := "{\"a\":1,\"b\":2}"
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

// TestPrettyFormatIndentJson - verifies the json string is indented.
func TestPrettyFormatIndentJson(t *testing.T) {
	t.Parallel()
	// arrange
	jsonString := "{\n\"a\": 1,\n\"b\": 2\n}"

	// act
	got, err := xjson.PrettyFormat([]byte(jsonString), false)
	if err != nil {
		t.Fatalf("got error: %v", err)
	}

	// arrange
	want := "{\n  \"a\": 1,\n  \"b\": 2\n}"
	if got != want {
		t.Errorf("got %s \n want %s", got, want)
	}
}

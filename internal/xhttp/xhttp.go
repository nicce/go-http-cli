package xhttp

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
)

// Call - retrieves the body from the given URL.
func Call(ctx context.Context, url string, method string, body []byte) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error from http client: %w", err)
	}

	defer func() { _ = res.Body.Close() }()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading body: %w", err)
	}

	return b, nil
}

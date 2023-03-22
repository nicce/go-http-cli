package xhttp

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

// HttpMethod - defines the HttpMethod type.
type HttpMethod string

const (
	// Get - Defines the GET HttpMethod.
	Get HttpMethod = "GET"
	// Post - Defines the POST HttpMethod.
	Post = "POST"
)

// HttpResponse - contains information about the http response.
type HttpResponse struct {
	Body        []byte
	Headers     map[string][]string
	LatencyInMs int64
	Status      int
}

// Call - retrieves the body from the given URL.
func Call(ctx context.Context, url string, method HttpMethod, headers map[string]string, body []byte) (*HttpResponse, error) {
	client := &http.Client{}

	req, err := http.NewRequestWithContext(ctx, string(method), url, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	start := time.Now()

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error from http client: %w", err)
	}

	defer func() { _ = res.Body.Close() }()

	latency := time.Since(start).Milliseconds()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading body: %w", err)
	}

	return &HttpResponse{
		Body:        b,
		Headers:     res.Header,
		LatencyInMs: latency,
		Status:      res.StatusCode,
	}, nil
}

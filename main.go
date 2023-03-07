package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "Http Client"
	app.Description = "Your cURL replacement"
	app.Version = "v0.0.1"
	app.Action = func(cCtx *cli.Context) error {
		rawURL := cCtx.Args().Get(0)

		b, err := Execute(cCtx.Context, rawURL)
		if err != nil {
			return fmt.Errorf("error executing action: %w", err)
		}

		var out bytes.Buffer
		err = json.Indent(&out, b, "", "\t")

		if err != nil {
			fmt.Print(string(b))
		}

		fmt.Print(out)

		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

// Execute - retrieves the response from the url or error if occurred.
func Execute(ctx context.Context, url string) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
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

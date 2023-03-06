package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/urfave/cli/v2"
	"io"
	"log"
	"net/http"
	"os"
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
			return err
		}
		var out bytes.Buffer
		err = json.Indent(&out, b, "", "\t")
		if err != nil {
			fmt.Print(string(b))
		}

		fmt.Print(out.String())
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func Execute(ctx context.Context, url string) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return io.ReadAll(res.Body)
}

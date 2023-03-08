package main

import (
	"fmt"
	"go-http-cli/internal/xhttp"
	"go-http-cli/internal/xjson"
	"log"
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

		b, err := xhttp.Call(cCtx.Context, rawURL, xhttp.Get, nil)
		if err != nil {
			return fmt.Errorf("error executing action: %w", err)
		}

		out, err := xjson.PrettyFormat(b)
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

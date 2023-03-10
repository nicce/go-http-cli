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
	var compact bool

	app := cli.NewApp()
	app.Name = "Http Client"
	app.Description = "Your cURL replacement"
	app.Version = "v0.0.1"
	app.Flags = []cli.Flag{
		&cli.BoolFlag{
			Name:        "compact",
			Value:       false,
			Usage:       "use to get compact json",
			Destination: &compact,
		},
	}
	app.Action = func(cCtx *cli.Context) error {
		rawURL := cCtx.Args().Get(0)

		res, err := xhttp.Call(cCtx.Context, rawURL, xhttp.Get, nil)
		if err != nil {
			return fmt.Errorf("error executing action: %w", err)
		}

		out, err := xjson.PrettyFormat(res.Body, compact)
		if err != nil {
			fmt.Printf("%d ms \n %v", res.LatencyInMs, string(res.Body))
		}

		fmt.Printf("request took: %dms \n %v", res.LatencyInMs, out)

		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

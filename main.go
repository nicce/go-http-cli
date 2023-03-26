package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/nicce/go-http-cli/internal/version"
	"github.com/nicce/go-http-cli/internal/xhttp"
	"github.com/nicce/go-http-cli/internal/xjson"

	"github.com/urfave/cli/v2"
)

const (
	red    = "\033[31m"
	white  = "\033[97m"
	green  = "\033[32m"
	cyan   = "\033[36m"
	yellow = "\033[33m"
)

func main() {
	var compact bool

	var include bool

	var requestMethod string

	var headers cli.StringSlice

	var body string

	app := cli.NewApp()
	app.Name = "Http Client"
	app.Description = "Your cURL replacement"
	app.Version = version.Version
	app.Flags = []cli.Flag{
		&cli.BoolFlag{
			Name:        "compact",
			Aliases:     []string{"c"},
			Usage:       "use to get compact json",
			Destination: &compact,
		},
		&cli.BoolFlag{
			Name:        "include",
			Aliases:     []string{"i"},
			Usage:       "use to include response headers",
			Destination: &include,
		},
		&cli.StringFlag{
			Name:        "request",
			Aliases:     []string{"r"},
			Usage:       "use to specify http request method",
			Destination: &requestMethod,
			Value:       string(xhttp.Get),
		},
		&cli.StringSliceFlag{
			Name:        "headers",
			Aliases:     []string{"H"},
			Usage:       "use to specify http request headers, like -H 'content-type: application/json'",
			Destination: &headers,
		},
		&cli.StringFlag{
			Name:        "data",
			Aliases:     []string{"d"},
			Usage:       "use to specify http json body",
			Destination: &body,
		},
	}
	app.Action = func(cCtx *cli.Context) error {
		rawURL := cCtx.Args().Get(0)

		h := parseHeaders(headers.Value())

		res, err := xhttp.Call(cCtx.Context, rawURL, xhttp.HttpMethod(requestMethod), h, body)
		if err != nil {
			return fmt.Errorf("error executing action: %w", err)
		}

		out, err := xjson.PrettyFormat(res.Body, compact)
		if err != nil {
			fmt.Printf("%d ms \n %v", res.LatencyInMs, string(res.Body))
		}

		fmt.Printf("%srequest latency: %s%dms \n", yellow, cyan, res.LatencyInMs)

		if include {
			for k, h := range res.Headers {
				fmt.Printf("%s%s: %s%v\n", red, k, white, h)
			}

			fmt.Printf("%s%s: %s%v\n", red, "Status", white, res.Status)
		}

		fmt.Printf(green + out)

		return nil
	}

	sort.Sort(cli.FlagsByName(app.Flags))

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func parseHeaders(headers []string) map[string]string {
	h := make(map[string]string)

	for _, s := range headers {
		b := strings.Split(s, ":")
		h[b[0]] = strings.TrimSpace(b[1])
	}

	return h
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"

	"github.com/atotto/clipboard"
	"github.com/urfave/cli"
)

var Flags = []cli.Flag{
	cli.BoolFlag{Name: "clip, c", Usage: "copy to clipboard"},
}

func init() {
    cli.AppHelpTemplate = `NAME:
   {{.Name}} - {{.Usage}}

USAGE:
   {{.HelpName}} {{if .VisibleFlags}}[options]{{end}}{{if .ArgsUsage}}{{.ArgsUsage}}{{else}} [arguments...]{{end}}
   {{if len .Authors}}
AUTHOR:
   {{range .Authors}}{{ . }}{{end}}
   {{end}}{{if .VisibleFlags}}
OPTIONS:
   {{range .VisibleFlags}}{{.}}
   {{end}}{{end}}{{if .Copyright }}
COPYRIGHT:
   {{.Copyright}}
   {{end}}{{if .Version}}
VERSION:
   {{.Version}}
   {{end}}
`
}

func Action(c *cli.Context) {

	if c.NArg() < 1 {
		fmt.Println("A url parameter is required.")
		os.Exit(1)
	}
	url := c.Args().Get(0)
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find("title,h1,h2").Each(func(i int, s *goquery.Selection) {
		title := s.Text()
		mdlink := fmt.Sprintf("[%s](%s)\n", strings.TrimSpace(title), url)
		if c.Bool("clip") {
			if err := clipboard.WriteAll(mdlink); err != nil {
				os.Exit(1)
			}
		}
	})
}

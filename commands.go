package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/urfave/cli"
)

var Flags = []cli.Flag{}

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
	if h1 := doc.Find("h1").First(); h1 != nil {
		fmt.Printf("[%s](%s)", h1.Text(), url)
	}
}

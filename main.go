package main

import (
	"fmt"
	"net/http"
	"os"
    "log"
    "github.com/PuerkitoBio/goquery"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("A url parameter is required.")
		os.Exit(1)
	}
    url := os.Args[1]
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

package main

import (
	"strings"
)

func SelectTop(links []string) string {
	return links[0]
}

func SelectAll(links []string) string {
	return strings.Join(links, "\n")
}

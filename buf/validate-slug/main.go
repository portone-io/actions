package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/mod/module"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: validate-slug <slug>")
		os.Exit(1)
	}
	slug := os.Args[1]
	if strings.ContainsAny(slug, "/\\") {
		fmt.Fprintln(os.Stderr, "slug must be a single directory name")
		os.Exit(1)
	}
	if err := module.CheckPath("github.com/portone-io/go/interface/" + slug); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

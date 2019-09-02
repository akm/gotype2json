package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/akm/gotype2json"
)

func main() {
	tm := gotype2json.TypeMap{}
	tm.Start(
		(*http.Request)(nil),
		(*http.Response)(nil),
	)

	if b, err := json.MarshalIndent(tm, "", "  "); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to marshal as JSON because of [%T] %v", err, err)
		os.Exit(1)
	} else {
		if _, err := os.Stdout.Write(b); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to output because of [%T] %v", err, err)
			os.Exit(1)
		}
	}
}

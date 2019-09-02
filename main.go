package main

import (
	"fmt"
	"os"

	"encoding/json"
	"net/http"
)

func main() {
	tm := TypeMap{}
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

package main

import (
	"fmt"
	"os"
)

var (
	TheVersion = "0.0.0-alpha"
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

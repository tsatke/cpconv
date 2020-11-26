package main

import (
	"os"

	"github.com/tsatke/cpconv"
	"golang.org/x/text/encoding/charmap"
)

/*
cd into this directory, then invoke go run .
*/

func main() {
	from, err := os.Open("myEbcdicFile.txt")
	if err != nil {
		panic(err)
	}

	if err := cpconv.Convert(from, charmap.CodePage037, os.Stdout, charmap.Windows1252); err != nil {
		panic(err)
	}
}

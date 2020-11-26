package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tsatke/cpconv"
	"golang.org/x/text/encoding/charmap"
)

var (
	rootCmd = &cobra.Command{
		Use:   "cpconv",
		Short: "Convert stdin from one codepage to another",
		Long: `Read input from stdin and interpret it as the --from codepage.
The interpreted bytes are the converted to the --to codepage.
The result is written to stdout.`,
		Example: "cat myFile.txt | cpconv --from IBM037 --to CP1252",
		Version: TheVersion,
		Run: func(cmd *cobra.Command, args []string) {
			if err := execute(); err != nil {
				_, _ = fmt.Fprintln(os.Stderr, err.Error())
			}
		},
	}
)

// flags
var (
	from string
	to   string
)

func init() {
	rootCmd.PersistentFlags().StringVar(&from, "from", "", "--from IBM037")
	rootCmd.PersistentFlags().StringVar(&to, "to", "", "--to CP1252")
	_ = rootCmd.MarkPersistentFlagRequired("from")
	_ = rootCmd.MarkPersistentFlagRequired("to")
}

var (
	cpNames = map[string]cpconv.Codepage{
		"CP1252":  charmap.Windows1252,
		"CP-1252": charmap.Windows1252,
		"IBM037":  charmap.CodePage037,
		"IBM-037": charmap.CodePage037,
	}
)

func execute() error {
	fromCp, ok := cpNames[from]
	if !ok {
		return fmt.Errorf("no codepage for name %s", from)
	}
	toCp, ok := cpNames[to]
	if !ok {
		return fmt.Errorf("no codepage for name %s", to)
	}

	return cpconv.Convert(os.Stdin, fromCp, os.Stdout, toCp)
}

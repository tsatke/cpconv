
<p align="center"></p>
<p align="center">
    <h1 align="center">cpconv</h1>
    <p align="center">A codepage converter. Has a Go API as well as a command line tool. Use it to convert streams from one codepage into another.</p>
    <p align="center">
        <a href="https://github.com/tsatke/cpconv/actions"><img src="https://github.com/tsatke/cpconv/workflows/Build/badge.svg"></a>
        <a href="https://github.com/tsatke/cpconv/actions"><img src="https://github.com/tsatke/cpconv/workflows/Tests/badge.svg"></a>
        <a href="https://github.com/tsatke/cpconv/actions"><img src="https://github.com/tsatke/cpconv/workflows/Static%20analysis/badge.svg"></a>
        <br>
        <img src="https://img.shields.io/badge/status-WIP-yellow">
    </p>
</p>

---

## Installation
Go get it with
```plain
go get github.com/tsatke/cpconv/... 
```
or just import it with
```plain
import "github.com/tsatke/cpconv"
```
and see a usage example below

## Usage (as command line tool)
```plain
Usage:
  cpconv [flags]

Examples:
cat myFile.txt | cpconv --from IBM037 --to CP1252

Flags:
      --from string   --from IBM037
  -h, --help          help for cpconv
      --to string     --to CP1252
      --version       version for cpconv
```

This tool does not read from files, it just reads from stdin and writes to stdout.
Pipes are your best friends with this tool, as you can see in the example.

## Supported codepages
| Codepage | supported | aliases |
| --- | :---: | --- |
| `IBM037` | :white_check_mark: | `IBM-037` |
| `CP1252` | :white_check_mark: | `CP-1252` |

## Usage (with the Go API)
**Please note** that there are no restrictions on what codepages are supported when using the Go API.
This is because we are not limited to strings here.
As long as you can provide an implementation, the conversion will work.

Working example:
```go
package main

import (
    "os"

    "github.com/tsatke/cpconv"
    "golang.org/x/text/encoding/charmap"
)

func main() {
    // this will convert EBCDIC input on stdin to CP1252 output on stdout 
    if err := cpconv.Convert(os.Stdin, charmap.CodePage037, os.Stdout, charmap.Windows1252); err != nil {
        panic(err)
    }
}
```

Please find other examples in the examples directory.

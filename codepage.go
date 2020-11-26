package cpconv

import (
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"
)

// ensure that *charmap.Charmap implements Codepage
var _ Codepage = (*charmap.Charmap)(nil)

type Codepage interface {
	NewDecoder() *encoding.Decoder
	NewEncoder() *encoding.Encoder
}

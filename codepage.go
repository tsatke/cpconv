package cpconv

import (
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"
)

// ensure that *charmap.Charmap implements Codepage
var _ Codepage = (*charmap.Charmap)(nil)

// Codepage is a component that deterministically provides an encoder and decoder.
// It describes the concept of being able to interpret and encode characters from
// and to bytes (codepage, charmap, charset etc.).
type Codepage interface {
	NewDecoder() *encoding.Decoder
	NewEncoder() *encoding.Encoder
}

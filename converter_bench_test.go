package cpconv

import (
	"bytes"
	"strconv"
	"strings"
	"testing"

	"golang.org/x/text/encoding/charmap"
)

var (
	result interface{}
	_      = result
)

func BenchmarkConverter_Convert(b *testing.B) {
	for i := 9; i <= 16; i++ {
		bufsz := 1 << i
		b.Run("bufsz="+strconv.Itoa(bufsz), func(b *testing.B) {
			b.StopTimer()

			conv := Converter{
				FromCodepage: charmap.Windows1252,
				ToCodepage:   charmap.CodePage037,
				BufferSize:   bufsz,
			}

			var err error

			for i := 0; i < b.N; i++ {
				source := strings.NewReader(hugeLipsum)
				var buf bytes.Buffer

				b.StartTimer()
				err = conv.Convert(source, &buf)
				b.StopTimer()

				if err != nil {
					panic(err)
				}
			}

			result = err
		})
	}
}

// BenchmarkConvert will be much slower than e.g. BenchmarkConverter_Convert, since it has to create a new
// converter in every run. This shows, that if a converter is reused often enough, it makes a lot of
// sense to reuse the converter object.
func BenchmarkConvert(b *testing.B) {
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		var err error

		for i := 0; i < b.N; i++ {
			source := strings.NewReader(hugeLipsum)
			var buf bytes.Buffer

			b.StartTimer()
			err = Convert(source, charmap.Windows1252, &buf, charmap.CodePage037)
			b.StopTimer()

			if err != nil {
				panic(err)
			}
		}

		result = err
	}
}

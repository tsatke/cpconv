package cpconv

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
	"golang.org/x/text/encoding/charmap"
)

var (
	originalText = []byte("[some text with a brackets]")
)

func TestConverterSuite(t *testing.T) {
	suite.Run(t, new(ConverterSuite))
}

type ConverterSuite struct {
	suite.Suite

	texts map[Codepage][]byte
}

func (suite *ConverterSuite) SetupSuite() {
	suite.texts = make(map[Codepage][]byte)
	for _, codepage := range []Codepage{
		charmap.Windows1252,
		charmap.CodePage037,
	} {
		encoded, err := codepage.NewEncoder().Bytes(originalText)
		suite.NoError(err, "codepage: %s", codepage)
		suite.texts[codepage] = encoded
	}
}

func (suite *ConverterSuite) TextForCodepage(cp Codepage) []byte {
	text, ok := suite.texts[cp]
	suite.True(ok, "No text for codepage %v", cp)
	return text
}

func (suite *ConverterSuite) CheckConverter(converter Converter) {
	text := suite.TextForCodepage(converter.FromCodepage)
	expectedText := suite.TextForCodepage(converter.ToCodepage)
	source := bytes.NewReader(text)
	var target bytes.Buffer
	err := converter.Convert(source, &target)
	suite.NoError(err)
	suite.Equal(expectedText, target.Bytes())
}

func (suite *ConverterSuite) TestConvertNoop() {
	conv := Converter{
		FromCodepage: charmap.Windows1252,
		ToCodepage:   charmap.Windows1252,
		BufferSize:   1024,
	}
	suite.CheckConverter(conv)
}

func (suite *ConverterSuite) TestZeroBuffer() {
	conv := Converter{
		FromCodepage: charmap.Windows1252,
		ToCodepage:   charmap.Windows1252,
		BufferSize:   0,
	}
	err := conv.Convert(nil, nil) // for this test, we don't need source or target
	suite.EqualError(err, "buffer size can not be 0")
}

func (suite *ConverterSuite) TestConversions() {
	for _, test := range []struct {
		fromCp, toCp Codepage
	}{
		{charmap.Windows1252, charmap.Windows1252},
		{charmap.CodePage037, charmap.Windows1252},
		{charmap.Windows1252, charmap.CodePage037},
	} {
		for i := 9; i <= 12; i++ {
			bufsz := 1 << i
			suite.Run(fmt.Sprintf("from=%s/to=%s/bufsz=%d", test.fromCp, test.toCp, bufsz), suite.createTestFn(test.fromCp, test.toCp, bufsz))
		}
	}
}

func (suite *ConverterSuite) createTestFn(from, to Codepage, bufferSize int) func() {
	return func() {
		conv := Converter{
			FromCodepage: from,
			ToCodepage:   to,
			BufferSize:   bufferSize,
		}
		suite.CheckConverter(conv)
	}
}

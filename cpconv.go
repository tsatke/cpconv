package cpconv

import "io"

func Convert(from io.Reader, sourceCodePage Codepage, to io.Writer, targetCodePage Codepage) error {
	return Converter{
		FromCodepage: sourceCodePage,
		ToCodepage:   targetCodePage,
		BufferSize:   1024,
	}.Convert(from, to)
}

package cpconv

import "io"

// Convert converts the given reader and writes the result to the given writer.
// The given codepages are used for decoding and encoding.
// The buffer size used for this is 1024. If you want to specify a different buffer size,
// use a Converter.
func Convert(from io.Reader, sourceCodePage Codepage, to io.Writer, targetCodePage Codepage) error {
	return Converter{
		FromCodepage: sourceCodePage,
		ToCodepage:   targetCodePage,
		BufferSize:   1024,
	}.Convert(from, to)
}

package cpconv

import (
	"fmt"
	"io"
)

// Converter is an object that can copy a reader to a writer while performing a codepage conversion.
// See Convert for more info.
type Converter struct {
	// FromCodepage is the codepage that is used to decode the reader.
	FromCodepage Codepage
	// ToCodepage is the codepage that will be used to encode decoded bytes from the reader.
	// Bytes of this codepage will be written to the writer.
	ToCodepage Codepage
	// BufferSize is the size of the buffer that will be used for conversion.
	BufferSize int
}

// Convert performs a codepage conversion of bytes from the reader and writes them to the writer.
// If the buffer size of this converter is zero (i.e. not specified), this will fail.
// If the converter uses the same codepage for reading as for writing, the reader will simply
// be copied to the writer, using the buffer size of this converter.
func (c Converter) Convert(from io.Reader, to io.Writer) error {
	// check if buffer size is 0, because if so, this would run indefinitely
	if c.BufferSize == 0 {
		return fmt.Errorf("buffer size can not be 0")
	}

	// optimization: if codepages are the same, just copy input to output
	if c.FromCodepage == c.ToCodepage {
		buffer := make([]byte, c.BufferSize)
		_, err := io.CopyBuffer(to, from, buffer)
		if err != nil {
			return fmt.Errorf("copy: %w", err)
		}
	}

	decoder := c.FromCodepage.NewDecoder()
	encoder := c.ToCodepage.NewEncoder()

	buffer := make([]byte, c.BufferSize)
	exitAfterConversion := false
	for !exitAfterConversion {
		n, err := from.Read(buffer)
		if err == io.EOF {
			// exit after conversion has performed on the last bytes
			// of the reader, AND the converted data has been flushed
			// to the writer
			exitAfterConversion = true
		} else if err != nil {
			return fmt.Errorf("read: %w", err)
		}
		decoded, err := decoder.Bytes(buffer[:n])
		if err != nil {
			return fmt.Errorf("decode: %w", err)
		}
		encoded, err := encoder.Bytes(decoded)
		if err != nil {
			return fmt.Errorf("encode: %w", err)
		}
		n, err = to.Write(encoded)
		if err != nil {
			return fmt.Errorf("write: %w", err)
		}
		if n != len(encoded) {
			return fmt.Errorf("unable to write all %d bytes, could only write %d bytes", len(encoded), n)
		}
	}
	return nil
}

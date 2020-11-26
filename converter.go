package cpconv

import (
	"fmt"
	"io"
)

type Converter struct {
	FromCodepage Codepage
	ToCodepage   Codepage
	BufferSize   int
}

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
		decoded, err := c.FromCodepage.NewDecoder().Bytes(buffer[:n])
		if err != nil {
			return fmt.Errorf("decode: %w", err)
		}
		encoded, err := c.ToCodepage.NewEncoder().Bytes(decoded)
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

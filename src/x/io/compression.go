package io

import (
	"bufio"
	"bytes"
	"fmt"
	"io"

	snappy "github.com/golang/snappy"
)

var snappyStreamID = []byte{0xff, 0x06, 0x00, 0x00, 0x73, 0x4e, 0x61, 0x50, 0x70, 0x59}

// CompressionMethod used for the read/write buffer
type CompressionMethod byte

var validCompressionMethods = []CompressionMethod{
	NoCompression,
	SnappyCompression,
}

const (
	// UnknownCompression mean that we haven't detected used compressiom method quite yet
	UnknownCompression CompressionMethod = iota

	// NoCompression means compression is disabled for the read/write buffer
	NoCompression

	// SnappyCompression enables snappy compression for read/write buffer
	SnappyCompression
)

func (cm CompressionMethod) String() string {
	switch cm {
	case NoCompression:
		return "none"
	case SnappyCompression:
		return "snappy"
	default:
		return ""
	}
}

// UnmarshalYAML unmarshals compression method from YAML configuration
func (cm *CompressionMethod) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var str string
	if err := unmarshal(&str); err != nil {
		return err
	}

	for _, valid := range validCompressionMethods {
		if str == valid.String() {
			*cm = valid
			return nil
		}
	}

	return fmt.Errorf("invalid CompressionMethod '%s' valid types are: %v", str, validCompressionMethods)
}

// SnappyResettableReader is a resettable reader with support for
// snappy compression. Supports fall-back to plain ResettableReader
// if received buffer is not snappy compressed.
type SnappyResettableReader struct {
	reader      *io.Reader
	opts        ResettableReaderOptions
	rr          ResettableReader
	compression CompressionMethod
}

// NewSnappyResettableReader returns a new reader supporting snappy compression.
func NewSnappyResettableReader(r io.Reader, opts ResettableReaderOptions) *SnappyResettableReader {
	return &SnappyResettableReader{&r, opts, nil, UnknownCompression}
}

// Read implements the io.Reader interface
func (s *SnappyResettableReader) Read(p []byte) (n int, err error) {
	if s.compression == UnknownCompression {
		// Try to detect the compression used
		streamHeader := make([]byte, 10)
		streamReader := *s.reader
		if n, err := streamReader.Read(streamHeader); err == nil {
			// Create the reader where reading should continue by adding back the already read header
			newStreamReader := io.MultiReader(bytes.NewReader(streamHeader), streamReader)
			// Check if the stream header matches the snappy stream identifier
			if bytes.Compare(streamHeader, snappyStreamID) == 0 {
				s.compression = SnappyCompression
				s.rr = snappy.NewReader(newStreamReader)
			} else {
				s.compression = NoCompression
				s.rr = bufio.NewReaderSize(newStreamReader, s.opts.ReadBufferSize)
			}
		} else {
			return n, err
		}
	}
	return s.rr.Read(p)
}

// Reset resets the reader state
func (s *SnappyResettableReader) Reset(r io.Reader) {
	s.reader = &r
	s.rr.Reset(r)
	s.compression = UnknownCompression
}

// SnappyResettableWriterFn returns a snappy compression enabled writer
func SnappyResettableWriterFn() ResettableWriterFn {
	return func(r io.Writer, opts ResettableWriterOptions) ResettableWriter {
		return snappy.NewBufferedWriter(r)
	}
}

// SnappyResettableReaderFn returns a snappy compression enabled reader
func SnappyResettableReaderFn() ResettableReaderFn {
	return func(r io.Reader, opts ResettableReaderOptions) ResettableReader {
		return NewSnappyResettableReader(r, opts)
	}
}

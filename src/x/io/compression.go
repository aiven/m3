package io

import (
	"fmt"
	"io"

	snappy "github.com/golang/snappy"
)

// CompressionMethod used for the read/write buffer
type CompressionMethod byte

var validCompressionMethods = []CompressionMethod{
	NoCompression,
	SnappyCompression,
}

const (
	// NoCompression means compression is disabled for the read/write buffer
	NoCompression CompressionMethod = iota

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

// SnappyResettableWriterFn returns a snappy compression enabled writer
func SnappyResettableWriterFn() ResettableWriterFn {
	return func(r io.Writer, opts ResettableWriterOptions) ResettableWriter {
		return snappy.NewBufferedWriter(r)
	}
}

// SnappyResettableReaderFn returns a snappy compression enabled reader
func SnappyResettableReaderFn() ResettableReaderFn {
	return func(r io.Reader, opts ResettableReaderOptions) ResettableReader {
		return snappy.NewReader(r)
	}
}

package io

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"math/rand"
	"testing"

	snappy "github.com/golang/snappy"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWriteBytes(t *testing.T) {
	tests := []struct {
		name        string
		compression CompressionMethod
	}{
		{
			name:        "Uncompressed data",
			compression: NoCompression,
		},
		{
			name:        "Snappy compressed data",
			compression: SnappyCompression,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(tt *testing.T) {
			// Create test data
			testData := make([]byte, 2048)
			rand.Read(testData)

			buf := new(bytes.Buffer)

			// Write uncompressed and compressed data in buffer
			writer := newTestWriter(testCase.compression, buf)
			bs, err := writer.Write(testData)

			require.NoError(t, err)
			require.NoError(t, writer.Flush())
			assert.Equal(t, len(testData), bs)

			r := bufio.NewReader(buf)

			// Read from buffer, supports both uncompressed on compressed data
			reader := NewSnappyResettableReader(r, ResettableReaderOptions{})
			partOne, err := ioutil.ReadAll(reader)

			require.NoError(t, err)
			assert.Equal(t, 2048, len(partOne))

			reader.Reset(r)

			assert.Equal(t, reader.compression, UnknownCompression)
		})
	}
}

func newTestWriter(c CompressionMethod, buf *bytes.Buffer) ResettableWriter {
	switch c {
	case SnappyCompression:
		return snappy.NewWriter(buf)
	default:
		return bufio.NewWriter(buf)
	}
}

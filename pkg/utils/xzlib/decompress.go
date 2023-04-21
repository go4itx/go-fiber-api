package xzlib

import (
	"bytes"
	"compress/zlib"
	"io"
)

// Decompress ...
func Decompress(data []byte) ([]byte, error) {
	b := bytes.NewReader(data)
	r, err := zlib.NewReader(b)
	if err != nil {
		return nil, err
	}

	defer r.Close()
	return io.ReadAll(r)
}

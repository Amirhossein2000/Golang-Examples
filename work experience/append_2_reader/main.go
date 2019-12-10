package main

import (
	"bytes"
	"io"
)

func AddToBody(reader1 io.Reader, reader2 io.Reader, value interface{}) io.Reader {
	buf1 := new(bytes.Buffer)
	buf2 := new(bytes.Buffer)

	_, _ = buf1.ReadFrom(reader1)
	_, _ = buf2.ReadFrom(reader2)

	return bytes.NewReader(append(buf1.Bytes(), buf2.Bytes()...))
}

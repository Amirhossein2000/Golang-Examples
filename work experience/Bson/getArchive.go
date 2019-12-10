package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

const MagicNumber uint32 = 0x8199e26d

type WrappedReadCloser struct {
	io.ReadCloser
	Inner io.ReadCloser
}

func (wrc *WrappedReadCloser) Close() error {
	outerErr := wrc.ReadCloser.Close()
	innerErr := wrc.Inner.Close()
	if outerErr != nil {
		return outerErr
	}
	return innerErr
}

func getArchiveReader(filPath string) (rc io.ReadCloser, err error) {
	rc, err = os.Open(filPath)
	if err != nil {
		return nil, err
	}

	gzrc, err := gzip.NewReader(rc)
	if err != nil {
		return nil, err
	}

	return &WrappedReadCloser{gzrc, rc}, nil

}

func FileValidation(in io.Reader) error {
	readMagicNumberBuf := make([]byte, 4)
	_, err := io.ReadAtLeast(in, readMagicNumberBuf, 4)
	if err != nil {
		return fmt.Errorf("I/O failure reading beginning of archive: %v", err)
	}
	readMagicNumber := uint32(
		(uint32(readMagicNumberBuf[0]) << 0) |
			(uint32(readMagicNumberBuf[1]) << 8) |
			(uint32(readMagicNumberBuf[2]) << 16) |
			(uint32(readMagicNumberBuf[3]) << 24),
	)

	if readMagicNumber != MagicNumber {
		return fmt.Errorf("stream or file does not appear to be a mongodump archive")
	}

	return nil
}

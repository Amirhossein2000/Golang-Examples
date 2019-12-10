package main

import (
	"fmt"
	"io"
)

const MaxBSONSize = 16 * 1024 * 1024
const minBSONSize = 4 + 1
const terminator = -1

type Parser struct {
	In     io.Reader
	buf    [MaxBSONSize]byte
	length int
}

type ParserConsumer interface {
	HeaderBSON([]byte) error
	BodyBSON([]byte) error
	End() error
}

func (parse *Parser) ReadAllBlocks(consumer ParserConsumer) (err error) {
	for err == nil {
		err = parse.ReadBlock(consumer)
	}
	endError := consumer.End()
	if err == io.EOF {
		return endError
	}
	return err
}

func (parse *Parser) ReadBlock(consumer ParserConsumer) (err error) {
	isTerminator, err := parse.readBSONOrTerminator()
	if err != nil {
		return err
	}
	if isTerminator {
		return fmt.Errorf("consecutive terminators / headerless blocks are not allowed.\n")
	}
	err = consumer.HeaderBSON(parse.buf[:parse.length])
	if err != nil {
		return fmt.Errorf("ParserConsumer.HeaderBSON(): %w", err)
	}
	for {
		isTerminator, err = parse.readBSONOrTerminator()
		if err != nil { // all errors, including EOF are errors here
			return fmt.Errorf("ParserConsumer.BodyBSON(): %w", err)
		}
		if isTerminator {
			return nil
		}
		err = consumer.BodyBSON(parse.buf[:parse.length])
		if err != nil {
			return fmt.Errorf("ParserConsumer.BodyBSON(): %w", err)
		}
	}
}

func (parse *Parser) readBSONOrTerminator() (isTerminator bool, err error) {
	parse.length = 0
	_, err = io.ReadFull(parse.In, parse.buf[0:4])
	if err == io.EOF {
		return false, err
	}
	if err != nil {
		return false, fmt.Errorf("I/O error reading length or terminator: %w", err)
	}
	size := int32(
		(uint32(parse.buf[0]) << 0) |
			(uint32(parse.buf[1]) << 8) |
			(uint32(parse.buf[2]) << 16) |
			(uint32(parse.buf[3]) << 24),
	)
	if size == terminator {
		return true, nil
	}
	if size < minBSONSize || size > MaxBSONSize {
		return false, fmt.Errorf("%v is neither a valid bson length nor a archive terminator", size)
	}
	// TODO Because we're reusing this same buffer for all of our IO, we are basically guaranteeing that we'll
	// copy the bytes twice.  At some point we should fix this. It's slightly complex, because we'll need consumer
	// methods closing one buffer and acquiring another
	_, err = io.ReadFull(parse.In, parse.buf[4:size])
	if err != nil {
		// any error, including EOF is an error so we wrap it up
		return false, fmt.Errorf("read bson: %w", err)
	}
	if parse.buf[size-1] != 0x00 {
		return false, fmt.Errorf("bson (size: %v, byte: %d) doesn't end with a null byte", size, parse.buf[size-1])
	}
	parse.length = int(size)
	return false, nil
}

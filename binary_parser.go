package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"math"
)

type BinaryParser struct {
	input *bufio.Reader
}

func NewBinaryParser(reader io.Reader) *BinaryParser {
	return &BinaryParser{
		input: bufio.NewReader(reader),
	}
}

func (p *BinaryParser) readBytes(size int) ([]byte, error) {
	bs := make([]byte, size)
	n, err := io.ReadFull(p.input, bs)
	if n != size {
		return nil, fmt.Errorf("Cannot read %d bytes. got %d bytes", size, n)
	}
	if err != nil {
		return nil, err
	}
	return bs, nil
}

func (p *BinaryParser) readUint8() (uint8, error) {
	return p.input.ReadByte()
}

func (p *BinaryParser) readUint16() (uint16, error) {
	bs, err := p.readBytes(2)
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint16(bs), nil
}

func (p *BinaryParser) readUint32() (uint32, error) {
	bs, err := p.readBytes(4)
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint32(bs), nil
}

func (p *BinaryParser) readUint64() (uint64, error) {
	bs, err := p.readBytes(8)
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint64(bs), nil
}

func (p *BinaryParser) readFloat() (float32, error) {
	bytes, err := p.readUint32()
	if err == nil {
		return math.Float32frombits(bytes), nil
	}
	return 0, err
}

func (p *BinaryParser) readDouble() (float64, error) {
	bytes, err := p.readUint64()
	if err == nil {
		return math.Float64frombits(bytes), nil
	}
	return 0, err
}

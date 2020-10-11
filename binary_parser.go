package parser

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"math"
)

type binaryParser struct {
	input *bufio.Reader
}

type BinaryParser interface {
	ReadBytes(size int) ([]byte, error)
	ReadUint8() (uint8, error)
	ReadUint16() (uint16, error)
	ReadUint32() (uint32, error)
	ReadUint64() (uint64, error)
	ReadFloat() (float32, error)
	ReadDouble() (float64, error)
}

var _ BinaryParser = (*binaryParser)(nil)

func NewBinaryParser(reader io.Reader) BinaryParser {
	return &binaryParser{
		input: bufio.NewReader(reader),
	}
}

func (p *binaryParser) ReadBytes(size int) ([]byte, error) {
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

func (p *binaryParser) ReadUint8() (uint8, error) {
	return p.input.ReadByte()
}

func (p *binaryParser) ReadUint16() (uint16, error) {
	bs, err := p.ReadBytes(2)
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint16(bs), nil
}

func (p *binaryParser) ReadUint32() (uint32, error) {
	bs, err := p.ReadBytes(4)
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint32(bs), nil
}

func (p *binaryParser) ReadUint64() (uint64, error) {
	bs, err := p.ReadBytes(8)
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint64(bs), nil
}

func (p *binaryParser) ReadFloat() (float32, error) {
	bytes, err := p.ReadUint32()
	if err == nil {
		return math.Float32frombits(bytes), nil
	}
	return 0, err
}

func (p *binaryParser) ReadDouble() (float64, error) {
	bytes, err := p.ReadUint64()
	if err == nil {
		return math.Float64frombits(bytes), nil
	}
	return 0, err
}

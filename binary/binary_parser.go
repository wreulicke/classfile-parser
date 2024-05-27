package binary

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"math"
)

type parser struct {
	input *bufio.Reader
}

type Parser interface {
	ReadBytes(size int) ([]byte, error)
	ReadUint8() (uint8, error)
	ReadUint16() (uint16, error)
	ReadUint32() (uint32, error)
	ReadUint64() (uint64, error)
	ReadFloat() (float32, error)
	ReadDouble() (float64, error)
}

var _ Parser = (*parser)(nil)

func NewParser(reader io.Reader) Parser {
	return &parser{
		input: bufio.NewReader(reader),
	}
}

func (p *parser) ReadBytes(size int) ([]byte, error) {
	bs := make([]byte, size)
	n, err := io.ReadFull(p.input, bs)
	if n != size {
		return nil, fmt.Errorf("cannot read %d bytes. got %d bytes", size, n)
	}
	if err != nil {
		return nil, err
	}
	return bs, nil
}

func (p *parser) ReadUint8() (uint8, error) {
	return p.input.ReadByte()
}

func (p *parser) ReadUint16() (uint16, error) {
	bs, err := p.ReadBytes(2)
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint16(bs), nil
}

func (p *parser) ReadUint32() (uint32, error) {
	bs, err := p.ReadBytes(4)
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint32(bs), nil
}

func (p *parser) ReadUint64() (uint64, error) {
	bs, err := p.ReadBytes(8)
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint64(bs), nil
}

func (p *parser) ReadFloat() (float32, error) {
	bytes, err := p.ReadUint32()
	if err == nil {
		return math.Float32frombits(bytes), nil
	}
	return 0, err
}

func (p *parser) ReadDouble() (float64, error) {
	bytes, err := p.ReadUint64()
	if err == nil {
		return math.Float64frombits(bytes), nil
	}
	return 0, err
}

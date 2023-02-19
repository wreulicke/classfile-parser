package parser

import (
	"errors"
	"fmt"
)

type ConstantPool struct {
	Constants []Constant
}

var ErrNotFoundConstant = errors.New("not found constant")

func (c *ConstantPool) LookupUtf8(index uint16) *ConstantUtf8 {
	var i int = int(index) - 1
	if i < 0 {
		return nil
	} else if i > len(c.Constants) {
		return nil
	}
	found := c.Constants[i]
	if utf8, ok := found.(*ConstantUtf8); ok {
		return utf8
	}
	return nil
}

func (c *ConstantPool) GetConstantUtf8(index uint16) (*ConstantUtf8, error) {
	var i int = int(index) - 1
	if i < 0 || i > len(c.Constants) {
		return nil, ErrNotFoundConstant
	}
	clazz, ok := c.Constants[i].(*ConstantUtf8)
	if !ok {
		return nil, fmt.Errorf("Unexpected constant. expected:ConstantUtf8, actual: %T", c.Constants[i])
	}
	return clazz, nil
}

func (c *ConstantPool) GetClassInfo(index uint16) (*ConstantClass, error) {
	var i int = int(index) - 1
	if i < 0 || i > len(c.Constants) {
		return nil, ErrNotFoundConstant
	}
	clazz, ok := c.Constants[i].(*ConstantClass)
	if !ok {
		return nil, fmt.Errorf("Unexpected constant. expected:ConstantClass, actual: %T", c.Constants[i])
	}
	return clazz, nil
}

func (c *ConstantPool) GetClassName(classNameIndex uint16) (string, error) {
	var i int = int(classNameIndex)
	if i < 1 || i > len(c.Constants) {
		return "", ErrNotFoundConstant
	}
	clazz, err := c.GetClassInfo(classNameIndex)
	if err != nil {
		return "", err
	}
	name, err := c.GetConstantUtf8(clazz.NameIndex)
	if err != nil {
		return "", err
	}
	return name.String(), nil
}

type Constant interface {
}

type ConstantClass struct {
	NameIndex uint16
}

type ConstantFieldref struct {
	ClassIndex       uint16
	NameAndTypeIndex uint16
}

type ConstantMethodref struct {
	ClassIndex       uint16
	NameAndTypeIndex uint16
}

type ConstantInterfaceMethodref struct {
	ClassIndex       uint16
	NameAndTypeIndex uint16
}

type ConstantString struct {
	StringIndex uint16
}

type ConstantInteger struct {
	Bytes uint32
}

type ConstantFloat struct {
	Bytes uint32
}

type ConstantLong struct {
	HighBytes uint32
	LowBytes  uint32
}

type ConstantDouble struct {
	HighBytes uint32
	LowBytes  uint32
}

type ConstantNameAndType struct {
	NameIndex       uint16
	DescriptorIndex uint16
}

type ConstantUtf8 struct {
	Length uint16
	Bytes  []byte
}

func (c *ConstantUtf8) String() string {
	return string(c.Bytes)
}

type ConstantMethodHandle struct {
	ReferenceKind  uint8
	ReferenceIndex uint16
}

type ConstantMethodType struct {
	DescriptorIndex uint16
}

type ConstantDynamic struct {
	BootstrapMethodAttrIndex uint16
	NameAndTypeIndex         uint16
}

type ConstantInvokeDynamic struct {
	BootstrapMethodAttrIndex uint16
	NameAndTypeIndex         uint16
}

type ConstantModule struct {
	NameIndex uint16
}

type ConstantPackage struct {
	NameIndex uint16
}

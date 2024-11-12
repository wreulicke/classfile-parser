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
	i := int(index) - 1
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
	i := int(index) - 1
	if i < 0 || i > len(c.Constants) {
		return nil, ErrNotFoundConstant
	}
	clazz, ok := c.Constants[i].(*ConstantUtf8)
	if !ok {
		return nil, fmt.Errorf("unexpected constant. expected:ConstantUtf8, actual: %T", c.Constants[i])
	}
	return clazz, nil
}

func (c *ConstantPool) GetClassInfo(index uint16) (*ConstantClass, error) {
	i := int(index) - 1
	if i < 0 || i > len(c.Constants) {
		return nil, ErrNotFoundConstant
	}
	clazz, ok := c.Constants[i].(*ConstantClass)
	if !ok {
		return nil, fmt.Errorf("unexpected constant. expected:ConstantClass, actual: %T", c.Constants[i])
	}
	return clazz, nil
}

func (c *ConstantPool) GetClassName(classNameIndex uint16) (string, error) {
	i := int(classNameIndex)
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
	Name() string
}

type ConstantClass struct {
	NameIndex uint16
}

func (c *ConstantClass) Name() string {
	return "Class"
}

type ConstantFieldref struct {
	ClassIndex       uint16
	NameAndTypeIndex uint16
}

func (c *ConstantFieldref) Name() string {
	return "Fieldref"
}

type ConstantMethodref struct {
	ClassIndex       uint16
	NameAndTypeIndex uint16
}

func (c *ConstantMethodref) Name() string {
	return "Methodref"
}

type ConstantInterfaceMethodref struct {
	ClassIndex       uint16
	NameAndTypeIndex uint16
}

func (c *ConstantInterfaceMethodref) Name() string {
	return "InterfaceMethodref"
}

type ConstantString struct {
	StringIndex uint16
}

func (c *ConstantString) Name() string {
	return "String"
}

type ConstantInteger struct {
	Bytes uint32
}

func (c *ConstantInteger) Name() string {
	return "Integer"
}

type ConstantFloat struct {
	Bytes uint32
}

func (c *ConstantFloat) Name() string {
	return "Float"
}

type ConstantLong struct {
	HighBytes uint32
	LowBytes  uint32
}

func (c *ConstantLong) Name() string {
	return "Long"
}

type ConstantDouble struct {
	HighBytes uint32
	LowBytes  uint32
}

func (c *ConstantDouble) Name() string {
	return "Double"
}

type ConstantNameAndType struct {
	NameIndex       uint16
	DescriptorIndex uint16
}

func (c *ConstantNameAndType) Name() string {
	return "NameAndType"
}

type ConstantUtf8 struct {
	Length uint16
	Bytes  []byte
}

func (c *ConstantUtf8) Name() string {
	return "Utf8"
}

func (c *ConstantUtf8) String() string {
	return string(c.Bytes)
}

type ConstantMethodHandle struct {
	ReferenceKind  uint8
	ReferenceIndex uint16
}

func (c *ConstantMethodHandle) Name() string {
	return "MethodHandle"
}

type ConstantMethodType struct {
	DescriptorIndex uint16
}

func (c *ConstantMethodType) Name() string {
	return "MethodType"
}

type ConstantDynamic struct {
	BootstrapMethodAttrIndex uint16
	NameAndTypeIndex         uint16
}

func (c *ConstantDynamic) Name() string {
	return "Dynamic"
}

type ConstantInvokeDynamic struct {
	BootstrapMethodAttrIndex uint16
	NameAndTypeIndex         uint16
}

func (c *ConstantInvokeDynamic) Name() string {
	return "InvokeDynamic"
}

type ConstantModule struct {
	NameIndex uint16
}

func (c *ConstantModule) Name() string {
	return "Module"
}

type ConstantPackage struct {
	NameIndex uint16
}

func (c *ConstantPackage) Name() string {
	return "Package"
}

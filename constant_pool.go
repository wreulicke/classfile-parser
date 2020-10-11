package parser

type ConstantPool struct {
	Constants []Constant
}

func (c *ConstantPool) LookupUtf8(index uint16) *ConstantUtf8 {
	i := index - 1
	if i < 0 {
		return nil
	} else if i > uint16(len(c.Constants)) {
		return nil
	}
	found := c.Constants[i]
	if utf8, ok := found.(*ConstantUtf8); ok {
		return utf8
	}
	return nil
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

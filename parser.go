package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
)

type Parser struct {
	*BinaryParser
	error error
}

func New(input io.Reader) *Parser {
	l := &Parser{BinaryParser: NewBinaryParser(input)}
	return l
}

type Classfile struct {
	MajorVersion uint16
	MinorVersion uint16
	ConstantPool *ConstantPool
	AccessFlags  uint16
	ThisClass    uint16
	SuperClass   uint16
	Interfaces   []uint16
	Fields       []*Field
	Methods      []*Method
	Attribute    []Attribute
}

func (p *Parser) Parse() (*Classfile, error) {
	err := p.readCaffbabe()
	if err != nil {
		return nil, err
	}
	c := &Classfile{}
	if err := p.readMinorVersion(c); err != nil {
		return nil, err
	}
	if err := p.readMajorVersion(c); err != nil {
		return nil, err
	}
	if err := p.readConstantPool(c); err != nil {
		return nil, err
	}
	if err := p.readAccessFlag(c); err != nil {
		return nil, err
	}
	if err := p.readThisClass(c); err != nil {
		return nil, err
	}
	if err := p.readSuperClass(c); err != nil {
		return nil, err
	}
	if err := p.readInterfaces(c); err != nil {
		return nil, err
	}
	if err := p.readFields(c); err != nil {
		return nil, err
	}
	if err := p.readMethods(c); err != nil {
		return nil, err
	}
	if c.Attribute, err = p.readAttributes(c.ConstantPool); err != nil {
		return nil, err
	}
	return c, nil
}

func (p *Parser) readCaffbabe() error {
	bs, err := p.readBytes(4)
	if err != nil {
		return err
	}
	if !bytes.Equal(bs, []byte{0xCA, 0xFE, 0xBA, 0xBE}) {
		return errors.New("magic is wrong")
	}
	return nil
}

func (p *Parser) readMinorVersion(c *Classfile) error {
	v, err := p.readUint16()
	if err != nil {
		return err
	}
	c.MinorVersion = v
	return nil
}

func (p *Parser) readMajorVersion(c *Classfile) error {
	v, err := p.readUint16()
	if err != nil {
		return err
	}
	c.MajorVersion = v
	return nil
}

func (p *Parser) readConstantPool(c *Classfile) error {
	count, err := p.readUint16()
	if err != nil {
		return err
	}
	var i uint16
	cp := &ConstantPool{Constants: make([]Constant, count-1)}
	c.ConstantPool = cp
	for ; i < count-1; i++ {
		tag, err := p.readUint8()
		if err != nil {
			return nil
		}
		switch tag {
		case 7:
			c := &ConstantClass{}
			cp.Constants[i] = c
			c.NameIndex, err = p.readUint16()
			if err != nil {
				return err
			}
		case 9:
			c := &ConstantFieldref{}
			cp.Constants[i] = c
			c.ClassIndex, err = p.readUint16()
			if err != nil {
				return err
			}
			c.NameAndTypeIndex, err = p.readUint16()
			if err != nil {
				return err
			}
		case 10:
			c := &ConstantMethodref{}
			cp.Constants[i] = c
			c.ClassIndex, err = p.readUint16()
			if err != nil {
				return err
			}
			c.NameAndTypeIndex, err = p.readUint16()
			if err != nil {
				return err
			}
		case 11:
			c := &ConstantInterfaceMethodref{}
			cp.Constants[i] = c
			c.ClassIndex, err = p.readUint16()
			if err != nil {
				return err
			}
			c.NameAndTypeIndex, err = p.readUint16()
			if err != nil {
				return err
			}
		case 8:
			c := &ConstantString{}
			cp.Constants[i] = c
			c.StringIndex, err = p.readUint16()
			if err != nil {
				return err
			}
		case 3:
			c := &ConstantInteger{}
			cp.Constants[i] = c
			c.Bytes, err = p.readUint32()
			if err != nil {
				return err
			}
		case 4:
			c := &ConstantFloat{}
			cp.Constants[i] = c
			c.Bytes, err = p.readUint32()
			if err != nil {
				return err
			}
		case 5:
			c := &ConstantLong{}
			cp.Constants[i] = c
			i++
			c.HighBytes, err = p.readUint32()
			if err != nil {
				return err
			}
			c.LowBytes, err = p.readUint32()
			if err != nil {
				return err
			}
		case 6:
			c := &ConstantDouble{}
			cp.Constants[i] = c
			i++
			c.HighBytes, err = p.readUint32()
			if err != nil {
				return err
			}
			c.LowBytes, err = p.readUint32()
			if err != nil {
				return err
			}
		case 12:
			c := &ConstantNameAndType{}
			cp.Constants[i] = c
			c.NameIndex, err = p.readUint16()
			if err != nil {
				return err
			}
			c.DescriptorIndex, err = p.readUint16()
			if err != nil {
				return err
			}
		case 1:
			c := &ConstantUtf8{}
			cp.Constants[i] = c
			count, err := p.readUint16()
			if err != nil {
				return err
			}

			c.Bytes, err = p.readBytes(int(count))
			if err != nil {
				return err
			}
		case 15:
			c := &ConstantMethodHandle{}
			cp.Constants[i] = c
			c.ReferenceKind, err = p.readUint8()
			if err != nil {
				return err
			}
			c.ReferenceIndex, err = p.readUint16()
			if err != nil {
				return err
			}
		case 16:
			c := &ConstantMethodType{}
			cp.Constants[i] = c
			c.DescriptorIndex, err = p.readUint16()
			if err != nil {
				return err
			}
		case 17:
			c := &ConstantDynamic{}
			cp.Constants[i] = c
			c.BootstrapMethodAttrIndex, err = p.readUint16()
			if err != nil {
				return err
			}
			c.NameAndTypeIndex, err = p.readUint16()
			if err != nil {
				return err
			}
		case 18:
			c := &ConstantInvokeDynamic{}
			cp.Constants[i] = c
			c.BootstrapMethodAttrIndex, err = p.readUint16()
			if err != nil {
				return err
			}
			c.NameAndTypeIndex, err = p.readUint16()
			if err != nil {
				return err
			}
		case 19:
			c := &ConstantModule{}
			cp.Constants[i] = c
			c.NameIndex, err = p.readUint16()
			if err != nil {
				return err
			}
		case 20:
			c := &ConstantPackage{}
			cp.Constants[i] = c
			c.NameIndex, err = p.readUint16()
			if err != nil {
				return err
			}
		default:
			return fmt.Errorf("Unsupported tags for constant pool. tag:%d", tag)
		}

	}
	return nil
}

func (p *Parser) readAccessFlag(c *Classfile) (err error) {
	c.AccessFlags, err = p.readUint16()
	return
}

func (p *Parser) readThisClass(c *Classfile) (err error) {
	c.ThisClass, err = p.readUint16()
	return
}

func (p *Parser) readSuperClass(c *Classfile) (err error) {
	c.SuperClass, err = p.readUint16()
	return
}

func (p *Parser) readInterfaces(c *Classfile) error {
	count, err := p.readUint16()
	if err != nil {
		return err
	}
	var i uint16
	for ; i < count; i++ {
		interfaceIndex, err := p.readUint16()
		if err != nil {
			return err
		}
		c.Interfaces = append(c.Interfaces, interfaceIndex)
	}
	return nil
}

func (p *Parser) readFields(c *Classfile) error {
	count, err := p.readUint16()
	if err != nil {
		return err
	}
	var i uint16
	for ; i < count; i++ {
		f := &Field{}
		c.Fields = append(c.Fields, f)
		f.AccessFlags, err = p.readUint16()
		if err != nil {
			return err
		}
		f.NameIndex, err = p.readUint16()
		if err != nil {
			return err
		}
		f.DescriptorIndex, err = p.readUint16()
		if err != nil {
			return err
		}
		f.Attributes, err = p.readAttributes(c.ConstantPool)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *Parser) readMethods(c *Classfile) error {
	count, err := p.readUint16()
	if err != nil {
		return err
	}
	var i uint16
	for ; i < count; i++ {
		m := &Method{}
		c.Methods = append(c.Methods, m)
		m.AccessFlags, err = p.readUint16()
		if err != nil {
			return err
		}
		m.NameIndex, err = p.readUint16()
		if err != nil {
			return err
		}
		m.DescriptorIndex, err = p.readUint16()
		if err != nil {
			return err
		}
		m.Attributes, err = p.readAttributes(c.ConstantPool)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *Parser) readAttributes(c *ConstantPool) ([]Attribute, error) {
	count, err := p.readUint16()
	if err != nil {
		return nil, err
	}
	as := make([]Attribute, 0, count)
	var i uint16
	for ; i < count; i++ {
		a, err := p.readAttribute(c)
		if err != nil {
			return nil, err
		}
		as = append(as, a)
	}
	return as, nil
}

func (p *Parser) readAttribute(constantPool *ConstantPool) (Attribute, error) {
	attributeNameIndex, err := p.readUint16()
	if err != nil {
		return nil, err
	}
	attributeLength, err := p.readUint32()
	if err != nil {
		return nil, err
	}
	u := constantPool.LookupUtf8(attributeNameIndex)
	if u == nil {
		return nil, fmt.Errorf("attribute name index is invalid: index:%d", attributeNameIndex)
	}
	bs, err := p.readBytes(int(attributeLength))
	if err != nil {
		return nil, err
	}
	parser := NewBinaryParser(bytes.NewBuffer(bs))
	return readAttribute(parser, attributeLength, u.String())
}

func readAttribute(p *BinaryParser, attributeLength uint32, attributeName string) (Attribute, error) {
	var err error
	switch attributeName {
	case "ConstantValue":
		a := &AttributeConstantValue{}
		a.ConstantValueIndex, err = p.readUint16()
		return a, err
	case "Code":
		goto notImplemented
	case "StackMapTable":
		goto notImplemented
	case "Exceptions":
		exceptionCount, err := p.readUint16()
		if err != nil {
			return nil, err
		}
		a := &AttributeExceptions{ExceptionIndexes: make([]uint16, 0, exceptionCount)}
		var i uint16
		for ; i < exceptionCount; i++ {
			exceptionIndex, err := p.readUint16()
			if err != nil {
				return nil, err
			}
			a.ExceptionIndexes = append(a.ExceptionIndexes, exceptionIndex)
		}
	case "InnerClasses":
		numberOfClasses, err := p.readUint16()
		if err != nil {
			return nil, err
		}
		a := &AttributeInnerClasses{InnerClasses: make([]*InnerClass, 0, numberOfClasses)}
		var i uint16
		for ; i < numberOfClasses; i++ {
			c := &InnerClass{}
			c.InnerClassInfoIndex, err = p.readUint16()
			if err != nil {
				return nil, err
			}
			c.OuterClassInfoIndex, err = p.readUint16()
			if err != nil {
				return nil, err
			}
			c.InnerNameIndex, err = p.readUint16()
			if err != nil {
				return nil, err
			}
			c.InnerClassAccessFlags, err = p.readUint16()
			if err != nil {
				return nil, err
			}
			a.InnerClasses = append(a.InnerClasses, c)
		}
		return a, nil
	case "EnclosingMethod":
		a := &AttributeEnclosingMethod{}
		a.ClassIndex, err = p.readUint16()
		if err != nil {
			return nil, err
		}
		a.MethodIndex, err = p.readUint16()
		return a, err
	case "Synthetic":
		a := &AttributeSynthetic{}
		return a, nil
	case "Signature":
		a := &AttributeSignature{}
		a.Signature, err = p.readUint16()
		return a, err
	case "SourceFile":
		a := &AttributeSourceFile{}
		a.SourcefileIndex, err = p.readUint16()
		return a, err
	case "SourceDebugExtension":
		a := &AttributeSourceDebugExtension{}
		a.DebugExtension, err = p.readBytes(int(attributeLength))
		return a, err
	case "LineNumberTable":
		lineNumberTableLength, err := p.readUint16()
		if err != nil {
			return nil, err
		}
		a := &AttributeLineNumberTable{LineNumberTable: make([]*LineNumber, 0, lineNumberTableLength)}
		var i uint16
		for ; i < lineNumberTableLength; i++ {
			ln := &LineNumber{}
			ln.StartPc, err = p.readUint16()
			if err != nil {
				return nil, err
			}
			ln.LineNumber, err = p.readUint16()
			if err != nil {
				return nil, err
			}
			a.LineNumberTable = append(a.LineNumberTable, ln)
		}
		return a, nil
	case "LocalVariableTable":
		goto notImplemented
	case "LocalVariableTypeTable":
		goto notImplemented
	case "Deprecated":
		if attributeLength != 0 {
			return nil, errors.New("Deprecated attribute length should be 2")
		}
		return &AttributeDeprecated{}, nil
	case "RuntimeVisibleAnnotations":
		goto notImplemented
	case "RuntimeInvisibleAnnotations":
		goto notImplemented
	case "RuntimeVisibleParameterAnnotations":
		goto notImplemented
	case "RuntimeInvisibleParameterAnnotations":
		goto notImplemented
	case "RuntimeVisibleTypeAnnotations":
		goto notImplemented
	case "RuntimeInvisibleTypeAnnotations":
		goto notImplemented
	case "AnnotationDefault":
		goto notImplemented
	case "BoostrapMethods":
		goto notImplemented
	case "MethodParameters":
		goto notImplemented
	case "Module":
		goto notImplemented
	case "ModulePackage":
		packageCount, err := p.readUint16()
		if err != nil {
			return nil, err
		}
		a := &AttributeModulePackage{PackageIndexes: make([]uint16, 0, packageCount)}
		var i uint16
		for ; i < packageCount; i++ {
			packageIndex, err := p.readUint16()
			if err != nil {
				return nil, err
			}
			a.PackageIndexes = append(a.PackageIndexes, packageIndex)
		}
		return a, nil
	case "ModuleMainClass":
		a := &AttributeModuleMainClass{}
		a.MainClassIndex, err = p.readUint16()
		return a, err
	case "NestHost":
		a := &AttributeNestHost{}
		a.HostClassIndex, err = p.readUint16()
		return a, err
	case "NestMembers":
		numberOfClasses, err := p.readUint16()
		if err != nil {
			return nil, err
		}
		a := &AttributeNestMembers{Classes: make([]uint16, 0, numberOfClasses)}
		var i uint16
		for ; i < numberOfClasses; i++ {
			class, err := p.readUint16()
			if err != nil {
				return nil, err
			}
			a.Classes = append(a.Classes, class)
		}
		return a, nil
	}
notImplemented:
	return nil, nil
}

func (p *Parser) readBytes(size int) ([]byte, error) {
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

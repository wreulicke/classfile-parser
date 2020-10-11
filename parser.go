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
		localVaribleTableLength, err := p.readUint16()
		if err != nil {
			return nil, err
		}
		a := &AttributeLocalVaribleTable{}
		var i uint16
		for ; i < localVaribleTableLength; i++ {
			ln := &LocalVarible{}
			ln.StartPc, err = p.readUint16()
			if err != nil {
				return nil, err
			}
			ln.Length, err = p.readUint16()
			if err != nil {
				return nil, err
			}
			ln.NameIndex, err = p.readUint16()
			if err != nil {
				return nil, err
			}
			ln.DescriptorInedx, err = p.readUint16()
			if err != nil {
				return nil, err
			}
			ln.Index, err = p.readUint16()
			if err != nil {
				return nil, err
			}
			a.LocalVaribleTable = append(a.LocalVaribleTable, ln)
		}
		return a, nil
	case "LocalVariableTypeTable":
		localVaribleTypeLength, err := p.readUint16()
		if err != nil {
			return nil, err
		}
		a := &AttributeLocalVaribleTypeTable{}
		var i uint16
		for ; i < localVaribleTypeLength; i++ {
			ln := &LocalVaribleType{}
			ln.StartPc, err = p.readUint16()
			if err != nil {
				return nil, err
			}
			ln.Length, err = p.readUint16()
			if err != nil {
				return nil, err
			}
			ln.NameIndex, err = p.readUint16()
			if err != nil {
				return nil, err
			}
			ln.SignatureInedx, err = p.readUint16()
			if err != nil {
				return nil, err
			}
			ln.Index, err = p.readUint16()
			if err != nil {
				return nil, err
			}
			a.LocalVaribleTypeTable = append(a.LocalVaribleTypeTable, ln)
		}
		return a, nil
	case "Deprecated":
		if attributeLength != 0 {
			return nil, errors.New("Deprecated attribute length should be 2")
		}
		return &AttributeDeprecated{}, nil
	case "RuntimeVisibleAnnotations":
		numAnnotations, err := p.readUint16()
		if err != nil {
			return nil, err
		}
		a := &AttributeRuntimeVisibleAnnotations{}
		var i uint16
		for ; i < numAnnotations; i++ {
			annot, err := readAnnotation(p)
			if err != nil {
				return nil, err
			}
			a.Annotations = append(a.Annotations, annot)
		}
		return a, nil
	case "RuntimeInvisibleAnnotations":
		numAnnotations, err := p.readUint16()
		if err != nil {
			return nil, err
		}
		a := &AttributeRuntimeInvisibleAnnotations{}
		var i uint16
		for ; i < numAnnotations; i++ {
			annot, err := readAnnotation(p)
			if err != nil {
				return nil, err
			}
			a.Annotations = append(a.Annotations, annot)
		}
		return a, nil
	case "RuntimeVisibleParameterAnnotations":
		numAnnotations, err := p.readUint16()
		if err != nil {
			return nil, err
		}
		a := &AttributeRuntimeVisibleParameterAnnotations{}
		var i uint16
		for ; i < numAnnotations; i++ {
			annot, err := readParameterAnnotation(p)
			if err != nil {
				return nil, err
			}
			a.ParameterAnnotations = append(a.ParameterAnnotations, annot)
		}
		return a, nil
	case "RuntimeInvisibleParameterAnnotations":
		numAnnotations, err := p.readUint16()
		if err != nil {
			return nil, err
		}
		a := &AttributeRuntimeVisibleParameterAnnotations{}
		var i uint16
		for ; i < numAnnotations; i++ {
			annot, err := readParameterAnnotation(p)
			if err != nil {
				return nil, err
			}
			a.ParameterAnnotations = append(a.ParameterAnnotations, annot)
		}
		return a, nil
	case "RuntimeVisibleTypeAnnotations":
		numAnnotations, err := p.readUint16()
		if err != nil {
			return nil, err
		}
		a := &AttributeRuntimeVisibleTypeAnnotations{}
		var i uint16
		for ; i < numAnnotations; i++ {
			annot, err := readTypeAnnotation(p)
			if err != nil {
				return nil, err
			}
			a.TypeAnnotations = append(a.TypeAnnotations, annot)
		}
		return a, nil
	case "RuntimeInvisibleTypeAnnotations":
		numAnnotations, err := p.readUint16()
		if err != nil {
			return nil, err
		}
		a := &AttributeRuntimeInvisibleTypeAnnotations{}
		var i uint16
		for ; i < numAnnotations; i++ {
			annot, err := readTypeAnnotation(p)
			if err != nil {
				return nil, err
			}
			a.TypeAnnotations = append(a.TypeAnnotations, annot)
		}
		return a, nil
	case "AnnotationDefault":
		value, err := readElementValue(p)
		if err != nil {
			return nil, err
		}
		return &AttributeAnnotationDefault{
			DefaultValue: value,
		}, nil
	case "BoostrapMethods":
		numBootstrapMethods, err := p.readUint16()
		if err != nil {
			return nil, err
		}
		attribute := &AttributeBootstrapMethods{}
		var i uint16
		for ; i < numBootstrapMethods; i++ {
			b := &BootstrapMethod{}
			b.BootstrapMethodRef, err = p.readUint16()
			if err != nil {
				return nil, err
			}
			numBootstrapArguments, err := p.readUint16()
			if err != nil {
				return nil, err
			}
			var j uint16
			for ; j < numBootstrapArguments; j++ {
				a, err := p.readUint16()
				if err != nil {
					return nil, err
				}
				b.BootstrapArguments = append(b.BootstrapArguments, a)
			}
			attribute.BootstrapMethods = append(attribute.BootstrapMethods, b)
		}
		return attribute, nil
	case "MethodParameters":
		parametersCount, err := p.readUint8()
		if err != nil {
			return nil, err
		}
		a := &AttributeMethodParameters{}
		var i uint8
		for ; i < parametersCount; i++ {
			param := &MethodParameter{}
			param.NameIndex, err = p.readUint16()
			if err != nil {
				return nil, err
			}
			param.AccessFlags, err = p.readUint16()
			if err != nil {
				return nil, err
			}
			a.Parameters = append(a.Parameters, param)
		}
		return a, nil
	case "Module":
		a := &AttributeModule{}
		a.ModuleNameIndex, err = p.readUint16()
		if err != nil {
			return nil, err
		}
		a.ModuleFlags, err = p.readUint16()
		if err != nil {
			return nil, err
		}
		a.ModuleVersionIndex, err = p.readUint16()
		if err != nil {
			return nil, err
		}
		requiresCount, err := p.readUint16()
		if err != nil {
			return nil, err
		}

		var i uint16
		for ; i < requiresCount; i++ {
			r := &Require{}
			r.RequiresIndex, err = p.readUint16()
			if err != nil {
				return nil, err
			}
			r.RequiresFlags, err = p.readUint16()
			if err != nil {
				return nil, err
			}
			r.RequiresVersionIndex, err = p.readUint16()
			if err != nil {
				return nil, err
			}
			a.Requires = append(a.Requires, r)
		}

		exportsCount, err := p.readUint16()
		if err != nil {
			return nil, err
		}
		for i = 0; i < exportsCount; i++ {
			e := &Export{}
			e.ExportsIndex, err = p.readUint16()
			if err != nil {
				return nil, err
			}
			e.ExportsFlags, err = p.readUint16()
			if err != nil {
				return nil, err
			}
			exportsToCount, err := p.readUint16()
			if err != nil {
				return nil, err
			}
			var j uint16
			for ; j < exportsToCount; j++ {
				t, err := p.readUint16()
				if err != nil {
					return nil, err
				}
				e.ExportsTo = append(e.ExportsTo, t)
			}
		}

		opensCount, err := p.readUint16()
		if err != nil {
			return nil, err
		}
		for i = 0; i < opensCount; i++ {
			e := &Open{}
			e.OpensIndex, err = p.readUint16()
			if err != nil {
				return nil, err
			}
			e.OpensFlags, err = p.readUint16()
			if err != nil {
				return nil, err
			}
			opensToCount, err := p.readUint16()
			if err != nil {
				return nil, err
			}
			var j uint16
			for ; j < opensToCount; j++ {
				t, err := p.readUint16()
				if err != nil {
					return nil, err
				}
				e.OpensTo = append(e.OpensTo, t)
			}
		}

		usesCount, err := p.readUint16()
		if err != nil {
			return nil, err
		}
		for i = 0; i < usesCount; i++ {
			use, err := p.readUint16()
			if err != nil {
				return nil, err
			}
			a.Uses = append(a.Uses, use)
		}

		providesCount, err := p.readUint16()
		if err != nil {
			return nil, err
		}
		for i = 0; i < providesCount; i++ {
			e := &Provide{}
			e.ProvidesIndex, err = p.readUint16()
			if err != nil {
				return nil, err
			}
			providesWithCount, err := p.readUint16()
			if err != nil {
				return nil, err
			}
			var j uint16
			for ; j < providesWithCount; j++ {
				t, err := p.readUint16()
				if err != nil {
					return nil, err
				}
				e.ProvidesWith = append(e.ProvidesWith, t)
			}
		}

		return a, nil
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

func readAnnotation(parser *BinaryParser) (*Annotation, error) {
	a := &Annotation{}
	var err error
	a.TypeIndex, err = parser.readUint16()
	if err != nil {
		return nil, err
	}
	a.ElementValuePairs, err = readElementValuePairs(parser)
	return a, err
}

func readElementValuePairs(parser *BinaryParser) ([]*ElementValuePair, error) {
	numElementValuePair, err := parser.readUint16()
	if err != nil {
		return nil, err
	}
	pairs := make([]*ElementValuePair, 0, numElementValuePair)
	var i uint16
	for ; i < numElementValuePair; i++ {
		p, err := readElementValuePair(parser)
		if err != nil {
			return nil, err
		}
		pairs = append(pairs, p)
	}
	return pairs, nil
}

func readElementValuePair(parser *BinaryParser) (*ElementValuePair, error) {
	p := &ElementValuePair{}
	var err error
	p.ElementNameIndex, err = parser.readUint16()
	return p, err
}

func readParameterAnnotation(parser *BinaryParser) (*ParameterAnnotation, error) {
	a := &ParameterAnnotation{}
	numAnnotations, err := parser.readUint16()
	if err != nil {
		return nil, err
	}
	var i uint16
	for ; i < numAnnotations; i++ {
		annot, err := readAnnotation(parser)
		if err != nil {
			return nil, err
		}
		a.Annotations = append(a.Annotations, annot)
	}
	return a, nil
}

func readTypeAnnotation(parser *BinaryParser) (*TypeAnnotation, error) {
	a := &TypeAnnotation{}
	targetType, err := parser.readUint8()
	if err != nil {
		return nil, err
	}
	switch targetType {
	case 0x00:
		a.TargetInfo, err = readTypeParameterTarget(parser)
	case 0x01:
		a.TargetInfo, err = readTypeParameterTarget(parser)
	case 0x10:
		a.TargetInfo, err = readSuperTypeTarget(parser)
	case 0x11:
		a.TargetInfo, err = readTypeParameterBoundTarget(parser)
	case 0x12:
		a.TargetInfo, err = readTypeParameterBoundTarget(parser)
	case 0x13:
		a.TargetInfo = &EmptyTarget{}
	case 0x14:
		a.TargetInfo = &EmptyTarget{}
	case 0x15:
		a.TargetInfo = &EmptyTarget{}
	case 0x16:
		a.TargetInfo, err = readFormalParameterTarget(parser)
	case 0x17:
		a.TargetInfo, err = readThrowsTarget(parser)
	case 0x40:
		a.TargetInfo, err = readLocalVarTarget(parser)
	case 0x41:
		a.TargetInfo, err = readLocalVarTarget(parser)
	case 0x42:
		a.TargetInfo, err = readCatchTarget(parser)
	case 0x43:
		a.TargetInfo, err = readOffsetTarget(parser)
	case 0x44:
		a.TargetInfo, err = readOffsetTarget(parser)
	case 0x45:
		a.TargetInfo, err = readOffsetTarget(parser)
	case 0x46:
		a.TargetInfo, err = readOffsetTarget(parser)
	case 0x47:
		a.TargetInfo, err = readTypeArgumentTarget(parser)
	case 0x48:
		a.TargetInfo, err = readTypeArgumentTarget(parser)
	case 0x49:
		a.TargetInfo, err = readTypeArgumentTarget(parser)
	case 0x4A:
		a.TargetInfo, err = readTypeArgumentTarget(parser)
	case 0x4B:
		a.TargetInfo, err = readTypeArgumentTarget(parser)
	default:
		return nil, fmt.Errorf("Unsupported target type for TypeAnnotation. tag: %d", targetType)
	}
	a.TargetPath, err = readTypePath(parser)
	if err != nil {
		return nil, err
	}
	a.TypeIndex, err = parser.readUint16()
	if err != nil {
		return nil, err
	}
	a.ElementValuePairs, err = readElementValuePairs(parser)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func readTypeParameterTarget(parser *BinaryParser) (TargetInfo, error) {
	target := &TypeParameterTarget{}
	var err error
	target.TypeParameterIndex, err = parser.readUint8()
	return target, err
}

func readSuperTypeTarget(parser *BinaryParser) (TargetInfo, error) {
	target := &SuperTypeTarget{}
	var err error
	target.SuperTypeIndex, err = parser.readUint16()
	return target, err
}

func readTypeParameterBoundTarget(parser *BinaryParser) (TargetInfo, error) {
	target := &TypeParameterBoundTarget{}
	var err error
	target.TypeParameterIndex, err = parser.readUint8()
	if err != nil {
		return nil, err
	}
	target.BoundIndex, err = parser.readUint8()
	return target, err
}

func readFormalParameterTarget(parser *BinaryParser) (TargetInfo, error) {
	target := &FormalParameterTarget{}
	var err error
	target.FormalParameterIndex, err = parser.readUint8()
	return target, err
}

func readThrowsTarget(parser *BinaryParser) (TargetInfo, error) {
	target := &ThrowsTarget{}
	var err error
	target.ThrowsTypeIndex, err = parser.readUint16()
	return target, err
}

func readLocalVarTarget(parser *BinaryParser) (TargetInfo, error) {
	target := &LocalVarTarget{}
	tableLength, err := parser.readUint16()
	if err != nil {
		return nil, err
	}
	var i uint16
	for ; i < tableLength; i++ {
		table, err := readLocalVarTargetTable(parser)
		if err != nil {
			return nil, err
		}
		target.LocalVarTargetTables = append(target.LocalVarTargetTables, table)
	}
	return target, err
}

func readLocalVarTargetTable(parser *BinaryParser) (*LocalVarTargetTable, error) {
	t := &LocalVarTargetTable{}
	var err error
	t.StartPc, err = parser.readUint16()
	if err != nil {
		return nil, err
	}
	t.Length, err = parser.readUint16()
	if err != nil {
		return nil, err
	}
	t.Index, err = parser.readUint16()
	return t, err
}

func readCatchTarget(parser *BinaryParser) (TargetInfo, error) {
	target := &CatchTarget{}
	var err error
	target.ExceptionTableIndex, err = parser.readUint16()
	return target, err
}

func readOffsetTarget(parser *BinaryParser) (TargetInfo, error) {
	target := &OffsetTarget{}
	var err error
	target.Offset, err = parser.readUint16()
	return target, err
}

func readTypeArgumentTarget(parser *BinaryParser) (TargetInfo, error) {
	target := &TypeArgumentTarget{}
	var err error
	target.Offset, err = parser.readUint16()
	if err != nil {
		return nil, err
	}
	target.TypeArgumentIndex, err = parser.readUint8()
	return target, err
}

func readTypePath(parser *BinaryParser) (*TypePath, error) {
	pathLength, err := parser.readUint8()
	if err != nil {
		return nil, err
	}
	typePath := &TypePath{}
	var i uint8
	for ; i < pathLength; i++ {
		path := &Path{}
		path.TypePathKind, err = parser.readUint8()
		if err != nil {
			return nil, err
		}
		path.TypeArgumentIndex, err = parser.readUint8()
		if err != nil {
			return nil, err
		}
		typePath.Paths = append(typePath.Paths, path)
	}
	return typePath, nil
}

func readElementValue(parser *BinaryParser) (*ElementValue, error) {
	tag, err := parser.readUint8()
	if err != nil {
		return nil, err
	}
	ev := &ElementValue{}
	switch tag {
	case 'B':
		ev.ConstValue, err = readElementValueConst(parser)
	case 'C':
		ev.ConstValue, err = readElementValueConst(parser)
	case 'F':
		ev.ConstValue, err = readElementValueConst(parser)
	case 'I':
		ev.ConstValue, err = readElementValueConst(parser)
	case 'J':
		ev.ConstValue, err = readElementValueConst(parser)
	case 'S':
		ev.ConstValue, err = readElementValueConst(parser)
	case 'Z':
		ev.ConstValue, err = readElementValueConst(parser)
	case 's':
		ev.ConstValue, err = readElementValueConst(parser)
	case 'e':
		ev.EnumConstValue, err = readElementValueEnumConst(parser)
	case 'c':
		ev.ClassInfo, err = readElementClassInfo(parser)
	case '@':
		ev.AnnotationValue, err = readAnnotation(parser)
	case '[':
		ev.ArrayValue, err = readElementArrayValue(parser)
	default:
		err = errors.New("Unsupported tag for element value")
	}
	return ev, err
}

func readElementValueConst(parser *BinaryParser) (*ElementValueConstValue, error) {
	e := &ElementValueConstValue{}
	var err error
	e.ConstValueIndex, err = parser.readUint16()
	return e, err
}

func readElementValueEnumConst(parser *BinaryParser) (*ElementValueEnumConstValue, error) {
	e := &ElementValueEnumConstValue{}
	var err error
	e.TypeNameIndex, err = parser.readUint16()
	if err != nil {
		return nil, err
	}
	e.ConstNameIndex, err = parser.readUint16()
	return e, err
}

func readElementClassInfo(parser *BinaryParser) (*ElementValueClassInfo, error) {
	e := &ElementValueClassInfo{}
	var err error
	e.ClassInfoIndex, err = parser.readUint16()
	return e, err
}

func readElementArrayValue(parser *BinaryParser) (*ElementValueArrayValue, error) {
	e := &ElementValueArrayValue{}
	numValues, err := parser.readUint16()
	var i uint16
	for ; i < numValues; i++ {
		v, err := readElementValue(parser)
		if err != nil {
			return nil, err
		}
		e.Values = append(e.Values, v)
	}
	return e, err
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

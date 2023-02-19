package parser

import (
	"bytes"
	"errors"
	"fmt"
	"io"
)

type Parser struct {
	BinaryParser
}

func New(input io.Reader) *Parser {
	l := &Parser{BinaryParser: NewBinaryParser(input)}
	return l
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
	if c.Attributes, err = readAttributes(p, c.ConstantPool); err != nil {
		return nil, err
	}
	return c, nil
}

func (p *Parser) readCaffbabe() error {
	bs, err := p.ReadBytes(4)
	if err != nil {
		return err
	}
	if !bytes.Equal(bs, []byte{0xCA, 0xFE, 0xBA, 0xBE}) {
		return errors.New("magic is wrong")
	}
	return nil
}

func (p *Parser) readMinorVersion(c *Classfile) error {
	v, err := p.ReadUint16()
	if err != nil {
		return err
	}
	c.MinorVersion = v
	return nil
}

func (p *Parser) readMajorVersion(c *Classfile) error {
	v, err := p.ReadUint16()
	if err != nil {
		return err
	}
	c.MajorVersion = v
	return nil
}

func (p *Parser) readConstantPool(c *Classfile) error {
	count, err := p.ReadUint16()
	if err != nil {
		return err
	}
	var i uint16
	cp := &ConstantPool{Constants: make([]Constant, count-1)}
	c.ConstantPool = cp
	for ; i < count-1; i++ {
		tag, err := p.ReadUint8()
		if err != nil {
			return nil
		}
		switch tag {
		case 7:
			c := &ConstantClass{}
			cp.Constants[i] = c
			c.NameIndex, err = p.ReadUint16()
			if err != nil {
				return err
			}
		case 9:
			c := &ConstantFieldref{}
			cp.Constants[i] = c
			c.ClassIndex, err = p.ReadUint16()
			if err != nil {
				return err
			}
			c.NameAndTypeIndex, err = p.ReadUint16()
			if err != nil {
				return err
			}
		case 10:
			c := &ConstantMethodref{}
			cp.Constants[i] = c
			c.ClassIndex, err = p.ReadUint16()
			if err != nil {
				return err
			}
			c.NameAndTypeIndex, err = p.ReadUint16()
			if err != nil {
				return err
			}
		case 11:
			c := &ConstantInterfaceMethodref{}
			cp.Constants[i] = c
			c.ClassIndex, err = p.ReadUint16()
			if err != nil {
				return err
			}
			c.NameAndTypeIndex, err = p.ReadUint16()
			if err != nil {
				return err
			}
		case 8:
			c := &ConstantString{}
			cp.Constants[i] = c
			c.StringIndex, err = p.ReadUint16()
			if err != nil {
				return err
			}
		case 3:
			c := &ConstantInteger{}
			cp.Constants[i] = c
			c.Bytes, err = p.ReadUint32()
			if err != nil {
				return err
			}
		case 4:
			c := &ConstantFloat{}
			cp.Constants[i] = c
			c.Bytes, err = p.ReadUint32()
			if err != nil {
				return err
			}
		case 5:
			c := &ConstantLong{}
			cp.Constants[i] = c
			i++
			c.HighBytes, err = p.ReadUint32()
			if err != nil {
				return err
			}
			c.LowBytes, err = p.ReadUint32()
			if err != nil {
				return err
			}
		case 6:
			c := &ConstantDouble{}
			cp.Constants[i] = c
			i++
			c.HighBytes, err = p.ReadUint32()
			if err != nil {
				return err
			}
			c.LowBytes, err = p.ReadUint32()
			if err != nil {
				return err
			}
		case 12:
			c := &ConstantNameAndType{}
			cp.Constants[i] = c
			c.NameIndex, err = p.ReadUint16()
			if err != nil {
				return err
			}
			c.DescriptorIndex, err = p.ReadUint16()
			if err != nil {
				return err
			}
		case 1:
			c := &ConstantUtf8{}
			cp.Constants[i] = c
			count, err := p.ReadUint16()
			if err != nil {
				return err
			}

			c.Bytes, err = p.ReadBytes(int(count))
			if err != nil {
				return err
			}
		case 15:
			c := &ConstantMethodHandle{}
			cp.Constants[i] = c
			c.ReferenceKind, err = p.ReadUint8()
			if err != nil {
				return err
			}
			c.ReferenceIndex, err = p.ReadUint16()
			if err != nil {
				return err
			}
		case 16:
			c := &ConstantMethodType{}
			cp.Constants[i] = c
			c.DescriptorIndex, err = p.ReadUint16()
			if err != nil {
				return err
			}
		case 17:
			c := &ConstantDynamic{}
			cp.Constants[i] = c
			c.BootstrapMethodAttrIndex, err = p.ReadUint16()
			if err != nil {
				return err
			}
			c.NameAndTypeIndex, err = p.ReadUint16()
			if err != nil {
				return err
			}
		case 18:
			c := &ConstantInvokeDynamic{}
			cp.Constants[i] = c
			c.BootstrapMethodAttrIndex, err = p.ReadUint16()
			if err != nil {
				return err
			}
			c.NameAndTypeIndex, err = p.ReadUint16()
			if err != nil {
				return err
			}
		case 19:
			c := &ConstantModule{}
			cp.Constants[i] = c
			c.NameIndex, err = p.ReadUint16()
			if err != nil {
				return err
			}
		case 20:
			c := &ConstantPackage{}
			cp.Constants[i] = c
			c.NameIndex, err = p.ReadUint16()
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
	c.AccessFlags, err = p.ReadUint16()
	return
}

func (p *Parser) readThisClass(c *Classfile) (err error) {
	c.ThisClass, err = p.ReadUint16()
	return
}

func (p *Parser) readSuperClass(c *Classfile) (err error) {
	c.SuperClass, err = p.ReadUint16()
	return
}

func (p *Parser) readInterfaces(c *Classfile) error {
	count, err := p.ReadUint16()
	if err != nil {
		return err
	}
	var i uint16
	for ; i < count; i++ {
		interfaceIndex, err := p.ReadUint16()
		if err != nil {
			return err
		}
		c.Interfaces = append(c.Interfaces, interfaceIndex)
	}
	return nil
}

func (p *Parser) readFields(c *Classfile) error {
	count, err := p.ReadUint16()
	if err != nil {
		return err
	}
	var i uint16
	for ; i < count; i++ {
		f := &Field{}
		c.Fields = append(c.Fields, f)
		f.AccessFlags, err = p.ReadUint16()
		if err != nil {
			return err
		}
		f.NameIndex, err = p.ReadUint16()
		if err != nil {
			return err
		}
		f.DescriptorIndex, err = p.ReadUint16()
		if err != nil {
			return err
		}
		f.Attributes, err = readAttributes(p, c.ConstantPool)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *Parser) readMethods(c *Classfile) error {
	count, err := p.ReadUint16()
	if err != nil {
		return err
	}
	var i uint16
	for ; i < count; i++ {
		m := &Method{}
		c.Methods = append(c.Methods, m)
		m.AccessFlags, err = p.ReadUint16()
		if err != nil {
			return err
		}
		m.NameIndex, err = p.ReadUint16()
		if err != nil {
			return err
		}
		m.DescriptorIndex, err = p.ReadUint16()
		if err != nil {
			return err
		}
		m.Attributes, err = readAttributes(p, c.ConstantPool)
		if err != nil {
			return err
		}
	}
	return nil
}

func readAttributes(p BinaryParser, c *ConstantPool) ([]Attribute, error) {
	count, err := p.ReadUint16()
	if err != nil {
		return nil, err
	}
	as := make([]Attribute, 0, count)
	var i uint16
	for ; i < count; i++ {
		attributeNameIndex, err := p.ReadUint16()
		if err != nil {
			return nil, err
		}
		attributeLength, err := p.ReadUint32()
		if err != nil {
			return nil, err
		}
		u := c.LookupUtf8(attributeNameIndex)
		if u == nil {
			return nil, fmt.Errorf("attribute name index is invalid: index:%d", attributeNameIndex)
		}
		bs, err := p.ReadBytes(int(attributeLength))
		if err != nil {
			return nil, err
		}
		parser := NewBinaryParser(bytes.NewBuffer(bs))
		a, err := readAttribute(parser, attributeLength, u.String(), c)
		if err != nil {
			return nil, err
		}
		as = append(as, a)
	}
	return as, nil
}

func readAttribute(p BinaryParser, attributeLength uint32, attributeName string, constantPool *ConstantPool) (Attribute, error) {
	var err error
	switch attributeName {
	case "ConstantValue":
		a := &AttributeConstantValue{}
		a.ConstantValueIndex, err = p.ReadUint16()
		return a, err
	case "Code":
		a := &AttributeCode{}
		a.MaxStack, err = p.ReadUint16()
		if err != nil {
			return nil, err
		}
		a.MaxLocals, err = p.ReadUint16()
		if err != nil {
			return nil, err
		}
		codeLength, err := p.ReadUint32()
		if err != nil {
			return nil, err
		}
		var i uint32
		for ; i < codeLength; i++ {
			code, err := p.ReadUint8()
			if err != nil {
				return nil, err
			}
			a.Codes = append(a.Codes, code)
		}

		exceptionTableLength, err := p.ReadUint16()
		if err != nil {
			return nil, err
		}
		var j uint16
		for ; j < exceptionTableLength; j++ {
			e := &Exception{}
			e.StartPc, err = p.ReadUint16()
			if err != nil {
				return nil, err
			}
			e.EndPc, err = p.ReadUint16()
			if err != nil {
				return nil, err
			}
			e.HandlerPc, err = p.ReadUint16()
			if err != nil {
				return nil, err
			}
			e.CatchType, err = p.ReadUint16()
			if err != nil {
				return nil, err
			}
			a.ExceptionTable = append(a.ExceptionTable, e)
		}

		return a, nil
	case "StackMapTable":
		numOfEntries, err := p.ReadUint16()
		if err != nil {
			return nil, err
		}
		a := &AttributeStackMapTable{}
		var i uint16
		for ; i < numOfEntries; i++ {
			e, err := readStackMapFrame(p)
			if err != nil {
				return nil, err
			}
			a.Entries = append(a.Entries, e)
		}
		return a, nil
	case "Exceptions":
		exceptionCount, err := p.ReadUint16()
		if err != nil {
			return nil, err
		}
		a := &AttributeExceptions{ExceptionIndexes: make([]uint16, 0, exceptionCount)}
		var i uint16
		for ; i < exceptionCount; i++ {
			exceptionIndex, err := p.ReadUint16()
			if err != nil {
				return nil, err
			}
			a.ExceptionIndexes = append(a.ExceptionIndexes, exceptionIndex)
		}
		return a, nil
	case "InnerClasses":
		numberOfClasses, err := p.ReadUint16()
		if err != nil {
			return nil, err
		}
		a := &AttributeInnerClasses{InnerClasses: make([]*InnerClass, 0, numberOfClasses)}
		var i uint16
		for ; i < numberOfClasses; i++ {
			c := &InnerClass{}
			c.InnerClassInfoIndex, err = p.ReadUint16()
			if err != nil {
				return nil, err
			}
			c.OuterClassInfoIndex, err = p.ReadUint16()
			if err != nil {
				return nil, err
			}
			c.InnerNameIndex, err = p.ReadUint16()
			if err != nil {
				return nil, err
			}
			c.InnerClassAccessFlags, err = p.ReadUint16()
			if err != nil {
				return nil, err
			}
			a.InnerClasses = append(a.InnerClasses, c)
		}
		return a, nil
	case "EnclosingMethod":
		a := &AttributeEnclosingMethod{}
		a.ClassIndex, err = p.ReadUint16()
		if err != nil {
			return nil, err
		}
		a.MethodIndex, err = p.ReadUint16()
		return a, err
	case "Synthetic":
		return synthetic, nil
	case "Signature":
		a := &AttributeSignature{}
		a.Signature, err = p.ReadUint16()
		return a, err
	case "SourceFile":
		a := &AttributeSourceFile{}
		a.SourcefileIndex, err = p.ReadUint16()
		return a, err
	case "SourceDebugExtension":
		a := &AttributeSourceDebugExtension{}
		a.DebugExtension, err = p.ReadBytes(int(attributeLength))
		return a, err
	case "LineNumberTable":
		lineNumberTableLength, err := p.ReadUint16()
		if err != nil {
			return nil, err
		}
		a := &AttributeLineNumberTable{LineNumberTable: make([]*LineNumber, 0, lineNumberTableLength)}
		var i uint16
		for ; i < lineNumberTableLength; i++ {
			ln := &LineNumber{}
			ln.StartPc, err = p.ReadUint16()
			if err != nil {
				return nil, err
			}
			ln.LineNumber, err = p.ReadUint16()
			if err != nil {
				return nil, err
			}
			a.LineNumberTable = append(a.LineNumberTable, ln)
		}
		return a, nil
	case "LocalVariableTable":
		localVaribleTableLength, err := p.ReadUint16()
		if err != nil {
			return nil, err
		}
		a := &AttributeLocalVariableTable{}
		var i uint16
		for ; i < localVaribleTableLength; i++ {
			ln := &LocalVariable{}
			ln.StartPc, err = p.ReadUint16()
			if err != nil {
				return nil, err
			}
			ln.Length, err = p.ReadUint16()
			if err != nil {
				return nil, err
			}
			ln.NameIndex, err = p.ReadUint16()
			if err != nil {
				return nil, err
			}
			ln.DescriptorInedx, err = p.ReadUint16()
			if err != nil {
				return nil, err
			}
			ln.Index, err = p.ReadUint16()
			if err != nil {
				return nil, err
			}
			a.LocalVaribleTable = append(a.LocalVaribleTable, ln)
		}
		return a, nil
	case "LocalVariableTypeTable":
		localVaribleTypeLength, err := p.ReadUint16()
		if err != nil {
			return nil, err
		}
		a := &AttributeLocalVariableTypeTable{}
		var i uint16
		for ; i < localVaribleTypeLength; i++ {
			ln := &LocalVariableType{}
			ln.StartPc, err = p.ReadUint16()
			if err != nil {
				return nil, err
			}
			ln.Length, err = p.ReadUint16()
			if err != nil {
				return nil, err
			}
			ln.NameIndex, err = p.ReadUint16()
			if err != nil {
				return nil, err
			}
			ln.SignatureInedx, err = p.ReadUint16()
			if err != nil {
				return nil, err
			}
			ln.Index, err = p.ReadUint16()
			if err != nil {
				return nil, err
			}
			a.LocalVaribleTypeTable = append(a.LocalVaribleTypeTable, ln)
		}
		return a, nil
	case "Deprecated":
		return deprecated, nil
	case "RuntimeVisibleAnnotations":
		numAnnotations, err := p.ReadUint16()
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
		numAnnotations, err := p.ReadUint16()
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
		numParameters, err := p.ReadUint8()
		if err != nil {
			return nil, err
		}
		a := &AttributeRuntimeVisibleParameterAnnotations{}
		var i uint8
		for ; i < numParameters; i++ {
			annot, err := readParameterAnnotation(p)
			if err != nil {
				return nil, err
			}
			a.ParameterAnnotations = append(a.ParameterAnnotations, annot)
		}
		return a, nil
	case "RuntimeInvisibleParameterAnnotations":
		numParameters, err := p.ReadUint8()
		if err != nil {
			return nil, err
		}
		a := &AttributeRuntimeVisibleParameterAnnotations{}
		var i uint8
		for ; i < numParameters; i++ {
			annot, err := readParameterAnnotation(p)
			if err != nil {
				return nil, err
			}
			a.ParameterAnnotations = append(a.ParameterAnnotations, annot)
		}
		return a, nil
	case "RuntimeVisibleTypeAnnotations":
		numAnnotations, err := p.ReadUint16()
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
		numAnnotations, err := p.ReadUint16()
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
	case "BootstrapMethods":
		numBootstrapMethods, err := p.ReadUint16()
		if err != nil {
			return nil, err
		}
		attribute := &AttributeBootstrapMethods{}
		var i uint16
		for ; i < numBootstrapMethods; i++ {
			b := &BootstrapMethod{}
			b.BootstrapMethodRef, err = p.ReadUint16()
			if err != nil {
				return nil, err
			}
			numBootstrapArguments, err := p.ReadUint16()
			if err != nil {
				return nil, err
			}
			var j uint16
			for ; j < numBootstrapArguments; j++ {
				a, err := p.ReadUint16()
				if err != nil {
					return nil, err
				}
				b.BootstrapArguments = append(b.BootstrapArguments, a)
			}
			attribute.BootstrapMethods = append(attribute.BootstrapMethods, b)
		}
		return attribute, nil
	case "MethodParameters":
		parametersCount, err := p.ReadUint8()
		if err != nil {
			return nil, err
		}
		a := &AttributeMethodParameters{}
		var i uint8
		for ; i < parametersCount; i++ {
			param := &MethodParameter{}
			param.NameIndex, err = p.ReadUint16()
			if err != nil {
				return nil, err
			}
			param.AccessFlags, err = p.ReadUint16()
			if err != nil {
				return nil, err
			}
			a.Parameters = append(a.Parameters, param)
		}
		return a, nil
	case "Module":
		a := &AttributeModule{}
		a.ModuleNameIndex, err = p.ReadUint16()
		if err != nil {
			return nil, err
		}
		a.ModuleFlags, err = p.ReadUint16()
		if err != nil {
			return nil, err
		}
		a.ModuleVersionIndex, err = p.ReadUint16()
		if err != nil {
			return nil, err
		}
		requiresCount, err := p.ReadUint16()
		if err != nil {
			return nil, err
		}

		var i uint16
		for ; i < requiresCount; i++ {
			r := &Require{}
			r.RequiresIndex, err = p.ReadUint16()
			if err != nil {
				return nil, err
			}
			r.RequiresFlags, err = p.ReadUint16()
			if err != nil {
				return nil, err
			}
			r.RequiresVersionIndex, err = p.ReadUint16()
			if err != nil {
				return nil, err
			}
			a.Requires = append(a.Requires, r)
		}

		exportsCount, err := p.ReadUint16()
		if err != nil {
			return nil, err
		}
		for i = 0; i < exportsCount; i++ {
			e := &Export{}
			e.ExportsIndex, err = p.ReadUint16()
			if err != nil {
				return nil, err
			}
			e.ExportsFlags, err = p.ReadUint16()
			if err != nil {
				return nil, err
			}
			exportsToCount, err := p.ReadUint16()
			if err != nil {
				return nil, err
			}
			var j uint16
			for ; j < exportsToCount; j++ {
				t, err := p.ReadUint16()
				if err != nil {
					return nil, err
				}
				e.ExportsTo = append(e.ExportsTo, t)
			}
		}

		opensCount, err := p.ReadUint16()
		if err != nil {
			return nil, err
		}
		for i = 0; i < opensCount; i++ {
			e := &Open{}
			e.OpensIndex, err = p.ReadUint16()
			if err != nil {
				return nil, err
			}
			e.OpensFlags, err = p.ReadUint16()
			if err != nil {
				return nil, err
			}
			opensToCount, err := p.ReadUint16()
			if err != nil {
				return nil, err
			}
			var j uint16
			for ; j < opensToCount; j++ {
				t, err := p.ReadUint16()
				if err != nil {
					return nil, err
				}
				e.OpensTo = append(e.OpensTo, t)
			}
		}

		usesCount, err := p.ReadUint16()
		if err != nil {
			return nil, err
		}
		for i = 0; i < usesCount; i++ {
			use, err := p.ReadUint16()
			if err != nil {
				return nil, err
			}
			a.Uses = append(a.Uses, use)
		}

		providesCount, err := p.ReadUint16()
		if err != nil {
			return nil, err
		}
		for i = 0; i < providesCount; i++ {
			e := &Provide{}
			e.ProvidesIndex, err = p.ReadUint16()
			if err != nil {
				return nil, err
			}
			providesWithCount, err := p.ReadUint16()
			if err != nil {
				return nil, err
			}
			var j uint16
			for ; j < providesWithCount; j++ {
				t, err := p.ReadUint16()
				if err != nil {
					return nil, err
				}
				e.ProvidesWith = append(e.ProvidesWith, t)
			}
		}

		return a, nil
	case "ModulePackages":
		packageCount, err := p.ReadUint16()
		if err != nil {
			return nil, err
		}
		a := &AttributeModulePackages{PackageIndexes: make([]uint16, 0, packageCount)}
		var i uint16
		for ; i < packageCount; i++ {
			packageIndex, err := p.ReadUint16()
			if err != nil {
				return nil, err
			}
			a.PackageIndexes = append(a.PackageIndexes, packageIndex)
		}
		return a, nil
	case "ModuleMainClass":
		a := &AttributeModuleMainClass{}
		a.MainClassIndex, err = p.ReadUint16()
		return a, err
	case "NestHost":
		a := &AttributeNestHost{}
		a.HostClassIndex, err = p.ReadUint16()
		return a, err
	case "NestMembers":
		numberOfClasses, err := p.ReadUint16()
		if err != nil {
			return nil, err
		}
		a := &AttributeNestMembers{Classes: make([]uint16, 0, numberOfClasses)}
		var i uint16
		for ; i < numberOfClasses; i++ {
			class, err := p.ReadUint16()
			if err != nil {
				return nil, err
			}
			a.Classes = append(a.Classes, class)
		}
		return a, nil
	case "Record":
		a := &AttributeRecord{}
		componentsCount, err := p.ReadUint16()
		if err != nil {
			return nil, err
		}
		var i uint16
		for ; i < componentsCount; i++ {
			ci, err := readRecordComponentInfo(p, constantPool)
			if err != nil {
				return nil, err
			}
			a.Components = append(a.Components, ci)
		}
		return a, nil
	default:
		return nil, errors.New("Unknown attributes:" + attributeName)
	}
}

func readRecordComponentInfo(parser BinaryParser, c *ConstantPool) (RecordComponentInfo, error) {
	i := RecordComponentInfo{}
	var err error
	i.NameIndex, err = parser.ReadUint16()
	if err != nil {
		return i, err
	}
	i.DescriptorIndex, err = parser.ReadUint16()
	if err != nil {
		return i, err
	}
	i.Attributes, err = readAttributes(parser, c)
	return i, err
}

func readStackMapFrame(parser BinaryParser) (StackMapFrame, error) {
	frameType, err := parser.ReadUint8()
	if err != nil {
		return nil, err
	}
	switch {
	case frameType <= 63:
		return &StackMapFrameSameFrame{FrameType: frameType}, nil
	case 64 <= frameType && frameType <= 127:
		f := &StackMapFrameSameLocals1StackItemFrame{
			FrameType: frameType,
		}
		f.stack, err = readVerificationType(parser)
		return f, err
	case frameType == 247:
		f := &StackMapFrameSameLocals1StackItemFrameExtended{
			FrameType: frameType,
		}
		f.stack, err = readVerificationType(parser)
		return f, err
	case 248 <= frameType && frameType <= 250:
		f := &StackMapFrameChopFrame{FrameType: frameType}
		f.OffsetDelta, err = parser.ReadUint16()
		return f, err
	case frameType == 251:
		f := &StackMapFrameSameFrameExtended{FrameType: frameType}
		f.OffsetDelta, err = parser.ReadUint16()
		return f, err
	case 252 <= frameType && frameType <= 254:
		f := &StackMapFrameAppendFrame{FrameType: frameType}
		f.OffsetDelta, err = parser.ReadUint16()
		if err != nil {
			return nil, err
		}
		for i := frameType - 251; i > 0; i-- {
			vt, err := readVerificationType(parser)
			if err != nil {
				return nil, err
			}
			f.Locals = append(f.Locals, vt)
		}
		return f, nil
	case frameType == 255:
		f := &StackMapFrameFullFrame{}
		f.OffsetDelta, err = parser.ReadUint16()
		if err != nil {
			return nil, err
		}
		numberOfLocals, err := parser.ReadUint16()
		if err != nil {
			return nil, err
		}
		var i uint16
		for ; i < numberOfLocals; i++ {
			vt, err := readVerificationType(parser)
			if err != nil {
				return nil, err
			}
			f.Locals = append(f.Locals, vt)
		}
		numberOfStacks, err := parser.ReadUint16()
		if err != nil {
			return nil, err
		}
		for i = 0; i < numberOfStacks; i++ {
			vt, err := readVerificationType(parser)
			if err != nil {
				return nil, err
			}
			f.Stacks = append(f.Stacks, vt)
		}
		return f, nil
	}
	return nil, errors.New("Not supported frame type")
}

func readVerificationType(parser BinaryParser) (VerificationTypeInfo, error) {
	tag, err := parser.ReadUint8()
	if err != nil {
		return nil, err
	}
	switch tag {
	case 0:
		return _verificationTypeInfoTopVaribleInfo, nil
	case 1:
		return _verificationTypeInfoIntegerVaribleInfo, nil
	case 2:
		return _verificationTypeInfoFloatVaribleInfo, nil
	case 5:
		return _verificationTypeInfoNullVaribleInfo, nil
	case 6:
		return _verificationTypeInfoUninitializedThisVaribleInfo, nil
	case 7:
		i := &VerificationTypeInfoObjectVaribleInfo{}
		i.CpoolIndex, err = parser.ReadUint16()
		return i, err
	case 8:
		i := &VerificationTypeInfoUninitializedVaribleInfo{}
		i.Offset, err = parser.ReadUint16()
		return i, err
	case 4:
		return _verificationTypeInfoLongVaribleInfo, nil
	case 3:
		return _verificationTypeInfoDoubleVaribleInfo, nil
	}
	return nil, errors.New("Unsupported verification type info")
}

func readAnnotation(parser BinaryParser) (*Annotation, error) {
	a := &Annotation{}
	var err error
	a.TypeIndex, err = parser.ReadUint16()
	if err != nil {
		return nil, err
	}
	a.ElementValuePairs, err = readElementValuePairs(parser)
	return a, err
}

func readElementValuePairs(parser BinaryParser) ([]*ElementValuePair, error) {
	numElementValuePair, err := parser.ReadUint16()
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

func readElementValuePair(parser BinaryParser) (*ElementValuePair, error) {
	p := &ElementValuePair{}
	var err error
	p.ElementNameIndex, err = parser.ReadUint16()
	if err != nil {
		return nil, err
	}
	p.ElementValue, err = readElementValue(parser)
	return p, err
}

func readParameterAnnotation(parser BinaryParser) (*ParameterAnnotation, error) {
	a := &ParameterAnnotation{}
	numAnnotations, err := parser.ReadUint16()
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

func readTypeAnnotation(parser BinaryParser) (*TypeAnnotation, error) {
	a := &TypeAnnotation{}
	targetType, err := parser.ReadUint8()
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
	if err != nil {
		return nil, err
	}
	a.TargetPath, err = readTypePath(parser)
	if err != nil {
		return nil, err
	}
	a.TypeIndex, err = parser.ReadUint16()
	if err != nil {
		return nil, err
	}
	a.ElementValuePairs, err = readElementValuePairs(parser)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func readTypeParameterTarget(parser BinaryParser) (TargetInfo, error) {
	target := &TypeParameterTarget{}
	var err error
	target.TypeParameterIndex, err = parser.ReadUint8()
	return target, err
}

func readSuperTypeTarget(parser BinaryParser) (TargetInfo, error) {
	target := &SuperTypeTarget{}
	var err error
	target.SuperTypeIndex, err = parser.ReadUint16()
	return target, err
}

func readTypeParameterBoundTarget(parser BinaryParser) (TargetInfo, error) {
	target := &TypeParameterBoundTarget{}
	var err error
	target.TypeParameterIndex, err = parser.ReadUint8()
	if err != nil {
		return nil, err
	}
	target.BoundIndex, err = parser.ReadUint8()
	return target, err
}

func readFormalParameterTarget(parser BinaryParser) (TargetInfo, error) {
	target := &FormalParameterTarget{}
	var err error
	target.FormalParameterIndex, err = parser.ReadUint8()
	return target, err
}

func readThrowsTarget(parser BinaryParser) (TargetInfo, error) {
	target := &ThrowsTarget{}
	var err error
	target.ThrowsTypeIndex, err = parser.ReadUint16()
	return target, err
}

func readLocalVarTarget(parser BinaryParser) (TargetInfo, error) {
	target := &LocalVarTarget{}
	tableLength, err := parser.ReadUint16()
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

func readLocalVarTargetTable(parser BinaryParser) (*LocalVarTargetTable, error) {
	t := &LocalVarTargetTable{}
	var err error
	t.StartPc, err = parser.ReadUint16()
	if err != nil {
		return nil, err
	}
	t.Length, err = parser.ReadUint16()
	if err != nil {
		return nil, err
	}
	t.Index, err = parser.ReadUint16()
	return t, err
}

func readCatchTarget(parser BinaryParser) (TargetInfo, error) {
	target := &CatchTarget{}
	var err error
	target.ExceptionTableIndex, err = parser.ReadUint16()
	return target, err
}

func readOffsetTarget(parser BinaryParser) (TargetInfo, error) {
	target := &OffsetTarget{}
	var err error
	target.Offset, err = parser.ReadUint16()
	return target, err
}

func readTypeArgumentTarget(parser BinaryParser) (TargetInfo, error) {
	target := &TypeArgumentTarget{}
	var err error
	target.Offset, err = parser.ReadUint16()
	if err != nil {
		return nil, err
	}
	target.TypeArgumentIndex, err = parser.ReadUint8()
	return target, err
}

func readTypePath(parser BinaryParser) (*TypePath, error) {
	pathLength, err := parser.ReadUint8()
	if err != nil {
		return nil, err
	}
	typePath := &TypePath{}
	var i uint8
	for ; i < pathLength; i++ {
		path := &Path{}
		path.TypePathKind, err = parser.ReadUint8()
		if err != nil {
			return nil, err
		}
		path.TypeArgumentIndex, err = parser.ReadUint8()
		if err != nil {
			return nil, err
		}
		typePath.Paths = append(typePath.Paths, path)
	}
	return typePath, nil
}

func readElementValue(parser BinaryParser) (ElementValue, error) {
	tag, err := parser.ReadUint8()
	if err != nil {
		return nil, err
	}
	switch tag {
	case 'B':
		return readElementValueConst(parser)
	case 'C':
		return readElementValueConst(parser)
	case 'F':
		return readElementValueConst(parser)
	case 'I':
		return readElementValueConst(parser)
	case 'J':
		return readElementValueConst(parser)
	case 'S':
		return readElementValueConst(parser)
	case 'Z':
		return readElementValueConst(parser)
	case 's':
		return readElementValueConst(parser)
	case 'e':
		return readElementValueEnumConst(parser)
	case 'c':
		return readElementClassInfo(parser)
	case '@':
		return readAnnotation(parser)
	case '[':
		return readElementArrayValue(parser)
	default:
		return nil, errors.New("Unsupported tag for element value")
	}
}

func readElementValueConst(parser BinaryParser) (ElementValue, error) {
	e := &ElementValueConstValue{}
	var err error
	e.ConstValueIndex, err = parser.ReadUint16()
	return e, err
}

func readElementValueEnumConst(parser BinaryParser) (*ElementValueEnumConstValue, error) {
	e := &ElementValueEnumConstValue{}
	var err error
	e.TypeNameIndex, err = parser.ReadUint16()
	if err != nil {
		return nil, err
	}
	e.ConstNameIndex, err = parser.ReadUint16()
	return e, err
}

func readElementClassInfo(parser BinaryParser) (*ElementValueClassInfo, error) {
	e := &ElementValueClassInfo{}
	var err error
	e.ClassInfoIndex, err = parser.ReadUint16()
	return e, err
}

func readElementArrayValue(parser BinaryParser) (*ElementValueArrayValue, error) {
	e := &ElementValueArrayValue{}
	numValues, err := parser.ReadUint16()
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

package parser

type Field struct {
	AccessFlags     uint16
	NameIndex       uint16
	DescriptorIndex uint16
	Attributes      []Attribute
}

func (f *Field) Name(c *ConstantPool) (string, error) {
	name, err := c.GetConstantUtf8(f.NameIndex)
	if err != nil {
		return "", err
	}
	return name.String(), nil
}

func (f *Field) Descriptor(c *ConstantPool) (string, error) {
	name, err := c.GetConstantUtf8(f.DescriptorIndex)
	if err != nil {
		return "", err
	}
	return name.String(), nil
}

func (f *Field) ConstantValue() *AttributeConstantValue {
	for _, e := range f.Attributes {
		if attr, ok := e.(*AttributeConstantValue); ok {
			return attr
		}
	}
	return nil
}

func (f *Field) Synthetic() *AttributeSynthetic {
	for _, e := range f.Attributes {
		if attr, ok := e.(*AttributeSynthetic); ok {
			return attr
		}
	}
	return nil
}

func (f *Field) Deprecated() *AttributeDeprecated {
	for _, e := range f.Attributes {
		if attr, ok := e.(*AttributeDeprecated); ok {
			return attr
		}
	}
	return nil
}

func (f *Field) Signature() *AttributeSignature {
	for _, e := range f.Attributes {
		if attr, ok := e.(*AttributeSignature); ok {
			return attr
		}
	}
	return nil
}

func (f *Field) RuntimeVisibleAnnotations() *AttributeRuntimeVisibleAnnotations {
	for _, e := range f.Attributes {
		if attr, ok := e.(*AttributeRuntimeVisibleAnnotations); ok {
			return attr
		}
	}
	return nil
}

func (f *Field) RuntimeInvisibleAnnotations() *AttributeRuntimeInvisibleAnnotations {
	for _, e := range f.Attributes {
		if attr, ok := e.(*AttributeRuntimeInvisibleAnnotations); ok {
			return attr
		}
	}
	return nil
}

func (f *Field) RuntimeVisibleTypeAnnotations() *AttributeRuntimeVisibleTypeAnnotations {
	for _, e := range f.Attributes {
		if attr, ok := e.(*AttributeRuntimeVisibleTypeAnnotations); ok {
			return attr
		}
	}
	return nil
}

func (f *Field) RuntimeInvisibleTypeAnnotations() *AttributeRuntimeInvisibleTypeAnnotations {
	for _, e := range f.Attributes {
		if attr, ok := e.(*AttributeRuntimeInvisibleTypeAnnotations); ok {
			return attr
		}
	}
	return nil
}

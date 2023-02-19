package parser

type Method struct {
	AccessFlags     uint16
	NameIndex       uint16
	DescriptorIndex uint16
	Attributes      []Attribute
}

func (m *Method) Name(c *ConstantPool) (string, error) {
	name, err := c.GetConstantUtf8(m.NameIndex)
	if err != nil {
		return "", err
	}
	return name.String(), nil
}

func (m *Method) Descriptor(c *ConstantPool) (string, error) {
	name, err := c.GetConstantUtf8(m.DescriptorIndex)
	if err != nil {
		return "", err
	}
	return name.String(), nil
}

func (m *Method) Code() *AttributeCode {
	for _, e := range m.Attributes {
		if attr, ok := e.(*AttributeCode); ok {
			return attr
		}
	}
	return nil
}

func (m *Method) Exceptions() *AttributeExceptions {
	for _, e := range m.Attributes {
		if attr, ok := e.(*AttributeExceptions); ok {
			return attr
		}
	}
	return nil
}

func (m *Method) RuntimeVisibleParameterAnnotations() *AttributeRuntimeVisibleParameterAnnotations {
	for _, e := range m.Attributes {
		if attr, ok := e.(*AttributeRuntimeVisibleParameterAnnotations); ok {
			return attr
		}
	}
	return nil
}

func (m *Method) RuntimeInisibleParameterAnnotations() *AttributeRuntimeInvisibleParameterAnnotations {
	for _, e := range m.Attributes {
		if attr, ok := e.(*AttributeRuntimeInvisibleParameterAnnotations); ok {
			return attr
		}
	}
	return nil
}

func (m *Method) AnnotationDefault() *AttributeAnnotationDefault {
	for _, e := range m.Attributes {
		if attr, ok := e.(*AttributeAnnotationDefault); ok {
			return attr
		}
	}
	return nil
}

func (m *Method) MethodParameters() *AttributeMethodParameters {
	for _, e := range m.Attributes {
		if attr, ok := e.(*AttributeMethodParameters); ok {
			return attr
		}
	}
	return nil
}

func (m *Method) Synthetic() *AttributeSynthetic {
	for _, e := range m.Attributes {
		if attr, ok := e.(*AttributeSynthetic); ok {
			return attr
		}
	}
	return nil
}

func (m *Method) Deprecated() *AttributeDeprecated {
	for _, e := range m.Attributes {
		if attr, ok := e.(*AttributeDeprecated); ok {
			return attr
		}
	}
	return nil
}

func (m *Method) Signature() *AttributeSignature {
	for _, e := range m.Attributes {
		if attr, ok := e.(*AttributeSignature); ok {
			return attr
		}
	}
	return nil
}

func (m *Method) RuntimeVisibleAnnotations() *AttributeRuntimeVisibleAnnotations {
	for _, e := range m.Attributes {
		if attr, ok := e.(*AttributeRuntimeVisibleAnnotations); ok {
			return attr
		}
	}
	return nil
}

func (m *Method) RuntimeInvisibleAnnotations() *AttributeRuntimeInvisibleAnnotations {
	for _, e := range m.Attributes {
		if attr, ok := e.(*AttributeRuntimeInvisibleAnnotations); ok {
			return attr
		}
	}
	return nil
}

func (m *Method) RuntimeVisibleTypeAnnotations() *AttributeRuntimeVisibleTypeAnnotations {
	for _, e := range m.Attributes {
		if attr, ok := e.(*AttributeRuntimeVisibleTypeAnnotations); ok {
			return attr
		}
	}
	return nil
}

func (m *Method) RuntimeInvisibleTypeAnnotations() *AttributeRuntimeInvisibleTypeAnnotations {
	for _, e := range m.Attributes {
		if attr, ok := e.(*AttributeRuntimeInvisibleTypeAnnotations); ok {
			return attr
		}
	}
	return nil
}

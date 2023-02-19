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

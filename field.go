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

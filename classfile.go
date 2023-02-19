package parser

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
	Attributes   []Attribute
}

func (c *Classfile) ReadThisClass() (string, error) {
	return c.ConstantPool.GetClassName(c.ThisClass)
}

func (c *Classfile) ReadSuperClass() (string, error) {
	return c.ConstantPool.GetClassName(c.SuperClass)
}

package main

type Field struct {
	AccessFlags     uint16
	NameIndex       uint16
	DescriptorIndex uint16
	Attributes      []Attribute
}

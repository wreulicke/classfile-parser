package parser

type elementValue struct{}

func (*elementValue) elementValueType() {}

type ElementValue interface {
	elementValueType()
}

type ElementValueConstValue struct {
	elementValue
	ConstValueIndex uint16
}

type ElementValueEnumConstValue struct {
	elementValue
	TypeNameIndex  uint16
	ConstNameIndex uint16
}

type ElementValueClassInfo struct {
	elementValue
	ClassInfoIndex uint16
}

type ElementValueArrayValue struct {
	elementValue
	Values []ElementValue
}

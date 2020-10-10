package main

type Attribute interface {
	Name() string
}

type AttributeConstantValue struct {
	ConstantValueIndex uint16
}

func (a *AttributeConstantValue) Name() string {
	return "ConstantValue"
}

type AttributeExceptions struct {
	ExceptionIndexes []uint16
}

func (a *AttributeExceptions) Name() string {
	return "Exceptions"
}

type AttributeInnerClasses struct {
	InnerClasses []*InnerClass
}

func (a *AttributeInnerClasses) Name() string {
	return "InnerClasses"
}

type InnerClass struct {
	InnerClassInfoIndex   uint16
	OuterClassInfoIndex   uint16
	InnerNameIndex        uint16
	InnerClassAccessFlags uint16
}

type AttributeEnclosingMethod struct {
	ClassIndex  uint16
	MethodIndex uint16
}

func (a *AttributeEnclosingMethod) Name() string {
	return "EnclosingMethod"
}

type AttributeSynthetic struct {
	Signature uint16
}

func (a *AttributeSynthetic) Name() string {
	return "Synthetic"
}

type AttributeSignature struct {
	Signature uint16
}

func (a *AttributeSignature) Name() string {
	return "Signature"
}

type AttributeSourceFile struct {
	SourcefileIndex uint16
}

func (a *AttributeSourceFile) Name() string {
	return "SourceFile"
}

type AttributeSourceDebugExtension struct {
	DebugExtension []byte
}

func (a *AttributeSourceDebugExtension) Name() string {
	return "SourceDebugExtension"
}

type AttributeLineNumberTable struct {
	LineNumberTable []*LineNumber
}

func (a *AttributeLineNumberTable) Name() string {
	return "LineNumberTable"
}

type LineNumber struct {
	StartPc    uint16
	LineNumber uint16
}

type AttributeLocalVaribleTable struct {
	LocalVaribleTable []*LocalVarible
}

func (a *AttributeLocalVaribleTable) Name() string {
	return "LocalVaribleTable"
}

type LocalVarible struct {
	StartPc         uint16
	Length          uint16
	NameIndex       uint16
	DescriptorInedx uint16
	Index           uint16
}

type AttributeLocalVaribleTypeTable struct {
	LocalVaribleTypeTable []*LocalVaribleType
}

func (a *AttributeLocalVaribleTypeTable) Name() string {
	return "LocalVaribleTypeTable"
}

type LocalVaribleType struct {
	StartPc        uint16
	Length         uint16
	NameIndex      uint16
	SignatureInedx uint16
	Index          uint16
}

type AttributeDeprecated struct{}

func (a *AttributeDeprecated) Name() string {
	return "Deprecated"
}

type AttributeRuntimeVisibleAnnotations struct {
	AttributeNameIndex uint16
	Annotations        []*Annotation
}

func (a *AttributeRuntimeVisibleAnnotations) Name() string {
	return "RuntimeVisibleAnnotations"
}

type Annotation struct {
	TypeIndex         uint16
	ElementValuePairs []*ElementValuePair
}

type ElementValuePair struct {
	ElementNameIndex uint16
	ElementValue     *ElementValue
}

type ElementValue struct {
	Tag uint8

	ConstValue *ElementValueConstValue

	EnumConstValue *ElementValueEnumConstValue

	ClassInfo *ElementValueClassInfo

	AnnotationValue *Annotation

	ArrayValue *ElementValueArrayValue
}

type ElementValueConstValue struct {
	ConstValueIndex uint16
}

type ElementValueEnumConstValue struct {
	TypeNameIndex  uint16
	ConstNameIndex uint16
}

type ElementValueClassInfo struct {
	ClassInfoIndex uint16
}

type ElementValueArrayValue struct {
	Values []*ElementValue
}

type AttributeModulePackage struct {
	PackageIndexes []uint16
}

func (a *AttributeModulePackage) Name() string {
	return "Module"
}

type AttributeModuleMainClass struct {
	MainClassIndex uint16
}

func (a *AttributeModuleMainClass) Name() string {
	return "ModuleMainClass"
}

type AttributeNestHost struct {
	HostClassIndex uint16
}

func (a *AttributeNestHost) Name() string {
	return "NestHost"
}

type AttributeNestMembers struct {
	Classes []uint16
}

func (a *AttributeNestMembers) Name() string {
	return "NestHost"
}

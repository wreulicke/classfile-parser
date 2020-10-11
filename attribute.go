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

type AttributeCode struct {
	MaxStack       uint16
	MaxLocals      uint16
	Codes          []uint8
	ExceptionTable []*Exception
	Attributes     []Attribute
}

type Exception struct {
	StartPc   uint16
	EndPc     uint16
	HandlerPc uint16
	CatchType uint16
}

func (a *AttributeCode) Name() string {
	return "Code"
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
	Annotations []*Annotation
}

func (a *AttributeRuntimeVisibleAnnotations) Name() string {
	return "RuntimeVisibleAnnotations"
}

type AttributeRuntimeInvisibleAnnotations struct {
	Annotations []*Annotation
}

func (a *AttributeRuntimeInvisibleAnnotations) Name() string {
	return "RuntimeInvisibleAnnotations"
}

type AttributeRuntimeVisibleParameterAnnotations struct {
	ParameterAnnotations []*ParameterAnnotation
}

func (a *AttributeRuntimeVisibleParameterAnnotations) Name() string {
	return "RuntimeVisibleParameterAnnotations"
}

type AttributeRuntimeInvisibleParameterAnnotations struct {
	ParameterAnnotations []*ParameterAnnotation
}

func (a *AttributeRuntimeInvisibleParameterAnnotations) Name() string {
	return "RuntimeInvisibleParameterAnnotations"
}

type ParameterAnnotation struct {
	Annotations []*Annotation
}

type AttributeRuntimeVisibleTypeAnnotations struct {
	TypeAnnotations []*TypeAnnotation
}

func (a *AttributeRuntimeVisibleTypeAnnotations) Name() string {
	return "RuntimeVisibleTypeAnnotations"
}

type AttributeRuntimeInvisibleTypeAnnotations struct {
	TypeAnnotations []*TypeAnnotation
}

func (a *AttributeRuntimeInvisibleTypeAnnotations) Name() string {
	return "RuntimeInvisibleTypeAnnotations"
}

type TypeAnnotation struct {
	TargetType        uint8
	TargetInfo        TargetInfo
	TargetPath        *TypePath
	TypeIndex         uint16
	ElementValuePairs []*ElementValuePair
}

type TargetInfo interface{}

type TypeParameterTarget struct {
	TypeParameterIndex uint8
}

type SuperTypeTarget struct {
	SuperTypeIndex uint16
}

type TypeParameterBoundTarget struct {
	TypeParameterIndex uint8
	BoundIndex         uint8
}
type EmptyTarget struct{}

type FormalParameterTarget struct {
	FormalParameterIndex uint8
}

type ThrowsTarget struct {
	ThrowsTypeIndex uint16
}

type LocalVarTarget struct {
	LocalVarTargetTables []*LocalVarTargetTable
}

type CatchTarget struct {
	ExceptionTableIndex uint16
}

type OffsetTarget struct {
	Offset uint16
}

type TypeArgumentTarget struct {
	Offset            uint16
	TypeArgumentIndex uint8
}

type LocalVarTargetTable struct {
	StartPc uint16
	Length  uint16
	Index   uint16
}

type TypePath struct {
	Paths []*Path
}

type Path struct {
	TypePathKind      uint8
	TypeArgumentIndex uint8
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

type AttributeAnnotationDefault struct {
	DefaultValue *ElementValue
}

func (a *AttributeAnnotationDefault) Name() string {
	return "AnnotationDefault"
}

type AttributeBootstrapMethods struct {
	BootstrapMethods []*BootstrapMethod
}

func (a *AttributeBootstrapMethods) Name() string {
	return "BootstrapMethods"
}

type BootstrapMethod struct {
	BootstrapMethodRef uint16
	BootstrapArguments []uint16
}

type AttributeMethodParameters struct {
	Parameters []*MethodParameter
}

func (a *AttributeMethodParameters) Name() string {
	return "MethodParameters"
}

type MethodParameter struct {
	NameIndex   uint16
	AccessFlags uint16
}

type AttributeModule struct {
	ModuleNameIndex    uint16
	ModuleFlags        uint16
	ModuleVersionIndex uint16

	Requires []*Require
	Exports  []*Export
	Opens    []*Open
	Uses     []uint16
	Provides []*Provide
}

func (*AttributeModule) Name() string {
	return "Module"
}

type Require struct {
	RequiresIndex        uint16
	RequiresFlags        uint16
	RequiresVersionIndex uint16
}

type Export struct {
	ExportsIndex uint16
	ExportsFlags uint16
	ExportsTo    []uint16
}
type Open struct {
	OpensIndex uint16
	OpensFlags uint16
	OpensTo    []uint16
}

type Provide struct {
	ProvidesIndex uint16
	ProvidesWith  []uint16
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

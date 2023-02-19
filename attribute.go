package parser

var (
	deprecated = &AttributeDeprecated{}
	synthetic  = &AttributeSynthetic{}
)

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

func (a *AttributeCode) Name() string {
	return "Code"
}

func (a *AttributeCode) LineNumberTable() *AttributeLineNumberTable {
	for _, e := range a.Attributes {
		if attr, ok := e.(*AttributeLineNumberTable); ok {
			return attr
		}
	}
	return nil
}

func (a *AttributeCode) LocalVariableTable() *AttributeLocalVariableTable {
	for _, e := range a.Attributes {
		if attr, ok := e.(*AttributeLocalVariableTable); ok {
			return attr
		}
	}
	return nil
}

func (a *AttributeCode) LocalVariableTypeTable() *AttributeLocalVariableTypeTable {
	for _, e := range a.Attributes {
		if attr, ok := e.(*AttributeLocalVariableTypeTable); ok {
			return attr
		}
	}
	return nil
}

func (a *AttributeCode) StackMapTable() *AttributeStackMapTable {
	for _, e := range a.Attributes {
		if attr, ok := e.(*AttributeStackMapTable); ok {
			return attr
		}
	}
	return nil
}

func (a *AttributeCode) RuntimeVisibleTypeAnnotations() *AttributeRuntimeVisibleTypeAnnotations {
	for _, e := range a.Attributes {
		if attr, ok := e.(*AttributeRuntimeVisibleTypeAnnotations); ok {
			return attr
		}
	}
	return nil
}

func (a *AttributeCode) RuntimeInvisibleTypeAnnotations() *AttributeRuntimeInvisibleTypeAnnotations {
	for _, e := range a.Attributes {
		if attr, ok := e.(*AttributeRuntimeInvisibleTypeAnnotations); ok {
			return attr
		}
	}
	return nil
}

type Exception struct {
	StartPc   uint16
	EndPc     uint16
	HandlerPc uint16
	CatchType uint16
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

type AttributeLocalVariableTable struct {
	LocalVaribleTable []*LocalVariable
}

func (a *AttributeLocalVariableTable) Name() string {
	return "LocalVariableTable"
}

type LocalVariable struct {
	StartPc         uint16
	Length          uint16
	NameIndex       uint16
	DescriptorInedx uint16
	Index           uint16
}

type AttributeLocalVariableTypeTable struct {
	LocalVaribleTypeTable []*LocalVariableType
}

func (a *AttributeLocalVariableTypeTable) Name() string {
	return "LocalVariableTypeTable"
}

type LocalVariableType struct {
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
	elementValue
	TypeIndex         uint16
	ElementValuePairs []*ElementValuePair
}

func (a *Annotation) Type(c *ConstantPool) (string, error) {
	typ, err := c.GetConstantUtf8(a.TypeIndex)
	if err != nil {
		return "", err
	}
	return typ.String(), nil
}

type ElementValuePair struct {
	ElementNameIndex uint16
	ElementValue     ElementValue
}

type AttributeAnnotationDefault struct {
	DefaultValue ElementValue
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

type AttributeRecord struct {
	Components []RecordComponentInfo
}

func (a *AttributeRecord) Name() string {
	return "Record"
}

type RecordComponentInfo struct {
	NameIndex       uint16
	DescriptorIndex uint16
	Attributes      []Attribute
}

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

type AttributeDeprecated struct{}

func (a *AttributeDeprecated) Name() string {
	return "Deprecated"
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

package parser

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

type AttributeModulePackages struct {
	PackageIndexes []uint16
}

func (a *AttributeModulePackages) Name() string {
	return "Module"
}

type AttributeModuleMainClass struct {
	MainClassIndex uint16
}

func (a *AttributeModuleMainClass) Name() string {
	return "ModuleMainClass"
}

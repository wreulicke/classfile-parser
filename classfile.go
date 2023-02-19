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

func (c *Classfile) ThisClassName() (string, error) {
	return c.ConstantPool.GetClassName(c.ThisClass)
}

func (c *Classfile) SuperClassName() (string, error) {
	return c.ConstantPool.GetClassName(c.SuperClass)
}

func (c *Classfile) SourceFile() *AttributeSourceFile {
	for _, e := range c.Attributes {
		if attr, ok := e.(*AttributeSourceFile); ok {
			return attr
		}
	}
	return nil
}

func (c *Classfile) InnerClasses() *AttributeInnerClasses {
	for _, e := range c.Attributes {
		if attr, ok := e.(*AttributeInnerClasses); ok {
			return attr
		}
	}
	return nil
}

func (c *Classfile) EnclosingMethod() *AttributeEnclosingMethod {
	for _, e := range c.Attributes {
		if attr, ok := e.(*AttributeEnclosingMethod); ok {
			return attr
		}
	}
	return nil
}

func (c *Classfile) SourceDebugExtension() *AttributeSourceDebugExtension {
	for _, e := range c.Attributes {
		if attr, ok := e.(*AttributeSourceDebugExtension); ok {
			return attr
		}
	}
	return nil
}

func (c *Classfile) BootstrapMethods() *AttributeBootstrapMethods {
	for _, e := range c.Attributes {
		if attr, ok := e.(*AttributeBootstrapMethods); ok {
			return attr
		}
	}
	return nil
}

func (c *Classfile) Module() *AttributeModule {
	for _, e := range c.Attributes {
		if attr, ok := e.(*AttributeModule); ok {
			return attr
		}
	}
	return nil
}

func (c *Classfile) ModulePackages() *AttributeModulePackages {
	for _, e := range c.Attributes {
		if attr, ok := e.(*AttributeModulePackages); ok {
			return attr
		}
	}
	return nil
}

func (c *Classfile) ModuleMainClass() *AttributeModuleMainClass {
	for _, e := range c.Attributes {
		if attr, ok := e.(*AttributeModuleMainClass); ok {
			return attr
		}
	}
	return nil
}

func (c *Classfile) NestHost() *AttributeNestHost {
	for _, e := range c.Attributes {
		if attr, ok := e.(*AttributeNestHost); ok {
			return attr
		}
	}
	return nil
}

func (c *Classfile) NestMembers() *AttributeNestMembers {
	for _, e := range c.Attributes {
		if attr, ok := e.(*AttributeNestMembers); ok {
			return attr
		}
	}
	return nil
}

func (c *Classfile) Synthetic() *AttributeSynthetic {
	for _, e := range c.Attributes {
		if attr, ok := e.(*AttributeSynthetic); ok {
			return attr
		}
	}
	return nil
}

func (c *Classfile) Deprecated() *AttributeDeprecated {
	for _, e := range c.Attributes {
		if attr, ok := e.(*AttributeDeprecated); ok {
			return attr
		}
	}
	return nil
}

func (c *Classfile) Signature() *AttributeSignature {
	for _, e := range c.Attributes {
		if attr, ok := e.(*AttributeSignature); ok {
			return attr
		}
	}
	return nil
}

func (c *Classfile) RuntimeVisibleAnnotations() *AttributeRuntimeVisibleAnnotations {
	for _, e := range c.Attributes {
		if attr, ok := e.(*AttributeRuntimeVisibleAnnotations); ok {
			return attr
		}
	}
	return nil
}

func (c *Classfile) RuntimeInvisibleAnnotations() *AttributeRuntimeInvisibleAnnotations {
	for _, e := range c.Attributes {
		if attr, ok := e.(*AttributeRuntimeInvisibleAnnotations); ok {
			return attr
		}
	}
	return nil
}

func (c *Classfile) RuntimeVisibleTypeAnnotations() *AttributeRuntimeVisibleTypeAnnotations {
	for _, e := range c.Attributes {
		if attr, ok := e.(*AttributeRuntimeVisibleTypeAnnotations); ok {
			return attr
		}
	}
	return nil
}

func (c *Classfile) RuntimeInvisibleTypeAnnotations() *AttributeRuntimeInvisibleTypeAnnotations {
	for _, e := range c.Attributes {
		if attr, ok := e.(*AttributeRuntimeInvisibleTypeAnnotations); ok {
			return attr
		}
	}
	return nil
}

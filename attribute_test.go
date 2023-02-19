package parser

func findAttribute(attrs []Attribute, name string) Attribute {
	for _, e := range attrs {
		if e.Name() == name {
			return e
		}
	}
	return nil
}

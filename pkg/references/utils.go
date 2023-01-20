package references

import "github.com/upbound/upjet/pkg/config"

type Reference config.Reference
type ReferenceFn interface {
	Get() config.Reference
}

// Get returns the Reference
func (r Reference) Get() config.Reference {
	return config.Reference(r)
}

// WithCustomRefsSelectors Add custom RefFieldName and SelectorFieldName
func (r Reference) WithCustomRefsSelectors(refFieldName, selectorFieldName string) Reference {
	r.RefFieldName = refFieldName
	r.SelectorFieldName = selectorFieldName
	return r
}

// WithCustomExtractor Add custom Extractor
func (r Reference) WithCustomExtractor(extractor string) Reference {
	r.Extractor = extractor
	return r
}

// WithCustomType Add custom Type
func (r Reference) WithCustomType(typeName string) Reference {
	r.Type = typeName
	return r
}

// WithoutResSelector Remove SelectorFieldName and RefFieldName
func (r Reference) WithoutRefsSelectors() Reference {
	r.RefFieldName = ""
	r.SelectorFieldName = ""
	return r
}

// WithoutExtractor Remove Extractor
func (r Reference) WithoutExtractor() Reference {
	r.Extractor = ""
	return r
}

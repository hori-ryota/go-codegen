package typeutil

import "go/types"

//genconstructor
type Printer struct {
	dstPackage *types.Package  `required:""`
	relativeTo types.Qualifier `required:"relativeTo(dstPackage)"`
}

func (p Printer) PrintRelativeType(t types.Type) string {
	return types.TypeString(t, p.relativeTo)
}

func (p Printer) PrintRelativeFuncName(f *types.Func) string {
	relativePkg := p.relativeTo(f.Pkg())
	if relativePkg == "" {
		return f.Name()
	}
	return relativePkg + "." + f.Name()
}

func relativeTo(pkg *types.Package) types.Qualifier {
	return func(other *types.Package) string {
		if SamePackage(pkg, other) {
			return ""
		}
		return other.Name()
	}
}

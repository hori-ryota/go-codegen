package typeutil

import (
	"go/types"
)

//genconstructor
type Printer struct {
	dstPackage    *types.Package    `required:""`
	pkgPathToName map[string]string `setter:""`
}

// implement types.Qualifier
func (p Printer) relativeTo(other *types.Package) string {
	if SamePackage(p.dstPackage, other) {
		return ""
	}
	if name, ok := p.pkgPathToName[other.Path()]; ok {
		return name
	}
	return other.Name()
}

func (p Printer) PrintRelativeType(t types.Type) string {
	return types.TypeString(t, p.relativeTo)
}

func (p Printer) PrintRelativeObject(t types.Object) string {
	return types.ObjectString(t, p.relativeTo)
}

func (p Printer) PrintRelativeConst(c *types.Const) string {
	prefix := p.relativeTo(c.Pkg())
	if prefix != "" {
		prefix += "."
	}
	return prefix + c.Name()
}

func (p Printer) PrintRelativeFuncName(f *types.Func) string {
	relativePkg := p.relativeTo(f.Pkg())
	if relativePkg == "" {
		return f.Name()
	}
	return relativePkg + "." + f.Name()
}

func relativeTo(pkg *types.Package, pkgPathToName map[string]string) types.Qualifier {
	return func(other *types.Package) string {
		if SamePackage(pkg, other) {
			return ""
		}
		if name, ok := pkgPathToName[other.Path()]; ok {
			return name
		}
		return other.Name()
	}
}

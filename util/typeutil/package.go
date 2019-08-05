package typeutil

import (
	"go/types"

	"github.com/hori-ryota/go-codegen/util/gocodeutil"
	"golang.org/x/tools/go/loader"
)

func SamePackage(a *types.Package, b *types.Package) bool {
	if a == b {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return a.Path() == b.Path()
}

func IsContainedInPackage(pkg *types.Package, typ types.Type) bool {
	p := TypeToPackage(typ)
	return SamePackage(pkg, p)
}

func TypeToPackage(typ types.Type) *types.Package {
	if named, ok := typ.(*types.Named); ok {
		return named.Obj().Pkg()
	}

	type pkger interface {
		Pkg() *types.Package
	}

	if t, ok := typ.(pkger); ok {
		return t.Pkg()
	}
	return nil
}

func FmtImports(pkgs []*types.Package) string {
	m := make(map[string]string, len(pkgs))
	for _, pkg := range pkgs {
		if pkg == nil {
			continue
		}
		m[pkg.Name()] = pkg.Path()
	}
	return gocodeutil.FmtImports(m)
}

func AllImported(pkg *loader.PackageInfo) []*types.Package {
	type hasPkg interface {
		Pkg() *types.Package
	}
	dst := make([]*types.Package, 0, 100)
	for _, tv := range pkg.Types {
		if tv.Type == nil {
			continue
		}
		if named, ok := tv.Type.(*types.Named); ok {
			dst = append(dst, named.Obj().Pkg())
			continue
		}
		if t, ok := tv.Type.(hasPkg); ok {
			dst = append(dst, t.Pkg())
			continue
		}
	}
	return dst
}

func SearchPackage(pkg *loader.PackageInfo, name string) *types.Package {
	for _, p := range AllImported(pkg) {
		if p.Name() == name {
			return p
		}
	}
	return nil
}

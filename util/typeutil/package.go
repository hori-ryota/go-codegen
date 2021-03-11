package typeutil

import (
	"go/types"
	"path"
	"strings"

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

func FmtImports(pkgs []*types.Package, current *types.Package) string {
	m := make(map[string]string, len(pkgs))
	for _, pkg := range pkgs {
		if pkg == nil {
			continue
		}
		if current != nil && pkg.Path() == current.Path() {
			continue
		}
		m[pkg.Name()] = pkg.Path()
	}
	if _, ok := m["fmt"]; !ok {
		m["fmt"] = "fmt"
	}
	if _, ok := m["runtime"]; !ok {
		m["runtime"] = "runtime"
	}
	return gocodeutil.FmtImports(m)
}

func AllImported(pkg *loader.PackageInfo) []*types.Package {
	type hasPkg interface {
		Pkg() *types.Package
	}
	dst := make([]*types.Package, 0, 100)
	for _, file := range pkg.Files {
		for _, is := range file.Imports {
			pat := strings.Trim(is.Path.Value, `"`)
			var name string
			if is.Name != nil {
				name = is.Name.Name
			} else {
				name = path.Base(pat)
			}
			dst = append(dst, types.NewPackage(pat, name))
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

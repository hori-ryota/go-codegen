package typeutil

import (
	"go/types"
	"strings"
)

func SearchConstructor(typ *types.Named) *types.Func {
	pkg := typ.Obj().Pkg()
	for _, name := range pkg.Scope().Names() {
		if strings.ToLower(name) != strings.ToLower("New"+typ.Obj().Name()) {
			continue
		}
		obj := pkg.Scope().Lookup(name)
		f, ok := obj.(*types.Func)
		if !ok {
			continue
		}
		return f
	}
	return nil
}

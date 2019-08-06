package typeutil

import "go/types"

func TypeToEnumValues(t types.Type) []*types.Const {
	named, ok := t.(*types.Named)
	if !ok {
		return nil
	}
	pkgScope := named.Obj().Pkg().Scope()
	values := make([]*types.Const, 0, 10)
	for _, n := range pkgScope.Names() {
		obj := pkgScope.Lookup(n)
		cst, ok := obj.(*types.Const)
		if !ok {
			continue
		}
		if named.String() != cst.Type().String() {
			continue
		}
		values = append(values, cst)
	}
	return values
}

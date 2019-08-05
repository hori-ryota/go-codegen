package typeutil

import (
	"go/types"
	"strings"

	"golang.org/x/tools/go/types/typeutil"
)

func SearchSetter(typ *types.Named, paramName string, onlyExported bool) *types.Func {
	for _, method := range typeutil.IntuitiveMethodSet(typ, nil) {
		if onlyExported && !method.Obj().Exported() {
			continue
		}
		if strings.ToLower(method.Obj().Name()) != strings.ToLower("Set"+paramName) {
			continue
		}
		return method.Obj().(*types.Func)
	}
	return nil
}

func ListSetters(typ *types.Named, onlyExported bool) []*types.Func {
	setters := make([]*types.Func, 0, 30)
	for _, method := range typeutil.IntuitiveMethodSet(typ, nil) {
		if onlyExported && !method.Obj().Exported() {
			continue
		}
		if !strings.HasPrefix(strings.ToLower(method.Obj().Name()), "set") {
			continue
		}
		setters = append(setters, method.Obj().(*types.Func))
	}
	return setters
}

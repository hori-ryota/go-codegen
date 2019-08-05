package typeutil

import (
	"fmt"
	"go/types"
	"strings"

	"golang.org/x/tools/go/types/typeutil"
)

func SearchGetter(typ *types.Named, paramName string, onlyExported bool) *types.Func {
	for _, method := range typeutil.IntuitiveMethodSet(typ, nil) {
		if onlyExported && !method.Obj().Exported() {
			continue
		}
		if strings.TrimPrefix(strings.ToLower(method.Obj().Name()), "get") != strings.ToLower(paramName) {
			continue
		}
		return method.Obj().(*types.Func)
	}
	return nil
}

func SearchGetValueStrAndResultType(typ types.Type, paramName string, onlyExported bool) (string, types.Type) {
	if ptr, ok := typ.(*types.Pointer); ok {
		typ = ptr.Elem()
	}
	if named, ok := typ.(*types.Named); ok {
		getter := SearchGetter(named, paramName, onlyExported)
		getterResults := FuncToResults(getter)
		if getter != nil && len(getterResults) > 0 {
			return fmt.Sprintf("%s()", getter.Name()), getterResults[0].Type()
		}
	}
	if strct, ok := typ.Underlying().(*types.Struct); ok {
		field := SearchField(strct, paramName, onlyExported)
		if field != nil {
			return field.Name(), field.Type()
		}
	}
	return "", nil
}

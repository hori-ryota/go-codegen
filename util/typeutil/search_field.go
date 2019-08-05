package typeutil

import (
	"go/types"
	"strings"
)

func SearchField(typ types.Type, paramName string, onlyExported bool) *types.Var {
	if ptr, ok := typ.(*types.Pointer); ok {
		typ = ptr.Elem()
	}
	for _, field := range TypeToFields(typ) {
		if onlyExported && !field.Exported() {
			continue
		}
		if strings.ToLower(field.Name()) != strings.ToLower(paramName) {
			continue
		}
		return field
	}
	return nil
}

package typeutil

import (
	"fmt"
	"go/types"
	"strings"

	"golang.org/x/tools/go/types/typeutil"
)

func (p Printer) PrintGetValue(
	parentName string,
	parentType types.Type,
	targetName string,
) (string, error) {
	for _, field := range TypeToFields(parentType) {
		if !IsContainedInPackage(p.dstPackage, parentType) && !field.Exported() {
			continue
		}
		if strings.ToLower(field.Name()) != strings.ToLower(targetName) {
			continue
		}
		return parentName + "." + field.Name(), nil
	}
	for _, method := range typeutil.IntuitiveMethodSet(parentType, nil) {
		if !IsContainedInPackage(p.dstPackage, parentType) && !method.Obj().Exported() {
			continue
		}
		if strings.ToLower(method.Obj().Name()) != strings.ToLower(targetName) {
			continue
		}
		return parentName + "." + method.Obj().Name() + "()", nil
	}
	return "", fmt.Errorf("failed to resolve getter for '%s' of '%s'", targetName, parentType)
}

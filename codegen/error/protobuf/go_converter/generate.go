package go_converter

import (
	"bytes"
	"go/types"

	"github.com/hori-ryota/go-codegen/util/gocodeutil"
	"github.com/hori-ryota/go-codegen/util/typeutil"
)

func Generate(
	errorDefinitionPkg *types.Package,
	protoPkg *types.Package,
	dstPackage *types.Package,
) (string, error) {
	errorPackagePrefix := ""
	if !typeutil.SamePackage(errorDefinitionPkg, dstPackage) {
		errorPackagePrefix = errorDefinitionPkg.Name() + "."
	}
	protoPackagePrefix := ""
	if !typeutil.SamePackage(protoPkg, dstPackage) {
		protoPackagePrefix = protoPkg.Name() + "."
	}

	out := new(bytes.Buffer)

	err := GoConverterTmpl.Execute(out, TmplParam{
		PackageName:        dstPackage.Name(),
		ImportPackages:     typeutil.FmtImports([]*types.Package{errorDefinitionPkg, protoPkg}, dstPackage),
		ErrorPackagePrefix: errorPackagePrefix,
		ProtoPackagePrefix: protoPackagePrefix,
	})
	if err != nil {
		return "", err
	}

	dst, err := gocodeutil.FormatGoFileFromString(out.String())
	if err != nil {
		return out.String(), err
	}
	return dst, nil
}

type TmplParam struct {
	PackageName        string
	ImportPackages     string
	ErrorPackagePrefix string
	ProtoPackagePrefix string
}

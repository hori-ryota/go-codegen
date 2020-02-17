package go_definition

import (
	"bytes"
	"go/types"

	"github.com/hori-ryota/go-codegen/codegen/error/internal"
	"github.com/hori-ryota/go-codegen/util/gocodeutil"
	"github.com/hori-ryota/go-codegen/util/typeutil"
	"golang.org/x/tools/go/loader"
)

func Generate(pkgInfo *loader.PackageInfo, errorCodes []string, dstPackage *types.Package) (string, error) {
	printer := typeutil.NewPrinter(dstPackage)

	detailErrorCodes, err := internal.ParseDetailErrorCodes(pkgInfo)
	if err != nil {
		return "", err
	}

	param := TmplParam{
		PackageName:      dstPackage.Name(),
		ErrorCodes:       errorCodes,
		DetailErrorCodes: detailErrorCodes,
		ImportPackages:   typeutil.FmtImports(append(typeutil.AllImported(pkgInfo), GodefTmplUsedPackages...), pkgInfo.Pkg),
		TypePrinter:      printer,
	}

	out := new(bytes.Buffer)

	err = GodefTmpl.Execute(out, param)
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
	PackageName      string
	ErrorCodes       []string
	DetailErrorCodes []internal.DetailErrorCodeInfo
	ImportPackages   string
	TypePrinter      typeutil.Printer
}

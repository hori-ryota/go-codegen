package httprpc

import (
	"bytes"
	"fmt"
	"go/types"
	"strings"

	"github.com/hori-ryota/go-codegen/util/gocodeutil"
	"github.com/hori-ryota/go-codegen/util/typeutil"
	"golang.org/x/tools/go/loader"
	xtypeutil "golang.org/x/tools/go/types/typeutil"
)

type Service struct {
	*types.Named
	RPCs []RPC
}

type RPC struct {
	*types.Func
	InputProtoType   *types.Named
	InputClientType  *types.Named
	OutputProtoType  *types.Named
	OutputClientType *types.Named
}

func Generate(
	usecasePkgInfo *loader.PackageInfo,
	protoPkgInfo *loader.PackageInfo,
	clientTypedefPkgInfo *loader.PackageInfo,
	dstPackage *types.Package,
) (string, error) {
	printer := typeutil.NewPrinter(dstPackage)

	services := make([]Service, 0, 100)
	for _, name := range usecasePkgInfo.Pkg.Scope().Names() {
		if !strings.HasSuffix(name, "Usecase") {
			continue
		}
		obj := usecasePkgInfo.Pkg.Scope().Lookup(name)
		named, ok := obj.Type().(*types.Named)
		if !ok {
			continue
		}

		methods := xtypeutil.IntuitiveMethodSet(named, nil)

		rpcs := make([]RPC, 0, len(methods))
		for _, method := range methods {
			rpc, err := ParseRPC(method.Obj().(*types.Func), protoPkgInfo, clientTypedefPkgInfo)
			if err != nil {
				return "", err
			}
			if rpc == nil {
				continue
			}
			rpcs = append(rpcs, *rpc)
		}
		if len(rpcs) > 0 {
			services = append(services, Service{
				Named: named,
				RPCs:  rpcs,
			})
		}
	}
	if len(services) == 0 {
		return "", nil
	}

	out := new(bytes.Buffer)

	importPackages := make([]*types.Package, 0, 100)
	importPackages = append(importPackages, ClientTemplateUsedPackages...)
	importPackages = append(importPackages, typeutil.AllImported(protoPkgInfo)...)
	importPackages = append(importPackages, typeutil.AllImported(clientTypedefPkgInfo)...)

	err := ClientTemplate.Execute(out, TemplateParam{
		PackageName:    dstPackage.Name(),
		ImportPackages: typeutil.FmtImports(importPackages),
		Services:       services,
		TypePrinter:    printer,
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

func ParseRPC(method *types.Func, protoPkgInfo *loader.PackageInfo, clientTypedefPkgInfo *loader.PackageInfo) (*RPC, error) {
	if !method.Exported() {
		return nil, nil
	}
	params := typeutil.FuncToParams(method)
	if len(params) < 2 {
		return nil, nil
	}
	inputType := params[1]
	namedInputType, ok := inputType.Type().(*types.Named)
	if !ok {
		return nil, nil
	}
	if _, ok := namedInputType.Underlying().(*types.Struct); !ok {
		return nil, nil
	}
	inputProtoType := extractPairType(namedInputType, protoPkgInfo)
	if inputProtoType == nil {
		return nil, fmt.Errorf("proto type is not found for '%s'", namedInputType)
	}
	inputClientType := extractPairType(namedInputType, clientTypedefPkgInfo)
	if inputClientType == nil {
		return nil, fmt.Errorf("client type is not found for '%s'", namedInputType)
	}
	rpc := &RPC{
		Func:            method,
		InputProtoType:  inputProtoType,
		InputClientType: inputClientType,
	}
	results := typeutil.FuncToResults(method)
	if len(results) < 2 {
		// without output
		return rpc, nil
	}
	outputType := results[0]
	namedOutputType, ok := outputType.Type().(*types.Named)
	if !ok {
		return nil, nil
	}
	if _, ok := namedOutputType.Underlying().(*types.Struct); !ok {
		return nil, nil
	}
	outputProtoType := extractPairType(namedOutputType, protoPkgInfo)
	if outputProtoType == nil {
		return nil, fmt.Errorf("proto type is not found for '%s'", namedOutputType)
	}
	outputClientType := extractPairType(namedOutputType, clientTypedefPkgInfo)
	if outputClientType == nil {
		return nil, fmt.Errorf("client type is not found for '%s'", namedOutputType)
	}
	rpc.OutputProtoType = outputProtoType
	rpc.OutputClientType = outputClientType
	return rpc, nil
}

func extractPairType(usecaseType *types.Named, pkgInfo *loader.PackageInfo) *types.Named {
	for _, name := range pkgInfo.Pkg.Scope().Names() {
		if strings.ToLower(usecaseType.Obj().Name()) != strings.ToLower(name) {
			continue
		}
		obj := pkgInfo.Pkg.Scope().Lookup(name)
		named, ok := obj.Type().(*types.Named)
		if !ok {
			return nil
		}
		return named
	}
	return nil
}

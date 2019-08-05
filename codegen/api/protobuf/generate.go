package protobuf

import (
	"bytes"
	"go/types"
	"strings"

	"github.com/hori-ryota/go-codegen/util/typeutil"
	"golang.org/x/tools/go/loader"
	xtypeutil "golang.org/x/tools/go/types/typeutil"
)

type Service struct {
	Name string
	RPCs []RPC
}

type RPC struct {
	MethodName string
	InputType  *types.Named
	OutputType *types.Named
}

func Generate(
	pkgInfo *loader.PackageInfo,
	protoPackage string,
	goPackage string,
	javaPackage string,
	javaOuterClassName string,
) (string, error) {
	services := make([]Service, 0, 100)
	for _, name := range pkgInfo.Pkg.Scope().Names() {
		if !strings.HasSuffix(name, "Usecase") {
			continue
		}
		obj := pkgInfo.Pkg.Scope().Lookup(name)
		named, ok := obj.Type().(*types.Named)
		if !ok {
			continue
		}

		methods := xtypeutil.IntuitiveMethodSet(named, nil)

		rpcs := make([]RPC, 0, len(methods))
		for _, method := range methods {
			rpc := ParseRPC(method.Obj().(*types.Func))
			if rpc == nil {
				continue
			}
			rpcs = append(rpcs, *rpc)
		}
		if len(rpcs) > 0 {
			services = append(services, Service{
				Name: name,
				RPCs: rpcs,
			})
		}
	}
	if len(services) == 0 {
		return "", nil
	}

	out := new(bytes.Buffer)

	err := Template.Execute(out, TemplateParam{
		Package:            protoPackage,
		GoPackage:          goPackage,
		JavaPackage:        javaPackage,
		JavaOuterClassName: javaOuterClassName,
		Services:           services,
	})
	if err != nil {
		return "", err
	}

	return out.String(), nil
}

func ParseRPC(method *types.Func) *RPC {
	if !method.Exported() {
		return nil
	}
	params := typeutil.FuncToParams(method)
	if len(params) < 2 {
		return nil
	}
	inputType := params[1]
	namedInputType, ok := inputType.Type().(*types.Named)
	if !ok {
		return nil
	}
	if _, ok := namedInputType.Underlying().(*types.Struct); !ok {
		return nil
	}
	rpc := &RPC{
		MethodName: method.Name(),
		InputType:  namedInputType,
	}
	results := typeutil.FuncToResults(method)
	if len(results) < 2 {
		// without output
		return rpc
	}
	outputType := results[0]
	namedOutputType, ok := outputType.Type().(*types.Named)
	if !ok {
		return nil
	}
	if _, ok := namedOutputType.Underlying().(*types.Struct); !ok {
		return nil
	}
	rpc.OutputType = namedOutputType
	return rpc
}

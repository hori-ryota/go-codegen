package structdef

import (
	"bytes"
	"go/types"
	"strings"

	"github.com/hori-ryota/go-codegen/util/typeutil"
	"golang.org/x/tools/go/loader"
	xtypeutil "golang.org/x/tools/go/types/typeutil"
)

type RPC struct {
	InputType  *types.Named
	OutputType *types.Named
}

func Generate(
	usecasePkgInfo *loader.PackageInfo,
	kotlinPackage string,
) (string, error) {
	typeDefs := make([]*types.Named, 0, 100)
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

		for _, method := range methods {
			rpc, err := ParseRPC(method.Obj().(*types.Func))
			if err != nil {
				return "", err
			}
			if rpc == nil {
				continue
			}
			typeDefs = append(typeDefs, rpc.InputType)
			if rpc.OutputType != nil {
				typeDefs = append(typeDefs, rpc.OutputType)
			}
		}
	}
	if len(typeDefs) == 0 {
		return "", nil
	}

	out := new(bytes.Buffer)

	err := StructDefTemplate.Execute(out, TemplateParam{
		KotlinPackage: kotlinPackage,
		TypeDefs:      typeDefs,
	})
	if err != nil {
		return "", err
	}

	return out.String(), nil
}

func ParseRPC(method *types.Func) (*RPC, error) {
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
	rpc := &RPC{
		InputType: namedInputType,
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
	rpc.OutputType = namedOutputType
	return rpc, nil
}

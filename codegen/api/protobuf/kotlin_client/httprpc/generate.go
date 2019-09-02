package httprpc

import (
	"bytes"
	"go/types"
	"strings"

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
	InputType  *types.Named
	OutputType *types.Named
}

func Generate(
	usecasePkgInfo *loader.PackageInfo,
	kotlinPackage string,
	kotlinClassName string,
	structdefPackage string,
) (string, error) {
	if structdefPackage == "" {
		structdefPackage = kotlinPackage
	}

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
			rpc, err := ParseRPC(method.Obj().(*types.Func))
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

	err := ClientTemplate.Execute(out, TemplateParam{
		KotlinPackage:    kotlinPackage,
		KotlinClassName:  kotlinClassName,
		StructdefPackage: structdefPackage,
		Services:         services,
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
		Func:      method,
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

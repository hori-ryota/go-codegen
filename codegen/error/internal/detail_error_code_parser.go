package internal

import (
	"fmt"
	"go/ast"
	"go/types"
	"strings"

	"github.com/hori-ryota/go-codegen/util/typeutil"
	"golang.org/x/tools/go/loader"
)

const (
	commentMarker = "//errcode"
)

type DetailErrorCodeInfo struct {
	Code   string
	Params []ParamInfo
}

type ParamInfo struct {
	Name string
	Type types.Type
}

func ParseDetailErrorCodes(pkgInfo *loader.PackageInfo) ([]DetailErrorCodeInfo, error) {
	errorDetailComments, err := extractErrorDetailComments(pkgInfo.Files)
	if err != nil {
		return nil, err
	}

	detailErrorCodes := make([]DetailErrorCodeInfo, len(errorDetailComments))
	for i, c := range errorDetailComments {
		c := strings.TrimPrefix(c, commentMarker)
		cs := strings.Split(c, ",")
		info := DetailErrorCodeInfo{
			Code:   strings.TrimSpace(cs[0]),
			Params: make([]ParamInfo, len(cs[1:])),
		}
		for i, param := range cs[1:] {
			ss := strings.Fields(param)
			typeObj := types.Universe.Lookup(ss[1])
			if typeObj != nil {
				info.Params[i] = ParamInfo{
					Name: ss[0],
					Type: typeObj.Type(),
				}
				continue
			}

			paramPkg := pkgInfo.Pkg
			typeName := ss[1]

			splitDotted := strings.SplitN(ss[1], ".", 2)
			if len(splitDotted) > 1 {
				paramPkg = typeutil.SearchPackage(pkgInfo, splitDotted[0])
				typeName = splitDotted[1]
			}
			if paramPkg == nil {
				return nil, fmt.Errorf("unknown param type '%s'", ss[1])
			}
			typObj := paramPkg.Scope().Lookup(typeName)
			if typObj == nil {
				return nil, fmt.Errorf("unknown param type '%s'", ss[1])
			}

			info.Params[i] = ParamInfo{
				Name: ss[0],
				Type: typObj.Type(),
			}
		}
		detailErrorCodes[i] = info
	}
	return detailErrorCodes, nil
}

func extractErrorDetailComments(files []*ast.File) ([]string, error) {
	results := make([]string, 0, 100)
	for _, file := range files {
		for _, cg := range file.Comments {
			for _, c := range cg.List {
				for _, line := range strings.Split(c.Text, "\n") {
					if strings.HasPrefix(line, commentMarker) {
						results = append(results, line)
					}
				}
			}
		}
	}
	return results, nil
}

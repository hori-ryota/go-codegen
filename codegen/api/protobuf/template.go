package protobuf

import (
	"bytes"
	"fmt"
	"go/token"
	"go/types"
	"io"
	"sort"
	"strconv"
	"strings"
	"text/template"

	"github.com/hori-ryota/go-codegen/util/typeutil"
	"github.com/hori-ryota/go-strcase"
)

type TemplateParam struct {
	Package            string
	GoPackage          string
	JavaPackage        string
	JavaOuterClassName string
	Services           []Service
}

var Template = template.Must(template.New("").Funcs(map[string]interface{}{
	"PrintMessageDefs": printMessageDefs,
	"ToImportList": func(rootParam TemplateParam) []string {
		importList := make([]string, 0, 10)
		for _, service := range rootParam.Services {
			for _, rpc := range service.RPCs {
				importList = append(importList, toImportList(rpc.InputType)...)
				if rpc.OutputType == nil {
					importList = append(importList, "google/protobuf/empty.proto")
				} else {
					importList = append(importList, toImportList(rpc.OutputType)...)
				}
			}
		}

		return removeDuplicateAndEmpty(importList)
	},
	"plus": func(a, b int) string { return strconv.Itoa(a + b) },
}).Parse(strings.TrimSpace(`
syntax = "proto3";

package {{.Package}};

{{- if .GoPackage}}
option go_package = "{{.GoPackage}}";
{{- end}}
{{- if .JavaPackage}}
option java_multiple_files = true;
option java_package = "{{.JavaPackage}}";
{{- end}}
{{- if .JavaOuterClassName}}
option java_outer_classname = "{{.JavaOuterClassName}}";
{{- end}}

{{- range ToImportList . }}
import "{{.}}";
{{- end}}
{{- range .Services}}

service {{.Name}} {
{{- range .RPCs}}
  rpc {{.MethodName}}({{.InputType.Obj.Name}}) returns ({{if .OutputType}}{{.OutputType.Obj.Name}}{{else}}google.protobuf.Empty{{end}});
{{- end}}
}
{{- end}}
{{PrintMessageDefs .Services}}
`)))

func removeDuplicateAndEmpty(s []string) []string {
	m := make(map[string]struct{}, len(s))
	for _, k := range s {
		if k == "" {
			continue
		}
		m[k] = struct{}{}
	}
	t := make([]string, 0, len(m))
	for k := range m {
		t = append(t, k)
	}
	sort.Strings(t)
	return t
}

func toImportList(t types.Type) []string {
	if types.TypeString(t, nil) == "time.Time" {
		return []string{"google/protobuf/timestamp.proto"}
	}

	if strct, ok := t.Underlying().(*types.Struct); ok {
		importList := make([]string, 0, 10)
		for _, field := range typeutil.StructToFields(strct) {
			importList = append(importList, toImportList(field.Type())...)
		}
		return importList
	}

	type wrapperType interface {
		Elem() types.Type
	}
	if wrapper, ok := t.(wrapperType); ok {
		return toImportList(wrapper.Elem())
	}
	return nil
}

func printMessageDefs(services []Service) string {
	w := new(bytes.Buffer)
	alreadyDefined := make(map[string]bool, 100)
	for _, s := range services {
		for _, rpc := range s.RPCs {
			printMessageDef(w, rpc.InputType, alreadyDefined)
			if rpc.OutputType != nil {
				printMessageDef(w, rpc.OutputType, alreadyDefined)
			}
		}
	}
	return w.String()
}

func printMessageDef(w io.Writer, def *types.Named, alreadyDefined map[string]bool) {
	name := strcase.ToUpperCamel(def.Obj().Name())
	if alreadyDefined[name] {
		return
	}

	needsMoreDef := make([]*types.Named, 0, 2)
	needsMoreEnumDef := make([]EnumDef, 0, 2)
	internalAlreadyDefined := make(map[string]bool, 2)

	fmt.Fprintf(w, "\nmessage %s {\n", name)

	for i, field := range typeutil.TypeToFields(def) {
		ft := field.Type()
		repeatedMark := ""

	PEELING:
		for {
			switch t := ft.Underlying().(type) {
			case *types.Pointer:
				ft = t.Elem()
			case *types.Slice:
				if t.String() == "[]byte" {
					break PEELING
				}
				ft = t.Elem()
				repeatedMark += "repeated "
			default:
				break PEELING
			}
		}

		var typeName string

		switch ft := ft.(type) {
		case *types.Struct:
			// for anonymous struct
			typeName = strcase.ToUpperCamel(field.Name())

			// define nested message
			b := new(bytes.Buffer)
			printMessageDef(
				b,
				types.NewNamed(types.NewTypeName(token.NoPos, nil, typeName, nil), ft, nil),
				internalAlreadyDefined,
			)
			fmt.Fprint(w, AddIndent(b.String(), 1))
		default:
			typeName = toTypeStr(ft, &needsMoreDef, &needsMoreEnumDef)
		}

		fmt.Fprint(w, AddIndent(
			fmt.Sprintf(
				"%s%s %s = %d;\n",
				repeatedMark,
				typeName,
				strcase.ToLowerSnake(field.Name()),
				i+1,
			),
			1,
		))
	}
	fmt.Fprintln(w, "}")

	alreadyDefined[strcase.ToUpperCamel(def.Obj().Name())] = true

	for _, moreDef := range needsMoreDef {
		printMessageDef(w, moreDef, alreadyDefined)
	}
	for _, enumDef := range needsMoreEnumDef {
		printEnumDef(w, enumDef, alreadyDefined)
	}
}

func printEnumDef(w io.Writer, def EnumDef, alreadyDefined map[string]bool) {
	if alreadyDefined[strcase.ToUpperCamel(def.Type.Obj().Name())] {
		return
	}

	fmt.Fprintf(w, "\nenum %s {\n", strcase.ToUpperCamel(def.Type.Obj().Name()))
	fmt.Fprintf(w, AddIndent(fmt.Sprintf("UNKNOWN_%s = 0;\n", strcase.ToUpperSnake(def.Type.Obj().Name())), 1))
	for i, v := range def.Values {
		if strings.Contains(strings.ToLower(v.Name()), "unknown") {
			continue
		}
		fmt.Fprint(w, AddIndent(
			fmt.Sprintf(
				"%s = %d;\n",
				strcase.ToUpperSnake(v.Name()),
				i+1,
			),
			1,
		))
	}

	fmt.Fprintln(w, "}")

	alreadyDefined[strcase.ToUpperCamel(def.Type.Obj().Name())] = true
}

func AddIndent(s string, indent int) string {
	ss := strings.Split(s, "\n")
	for i := range ss {
		if ss[i] == "" {
			continue
		}
		ss[i] = strings.Repeat("  ", indent) + ss[i]
	}
	return strings.Join(ss, "\n")
}

func KnownTypesToProtoType(t types.Type) (string, bool) {
	underlyingType := t.Underlying().String()
	switch underlyingType {
	case "bool",
		"string",
		"int32", "int64",
		"uint32", "uint64":
		return underlyingType, true
	case "float64":
		return "double", true
	case "float32":
		return "float", true
	case "int8", "int16":
		return "int32", true
	case "int":
		return "int64", true
	case "uint8", "uint16":
		return "uint32", true
	case "uint":
		return "uint64", true
	case "[]byte":
		return "bytes", true
	case "time.Time":
		return "google.protobuf.Timestamp", true
	default:
		return "", false
	}
}

func toTypeStr(
	t types.Type,
	needsMoreDef *[]*types.Named,
	needsMoreEnumDef *[]EnumDef,
) string {
	if named, ok := t.(*types.Named); ok {
		enumValues := typeutil.TypeToEnumValues(named)
		if len(enumValues) > 0 {
			*needsMoreEnumDef = append(*needsMoreEnumDef, EnumDef{
				Type:   named,
				Values: enumValues,
			})
			return strcase.ToUpperCamel(named.Obj().Name())
		}
	}

	if typeStr, ok := KnownTypesToProtoType(t); ok {
		return typeStr
	}

	if named, ok := t.(*types.Named); ok {
		if _, ok := t.Underlying().(*types.Struct); ok {
			*needsMoreDef = append(*needsMoreDef, named)
			return strcase.ToUpperCamel(named.Obj().Name())
		}
	}
	panic(fmt.Errorf("unknown type '%s", t))
}

type EnumDef struct {
	Type   *types.Named
	Values []*types.Const
}

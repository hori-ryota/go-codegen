package protobuf

import (
	"bytes"
	"fmt"
	"go/types"
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
	"PrintMessageDef": func(t *types.Named) string {
		if t == nil {
			return ""
		}
		s, err := printMessageDef(t.Obj().Name(), t.Underlying().(*types.Struct))
		if err != nil {
			panic(err)
		}
		return s
	},
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
  rpc {{.MethodName}} ({{.InputType.Obj.Name}}) returns ({{if .OutputType}}{{.OutputType.Obj.Name}}{{else}}google.protobuf.Empty{{end}});
{{- end}}
}
{{- end}}

{{- range .Services}}
{{- range .RPCs}}

{{PrintMessageDef .InputType}}
{{- if .OutputType}}

{{PrintMessageDef .OutputType}}
{{- end}}
{{- end}}
{{- end}}

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

func printMessageDef(name string, strct *types.Struct) (string, error) {
	w := new(bytes.Buffer)
	fmt.Fprintf(w, "message %s {\n", strcase.ToUpperCamel(name))

	alreadyDefined := make(map[string]bool)

	for i, field := range typeutil.TypeToFields(strct) {
		ft := field.Type()
		repeatedStrIfNeeded := ""
		if slice, ok := ft.Underlying().(*types.Slice); ok {
			repeatedStrIfNeeded = "repeated "
			ft = slice.Elem()
		}
		if ptr, ok := ft.Underlying().(*types.Pointer); ok {
			ft = ptr.Elem()
		}

		if named, ok := ft.(*types.Named); ok {
			enumValues := typeutil.TypeToEnumValues(named)
			if len(enumValues) > 0 {
				enumName := named.Obj().Name()
				if !alreadyDefined[enumName] {
					enumDef := printEnumDef(named, enumValues)
					fmt.Fprintln(w, AddIndent(enumDef, 1))
					alreadyDefined[enumName] = true
				}

				fmt.Fprint(w, AddIndent(
					fmt.Sprintf(
						"%s%s %s = %d;\n",
						repeatedStrIfNeeded,
						strcase.ToUpperCamel(enumName),
						strcase.ToLowerSnake(field.Name()),
						i+1,
					),
					1,
				))
				continue
			}
		}

		if typeStr, ok := KnownTypesToProtoType(ft); ok {
			fmt.Fprint(w, AddIndent(
				fmt.Sprintf(
					"%s%s %s = %d;\n",
					repeatedStrIfNeeded,
					typeStr,
					strcase.ToLowerSnake(field.Name()),
					i+1,
				),
				1,
			))
			continue
		}

		if strct, ok := ft.Underlying().(*types.Struct); ok {
			nestedTypeName := field.Name()
			if named, ok := ft.(*types.Named); ok {
				nestedTypeName = named.Obj().Name()
			}
			if !alreadyDefined[nestedTypeName] {
				nestedMessageDef, err := printMessageDef(nestedTypeName, strct)
				if err != nil {
					return "", err
				}
				fmt.Fprintln(w, AddIndent(nestedMessageDef, 1))
				alreadyDefined[nestedTypeName] = true
			}

			fmt.Fprint(w, AddIndent(
				fmt.Sprintf(
					"%s%s %s = %d;\n",
					repeatedStrIfNeeded,
					strcase.ToUpperCamel(nestedTypeName),
					strcase.ToLowerSnake(field.Name()),
					i+1,
				),
				1,
			))
			continue
		}
		return "", fmt.Errorf("unknown field type '%s", field)
	}
	fmt.Fprint(w, "}")

	return w.String(), nil
}

func printEnumDef(t *types.Named, values []*types.Const) string {
	w := new(bytes.Buffer)
	fmt.Fprintf(w, "enum %s {\n", strcase.ToUpperCamel(t.Obj().Name()))
	fmt.Fprintf(w, AddIndent(fmt.Sprintf("UNKNOWN_%s = 0;\n", strcase.ToUpperSnake(t.Obj().Name())), 1))
	for i, v := range values {
		if strings.ToLower(v.Name()) == "unknown" {
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

	fmt.Fprint(w, "}")

	return w.String()
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
	case "time.Time":
		return "google.protobuf.Timestamp", true
	default:
		return "", false
	}
}

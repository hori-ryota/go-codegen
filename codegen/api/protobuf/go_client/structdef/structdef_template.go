package structdef

import (
	"bytes"
	"fmt"
	"go/types"
	"io"
	"strings"
	"text/template"

	"github.com/hori-ryota/go-codegen/util/typeutil"
	"github.com/hori-ryota/go-strcase"
)

type TemplateParam struct {
	PackageName    string
	ImportPackages string
	TypeDefs       []*types.Named
	TypePrinter    typeutil.Printer
}

var StructDefTemplate = template.Must(template.New("").Funcs(map[string]interface{}{
	"ToLowerCamel": strcase.ToLowerCamel,
	"ToUpperCamel": strcase.ToUpperCamel,
	"PrintDefs":    printDefs,
}).Parse(strings.TrimSpace(`
// Code generated by go-codegen api protobuf go_client httprpc structdef; DO NOT EDIT

package {{.PackageName}}

{{.ImportPackages}}

{{PrintDefs .TypePrinter .TypeDefs}}
`)))

func printDefs(typePrinter typeutil.Printer, typeDefs []*types.Named) string {
	w := new(bytes.Buffer)
	alreadyDefined := make(map[string]bool, 100)
	for _, t := range typeDefs {
		printDef(w, typePrinter, t, alreadyDefined)
	}
	return w.String()
}

func printDef(w io.Writer, typePrinter typeutil.Printer, def *types.Named, alreadyDefined map[string]bool) {
	if alreadyDefined[strcase.ToUpperCamel(def.Obj().Name())] {
		return
	}

	needsMoreDef := make([]*types.Named, 0, 2)
	needsMoreEnumDef := make([]EnumDef, 0, 2)

	fmt.Fprintf(
		w,
		"type %s struct {\n",
		strcase.ToUpperCamel(def.Obj().Name()),
	)

	for _, field := range typeutil.TypeToFields(def) {
		fmt.Fprintf(w, "%s %s\n", strcase.ToUpperCamel(field.Name()), toTypeStr(typePrinter, field.Type(), &needsMoreDef, &needsMoreEnumDef))
	}
	fmt.Fprintln(w, "}")

	fmt.Fprintf(
		w,
		"func New%s(\n",
		strcase.ToUpperCamel(def.Obj().Name()),
	)

	for _, field := range typeutil.TypeToFields(def) {
		fmt.Fprintf(w, "%s %s,\n", strcase.ToLowerCamel(field.Name()), toTypeStr(typePrinter, field.Type(), &needsMoreDef, &needsMoreEnumDef))
	}

	fmt.Fprintf(w, ") %s {\n", strcase.ToUpperCamel(def.Obj().Name()))
	fmt.Fprintf(w, "return %s{\n", strcase.ToUpperCamel(def.Obj().Name()))

	for _, field := range typeutil.TypeToFields(def) {
		fmt.Fprintf(w, "%s:%s,\n", strcase.ToUpperCamel(field.Name()), strcase.ToLowerCamel(field.Name()))
	}

	fmt.Fprintln(w, "}")
	fmt.Fprintln(w, "}")

	alreadyDefined[strcase.ToUpperCamel(def.Obj().Name())] = true

	for _, moreDef := range needsMoreDef {
		printDef(w, typePrinter, moreDef, alreadyDefined)
	}
	for _, enumDef := range needsMoreEnumDef {
		printEnumDef(w, typePrinter, enumDef, alreadyDefined)
	}
}

func printEnumDef(w io.Writer, typePrinter typeutil.Printer, def EnumDef, alreadyDefined map[string]bool) {
	if alreadyDefined[strcase.ToUpperCamel(def.Type.Obj().Name())] {
		return
	}

	fmt.Fprintf(w, "type %s %s\n", def.Type.Obj().Name(), typePrinter.PrintRelativeType(def.Type.Underlying()))
	fmt.Fprintf(w, "const (\n")
	for _, v := range def.Values {
		fmt.Fprintf(w, "%s %s = %s\n", v.Name(), def.Type.Obj().Name(), v.Val().ExactString())
	}
	fmt.Fprintf(w, ")\n")

	alreadyDefined[strcase.ToUpperCamel(def.Type.Obj().Name())] = true
}

func IsKnownNamedType(t *types.Named) bool {
	switch t.String() {
	case "time.Time", "time.Duration":
		return true
	}
	return false
}

var StructDefTemplateUsedPackages = []*types.Package{
	types.NewPackage("time", "time"),
}

func toTypeStr(typePrinter typeutil.Printer, t types.Type, needsMoreDef *[]*types.Named, needsMoreEnumDef *[]EnumDef) string {
	if s, ok := t.Underlying().(*types.Slice); ok {
		return fmt.Sprintf("[]%s", toTypeStr(typePrinter, s.Elem(), needsMoreDef, needsMoreEnumDef))
	}
	if s, ok := t.Underlying().(*types.Pointer); ok {
		return fmt.Sprintf("*%s", toTypeStr(typePrinter, s.Elem(), needsMoreDef, needsMoreEnumDef))
	}
	if named, ok := t.(*types.Named); ok {
		if IsKnownNamedType(named) {
			return typePrinter.PrintRelativeType(t)
		}
		enumValues := typeutil.TypeToEnumValues(named)
		if len(enumValues) > 0 {
			*needsMoreEnumDef = append(*needsMoreEnumDef, EnumDef{
				Type:   named,
				Values: enumValues,
			})
			return named.Obj().Name()
		}

		if _, ok := t.Underlying().(*types.Struct); ok {
			*needsMoreDef = append(*needsMoreDef, named)
			return strcase.ToUpperCamel(named.Obj().Name())
		}
		return typePrinter.PrintRelativeType(named.Underlying())
	}
	return typePrinter.PrintRelativeType(t)
}

type EnumDef struct {
	Type   *types.Named
	Values []*types.Const
}

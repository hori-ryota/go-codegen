package gocodeutil

import (
	"bytes"
	"go/format"
	"go/parser"
	"go/token"
)

func FormatGoFileFromString(s string) (string, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "", s, parser.ParseComments)
	if err != nil {
		return "", err
	}
	RemoveUnusedImports(fset, file)
	b := new(bytes.Buffer)
	if err := format.Node(b, fset, file); err != nil {
		return "", nil
	}
	return b.String(), nil
}

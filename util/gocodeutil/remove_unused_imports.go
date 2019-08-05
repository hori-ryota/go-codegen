package gocodeutil

import (
	"go/ast"
	"go/token"
	"path"
	"strings"

	"golang.org/x/tools/go/ast/astutil"
)

func RemoveUnusedImports(fset *token.FileSet, file *ast.File) {
	unused := make([]*ast.ImportSpec, 0, len(file.Imports))
	for _, pkg := range file.Imports {
		if pkg.Name != nil && (pkg.Name.Name == "." || pkg.Name.Name == "_") {
			continue
		}
		pkgName := path.Base(strings.Trim(pkg.Path.Value, "\""))
		if pkg.Name != nil {
			pkgName = pkg.Name.Name
		}
		if usesImport(file, pkgName) {
			continue
		}
		unused = append(unused, pkg)
	}

	for _, pkg := range unused {
		pkgPath := strings.Trim(pkg.Path.Value, "\"")
		if pkg.Name == nil {
			astutil.DeleteImport(fset, file, pkgPath)
			continue
		}
		astutil.DeleteNamedImport(fset, file, pkg.Name.Name, pkgPath)
	}
}

func usesImport(f *ast.File, pkgName string) bool {
	used := false
	ast.Inspect(f, func(n ast.Node) bool {
		if n == nil {
			return true
		}
		sel, ok := n.(*ast.SelectorExpr)
		if !ok {
			return true
		}
		id, ok := sel.X.(*ast.Ident)
		if !ok {
			return true
		}
		if id.Obj != nil {
			// is not top name
			return true
		}
		if id.Name != pkgName {
			return true
		}

		used = true
		return false
	})
	return used
}

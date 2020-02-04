package prototoolyaml

import (
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/hori-ryota/go-codegen/util/gocodeutil"
)

type TemplateParam struct {
	OutputDir     string
	JavaOutputDir string
}

var Template = template.Must(template.New("").Funcs(template.FuncMap{
	"LocalPathToGoPackagePath": func(s string) string {
		p, err := gocodeutil.LocalPathToGoPackagePath(s)
		if err != nil {
			panic(err)
		}
		return p
	},
	"ToGoPath": func(s string) string {
		p, err := filepath.Abs(filepath.Dir(filepath.FromSlash(s)))
		if err != nil {
			panic(err)
		}
		gopath := os.Getenv("GOPATH")
		if gopath == "" {
			panic("GOPATH is empty")
		}
		toGoPath, err := filepath.Rel(p, gopath)
		if err != nil {
			panic(err)
		}
		return filepath.ToSlash(toGoPath)
	},
	"Rel": func(base, t string) string {
		rel, err := filepath.Rel(base, t)
		if err != nil {
			panic(err)
		}
		return filepath.ToSlash(rel)
	},
}).Parse(strings.TrimSpace(`
lint:
  group: empty

generate:
  go_options:
    import_path: {{LocalPathToGoPackagePath .OutputDir}}

  plugins:
    - name: go
      output: {{ToGoPath .OutputDir}} # GOPATH/src
	{{- if .JavaOutputDir}}
    - name: java
      output: {{Rel .OutputDir .JavaOutputDir}}
	{{- end}}
`)))

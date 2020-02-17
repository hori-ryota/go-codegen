package protobuf

import (
	"strings"
	"text/template"
)

type TemplateParam struct {
	Package            string
	GoPackage          string
	JavaPackage        string
	JavaOuterClassName string
}

var Template = template.Must(template.New("").Parse(strings.TrimSpace(`
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

message Error {
  string code = 1;
  repeated ErrorDetail details = 2;
}

message ErrorDetail {
  string code = 1;
  repeated string args = 2;
}
`)))

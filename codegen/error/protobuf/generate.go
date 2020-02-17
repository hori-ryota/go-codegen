package protobuf

import (
	"bytes"

	"golang.org/x/tools/go/loader"
)

func Generate(
	pkgInfo *loader.PackageInfo,
	protoPackage string,
	goPackage string,
	javaPackage string,
	javaOuterClassName string,
) (string, error) {
	out := new(bytes.Buffer)

	err := Template.Execute(out, TemplateParam{
		Package:            protoPackage,
		GoPackage:          goPackage,
		JavaPackage:        javaPackage,
		JavaOuterClassName: javaOuterClassName,
	})
	if err != nil {
		return "", err
	}

	return out.String(), nil
}

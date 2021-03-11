package bufgenyaml

import (
	"bytes"
)

func Generate(
	outputDir string,
	javaOutputDir string,
) (string, error) {
	out := new(bytes.Buffer)

	err := Template.Execute(out, TemplateParam{
		OutputDir:     outputDir,
		JavaOutputDir: javaOutputDir,
	})
	if err != nil {
		return "", err
	}

	return out.String(), nil
}

package prototoolyaml_test

import (
	cmd "github.com/hori-ryota/go-codegen/codegen"
	"github.com/mattn/go-shellwords"
	"github.com/spf13/cobra"
)

func Example() {
	input := "api protobuf prototoolyaml -o ../../internal/testdata/external/adapter/protobuf --javaOutputDir ../../internal/testdata/external/adapter/protobuf/java/src/main/java"

	cmd := prepareCmdWithStdout(input)
	err := cmd.Execute()
	if err != nil {
		panic(err)
	}
	// Output:
	// lint:
	//   group: empty
	//
	// generate:
	//   go_options:
	//     import_path: github.com/hori-ryota/go-codegen/codegen/api/internal/testdata/external/adapter/protobuf
	//
	//   plugins:
	//     - name: go
	//       output: ../../../../../../../../../.. # GOPATH
	//     - name: java
	//       output: java/src/main/java
}

func prepareCmdWithStdout(input string) *cobra.Command {
	input += " --useStdOut"
	args, err := shellwords.Parse(input)
	if err != nil {
		panic(err)
	}
	cmd := cmd.NewRootCmd()
	cmd.SetArgs(args)
	return cmd
}

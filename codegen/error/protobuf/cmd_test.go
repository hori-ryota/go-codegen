package protobuf_test

import (
	cmd "github.com/hori-ryota/go-codegen/codegen"
	"github.com/mattn/go-shellwords"
	"github.com/spf13/cobra"
)

func Example() {
	input := "error protobuf -t ../internal/testdata -o ../internal/testdata/error.proto --protoPackage testdata --javaPackage com.github.horiryota.gocodegen.error.testdata --javaOuterClassName ErrorProto"

	cmd := prepareCmdWithStdout(input)
	err := cmd.Execute()
	if err != nil {
		panic(err)
	}
	// Output:
	// syntax = "proto3";
	//
	// package testdata;
	// option go_package = "github.com/hori-ryota/go-codegen/codegen/error/internal/testdata";
	// option java_multiple_files = true;
	// option java_package = "com.github.horiryota.gocodegen.error.testdata";
	// option java_outer_classname = "ErrorProto";
	//
	// message Error {
	//   string code = 1;
	//   repeated ErrorDetail details = 2;
	// }
	//
	// message ErrorDetail {
	//   string code = 1;
	//   repeated string args = 2;
	// }
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

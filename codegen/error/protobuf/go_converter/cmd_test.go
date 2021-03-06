package go_converter_test

import (
	cmd "github.com/hori-ryota/go-codegen/codegen"
	"github.com/mattn/go-shellwords"
	"github.com/spf13/cobra"
)

func Example() {
	input := "error protobuf go_converter -t ../../internal/testdata -p ../../internal/testdata/proto -o ../internal/testdata/converter"

	cmd := prepareCmdWithStdout(input)
	err := cmd.Execute()
	if err != nil {
		panic(err)
	}
	// Output:
	// // Code generated by go-codegen error protobuf go_converter; DO NOT EDIT
	//
	// package converter
	//
	// import (
	// 	"github.com/hori-ryota/go-codegen/codegen/error/internal/testdata"
	// 	"github.com/hori-ryota/go-codegen/codegen/error/internal/testdata/proto"
	// )
	//
	// func ConvertErrorToProto(derr testdata.Error) *proto.Error {
	// 	details := make([]*proto.ErrorDetail, len(derr.Details()))
	// 	for i := range derr.Details() {
	// 		details[i] = &proto.ErrorDetail{
	// 			Code: derr.Details()[i].Code().String(),
	// 			Args: derr.Details()[i].Args(),
	// 		}
	// 	}
	// 	return &proto.Error{
	// 		Code:    derr.Code().String(),
	// 		Details: details,
	// 	}
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

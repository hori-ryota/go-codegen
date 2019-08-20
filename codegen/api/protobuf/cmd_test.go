package protobuf_test

import (
	cmd "github.com/hori-ryota/go-codegen/codegen"
	"github.com/mattn/go-shellwords"
	"github.com/spf13/cobra"
)

func Example() {
	input := "api protobuf -t ../internal/testdata/internal/application -o ../internal/testdata/external/adapter/protobuf/example.proto --protoPackage codegen --javaPackage com.github.hori-ryota/go-codegen/api/example/codegen --javaOuterClassName CodegenProto"

	cmd := prepareCmdWithStdout(input)
	err := cmd.Execute()
	if err != nil {
		panic(err)
	}
	// Output:
	// syntax = "proto3";
	// package codegen;
	// option go_package = "github.com/hori-ryota/go-codegen/codegen/api/internal/testdata/external/adapter/protobuf";
	// option java_multiple_files = true;
	// option java_package = "com.github.hori-ryota/go-codegen/api/example/codegen";
	// option java_outer_classname = "CodegenProto";
	// import "google/protobuf/empty.proto";
	//
	// service DoingSomethingWithOutputAndActorUsecase {
	//   rpc DoSomethingWithOutputAndActor (DoingSomethingWithOutputAndActorUsecaseInput) returns (DoingSomethingWithOutputAndActorUsecaseOutput);
	// }
	//
	// service DoingSomethingWithOutputWithoutActorUsecase {
	//   rpc DoSomethingWithOutputWithoutActor (DoingSomethingWithOutputWithoutActorUsecaseInput) returns (DoingSomethingWithOutputWithoutActorUsecaseOutput);
	// }
	//
	// service DoingSomethingWithoutOutputAndActorUsecase {
	//   rpc DoSomethingWithoutOutputAndActor (DoingSomethingWithoutOutputAndActorUsecaseInput) returns (google.protobuf.Empty);
	// }
	//
	// service DoingSomethingWithoutOutputWithActorUsecase {
	//   rpc DoSomethingWithoutOutputWithActor (DoingSomethingWithoutOutputWithActorUsecaseInput) returns (google.protobuf.Empty);
	// }
	//
	// message DoingSomethingWithOutputAndActorUsecaseInput {
	//   string string_param = 1;
	//   int64 int_param = 2;
	//   int64 int64_param = 3;
	//   uint64 uint_param = 4;
	//   uint64 uint64_param = 5;
	//   float float32_param = 6;
	//   double float64_param = 7;
	//   bytes bytes_param = 8;
	//   message AnonymousNestedStructParam {
	//     string string_param = 1;
	//   }
	//   AnonymousNestedStructParam anonymous_nested_struct_param = 9;
	//   message NamedSomeType {
	//     string string_param = 1;
	//   }
	//   NamedSomeType named_nested_struct_param = 10;
	//   enum StringEnum {
	//     UNKNOWN_STRING_ENUM = 0;
	//     STRING_A = 1;
	//     STRING_B = 2;
	//     STRING_C = 3;
	//   }
	//   StringEnum string_enum_param = 11;
	//   enum IntEnum {
	//     UNKNOWN_INT_ENUM = 0;
	//     INT_ONE = 1;
	//     INT_THREE = 2;
	//     INT_TWO = 3;
	//   }
	//   IntEnum int_enum_param = 12;
	//   repeated string string_slice_param = 13;
	//   repeated int64 int_slice_param = 14;
	//   repeated int64 int64_slice_param = 15;
	//   repeated uint64 uint_slice_param = 16;
	//   repeated uint64 uint64_slice_param = 17;
	//   repeated float float32_slice_param = 18;
	//   repeated double float64_slice_param = 19;
	//   repeated bytes bytes_slice_param = 20;
	//   message AnonymousNestedStructSliceParam {
	//     string string_param = 1;
	//   }
	//   repeated AnonymousNestedStructSliceParam anonymous_nested_struct_slice_param = 21;
	//   repeated NamedSomeType named_nested_struct_slice_param = 22;
	//   repeated StringEnum string_enum_slice_param = 23;
	//   repeated IntEnum int_enum_slice_param = 24;
	// }
	//
	// message DoingSomethingWithOutputAndActorUsecaseOutput {
	//   string string_param = 1;
	// }
	//
	// message DoingSomethingWithOutputWithoutActorUsecaseInput {
	//   string string_param = 1;
	// }
	//
	// message DoingSomethingWithOutputWithoutActorUsecaseOutput {
	//   string string_param = 1;
	// }
	//
	// message DoingSomethingWithoutOutputAndActorUsecaseInput {
	//   string string_param = 1;
	// }
	//
	// message DoingSomethingWithoutOutputWithActorUsecaseInput {
	//   string string_param = 1;
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

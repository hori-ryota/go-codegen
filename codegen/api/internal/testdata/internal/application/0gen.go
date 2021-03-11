package application

// accessor
//go:generate go-codegen go_accessor
// constructor
//go:generate go-codegen go_constructor

// api

//// protobuf
//go:generate go-codegen api protobuf -o ../../external/adapter/protobuf/codegen.proto --protoPackage codegen --javaPackage com.github.horiryota.gocodegen.api.example.codegen --javaOuterClassName CodegenProto
//// protobuf of error
//go:generate go-codegen error protobuf -o ../../external/adapter/protobuf/error.proto --protoPackage codegen --javaPackage com.github.horiryota.gocodegen.api.example.codegen --javaOuterClassName ErrorProto
//// buf.gen.yaml
//go:generate go-codegen api protobuf bufgenyaml -o ../../external/adapter/protobuf
//// buf generate
//go:generate buf generate --template ../../external/adapter/protobuf/buf.gen.yaml -o ../../external/adapter/protobuf ../../external/adapter/protobuf

//// httprpc server of Go
//go:generate go-codegen api protobuf go_server httprpc -p ../../external/adapter/protobuf -o ../adapter/protobuf/httprpc
//go:generate go-codegen error protobuf go_converter -t ../domain -p ../../external/adapter/protobuf -o ../adapter/protobuf/httprpc

//// httprpc client of Go

////// structdef
//go:generate go-codegen api protobuf go_client structdef -o ../../external/adapter/protobuf/httprpc
////// client of go
//go:generate go-codegen api protobuf go_client httprpc -p ../../external/adapter/protobuf -c ../../external/adapter/protobuf/httprpc -o ../../external/adapter/protobuf/httprpc

//// httprpc client of Kotlin multiplatform
////// structdef
//go:generate go-codegen api protobuf kotlin_client structdef -o ../../external/adapter/protobuf/kotlinmpp/src/commonMain/kotlin --kotlinPackage com.github.horiryota.gocodegen.api.example.codegen
////// client of kotlin multiplatform
//go:generate go-codegen api protobuf kotlin_client httprpc -o ../../external/adapter/protobuf/kotlinmpp/src/commonMain/kotlin --kotlinPackage com.github.horiryota.gocodegen.api.example.codegen --kotlinClassName CodegenExampleApi

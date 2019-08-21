package application

// accessor
//go:generate go run ../../../../../../ go_accessor
// constructor
//go:generate go run ../../../../../../ go_constructor

// api

//// protobuf
//go:generate go run ../../../../../../ api protobuf -o ../../external/adapter/protobuf/codegen.proto --protoPackage codegen --javaPackage com.github.horiryota.gocodegen.api.example.codegen --javaOuterClassName CodegenProto
//// prototool.yaml
//go:generate go run ../../../../../../ api protobuf prototoolyaml -o ../../external/adapter/protobuf --javaOutputDir ../../external/adapter/protobuf/java/src/main/java
//// prototool generate
//go:generate prototool generate ../../external/adapter/protobuf

//// httprpc server of Go
//go:generate go run ../../../../../../ api protobuf go_server httprpc -p ../../external/adapter/protobuf -o ../adapter/protobuf/httprpc

//// httprpc client of Go

////// structdef
//go:generate go run ../../../../../../ api protobuf go_client structdef -o ../../external/adapter/protobuf/httprpc
////// client of go
//go:generate go run ../../../../../../ api protobuf go_client httprpc -p ../../external/adapter/protobuf -c ../../external/adapter/protobuf/httprpc -o ../../external/adapter/protobuf/httprpc

//// httprpc client of Kotlin
////// structdef
//go:generate go run ../../../../../../ api protobuf kotlin_client structdef -o ../../external/adapter/protobuf/kotlin/src/commonMain/kotlin --kotlinPackage com.github.horiryota.gocodegen.api.example.codegen

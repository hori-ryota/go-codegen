package testdata

import "time"

//go:generate go-codegen error go_definition -t . -o . --codes Unknown,NotFound,BadRequest
//go:generate go-codegen error protobuf -t . -o ./proto/error.proto --protoPackage testdata --javaPackage com.github.horiryota.gocodegen.error.testdata --javaOuterClassName ErrorProto
//go:generate protoc --go_out=$GOPATH/src ./proto/error.proto

type WorkKind string

type WorkingTime struct {
	WorkKind WorkKind
	StartAt  time.Time
	EndAt    time.Time
}

type WorkingTimeSpec struct {
	allowedWorkKinds []WorkKind
}

func (s WorkingTimeSpec) Validate(workingTime WorkingTime) Error {
	if s.IsAllowedWorkKind(workingTime.WorkKind) {
		//errcode InvalidWorkKind,workKind WorkKind
		return nil // TODO
	}

	//errcode InvalidWorkingTime,startAt int64,endAt int64
	return nil
}

func (s WorkingTimeSpec) IsAllowedWorkKind(workKind WorkKind) bool {
	for _, w := range s.allowedWorkKinds {
		if w == workKind {
			return true
		}
	}
	return false
}

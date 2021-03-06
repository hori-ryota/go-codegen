package go_definition_test

import (
	cmd "github.com/hori-ryota/go-codegen/codegen"
	"github.com/mattn/go-shellwords"
	"github.com/spf13/cobra"
)

func Example() {
	input := "error go_definition -t ../internal/testdata -o ../internal/testdata --codes Unknown,NotFound,BadRequest"

	cmd := prepareCmdWithStdout(input)
	err := cmd.Execute()
	if err != nil {
		panic(err)
	}
	// Output:
	// // Code generated by go-codegen error go_definition; DO NOT EDIT
	//
	// package testdata
	//
	// import (
	// 	"fmt"
	// 	"strconv"
	// 	"strings"
	//
	// 	"github.com/hori-ryota/zaperr"
	// 	"go.uber.org/zap"
	// 	"go.uber.org/zap/zapcore"
	// )
	//
	// type ErrorCode string
	//
	// const (
	// 	errorUnknown    ErrorCode = "Unknown"
	// 	errorNotFound   ErrorCode = "NotFound"
	// 	errorBadRequest ErrorCode = "BadRequest"
	// )
	//
	// func (c ErrorCode) String() string {
	// 	return string(c)
	// }
	//
	// type Error interface {
	// 	Error() string
	// 	Code() ErrorCode
	// 	Details() []ErrorDetail
	//
	// 	IsUnknown() bool
	// 	IsNotFound() bool
	// 	IsBadRequest() bool
	// }
	//
	// func newError(source error, code ErrorCode, details ...ErrorDetail) Error {
	// 	return errorImpl{
	// 		source:  source,
	// 		code:    code,
	// 		details: details,
	// 	}
	// }
	//
	// func ErrorUnknown(source error, details ...ErrorDetail) Error {
	// 	return newError(source, errorUnknown, details...)
	// }
	// func ErrorNotFound(source error, details ...ErrorDetail) Error {
	// 	return newError(source, errorNotFound, details...)
	// }
	// func ErrorBadRequest(source error, details ...ErrorDetail) Error {
	// 	return newError(source, errorBadRequest, details...)
	// }
	//
	// type errorImpl struct {
	// 	source  error
	// 	code    ErrorCode
	// 	details []ErrorDetail
	// }
	//
	// func (e errorImpl) Error() string {
	// 	return fmt.Sprintf("%s:%s:%s", e.code, e.details, e.source)
	// }
	// func (e errorImpl) Code() ErrorCode {
	// 	return e.code
	// }
	// func (e errorImpl) Details() []ErrorDetail {
	// 	return e.details
	// }
	//
	// func (e errorImpl) IsUnknown() bool {
	// 	return e.code == errorUnknown
	// }
	// func (e errorImpl) IsNotFound() bool {
	// 	return e.code == errorNotFound
	// }
	// func (e errorImpl) IsBadRequest() bool {
	// 	return e.code == errorBadRequest
	// }
	//
	// type ErrorDetail struct {
	// 	code ErrorDetailCode
	// 	args []string
	// }
	//
	// func newErrorDetail(code ErrorDetailCode, args ...string) ErrorDetail {
	// 	return ErrorDetail{
	// 		code: code,
	// 		args: args,
	// 	}
	// }
	//
	// func (e ErrorDetail) String() string {
	// 	return strings.Join(append([]string{e.code.String()}, e.args...), ",")
	// }
	//
	// func (c ErrorDetail) Code() ErrorDetailCode {
	// 	return c.code
	// }
	//
	// func (c ErrorDetail) Args() []string {
	// 	return c.args
	// }
	//
	// type ErrorDetailCode string
	//
	// func (c ErrorDetailCode) String() string {
	// 	return string(c)
	// }
	//
	// const ErrorDetailInvalidWorkKind ErrorDetailCode = "InvalidWorkKind"
	//
	// func InvalidWorkKindError(
	// 	workKind WorkKind,
	// ) ErrorDetail {
	// 	return newErrorDetail(
	// 		ErrorDetailInvalidWorkKind,
	// 		string(workKind),
	// 	)
	// }
	//
	// const ErrorDetailInvalidWorkingTime ErrorDetailCode = "InvalidWorkingTime"
	//
	// func InvalidWorkingTimeError(
	// 	startAt int64,
	// 	endAt int64,
	// ) ErrorDetail {
	// 	return newErrorDetail(
	// 		ErrorDetailInvalidWorkingTime,
	// 		strconv.FormatInt(startAt, 10),
	// 		strconv.FormatInt(endAt, 10),
	// 	)
	// }
	//
	// func (e errorImpl) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	// 	zaperr.ToNamedField("sourceError", e.source).AddTo(enc)
	// 	zap.String("code", string(e.code)).AddTo(enc)
	// 	zap.Any("details", e.details)
	// 	return nil
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

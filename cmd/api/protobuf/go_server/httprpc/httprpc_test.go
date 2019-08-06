package httprpc_test

import (
	"github.com/hori-ryota/go-codegen/cmd"
	"github.com/mattn/go-shellwords"
	"github.com/spf13/cobra"
)

func Example() {
	input := "api protobuf go_server httprpc -t ../../../internal/testdata/internal/application -p ../../../internal/testdata/external/adapter/protobuf -f ../../../internal/testdata/internal/adapter/usecasefactory -o ../../../internal/testdata/internal/adapter/protobuf/httprpc"

	cmd := prepareCmdWithStdout(input)
	err := cmd.Execute()
	if err != nil {
		panic(err)
	}
	// Output:
	// // Code generated by go-codegen api protobuf go_server httprpc; DO NOT EDIT
	//
	// package httprpc
	//
	// import (
	// 	ioutil "io/ioutil"
	// 	http "net/http"
	//
	// 	proto "github.com/golang/protobuf/proto"
	// 	protobuf "github.com/hori-ryota/go-codegen/cmd/api/internal/testdata/external/adapter/protobuf"
	// 	usecasefactory "github.com/hori-ryota/go-codegen/cmd/api/internal/testdata/internal/adapter/usecasefactory"
	// 	application "github.com/hori-ryota/go-codegen/cmd/api/internal/testdata/internal/application"
	// 	domain "github.com/hori-ryota/go-codegen/cmd/api/internal/testdata/internal/domain"
	// )
	//
	// func NewHandlers(
	// 	handleError func(w http.ResponseWriter, r *http.Request, err error),
	// 	doingSomethingWithOutputAndActorUsecaseFactory usecasefactory.DoingSomethingWithOutputAndActorUsecaseFactory,
	// 	doingSomethingWithOutputWithoutActorUsecaseFactory usecasefactory.DoingSomethingWithOutputWithoutActorUsecaseFactory,
	// 	doingSomethingWithoutOutputAndActorUsecaseFactory usecasefactory.DoingSomethingWithoutOutputAndActorUsecaseFactory,
	// 	doingSomethingWithoutOutputWithActorUsecaseFactory usecasefactory.DoingSomethingWithoutOutputWithActorUsecaseFactory,
	// ) Handlers {
	// 	return Handlers{
	// 		HandleError: handleError,
	// 		DoingSomethingWithOutputAndActorUsecaseFactory:     doingSomethingWithOutputAndActorUsecaseFactory,
	// 		DoingSomethingWithOutputWithoutActorUsecaseFactory: doingSomethingWithOutputWithoutActorUsecaseFactory,
	// 		DoingSomethingWithoutOutputAndActorUsecaseFactory:  doingSomethingWithoutOutputAndActorUsecaseFactory,
	// 		DoingSomethingWithoutOutputWithActorUsecaseFactory: doingSomethingWithoutOutputWithActorUsecaseFactory,
	// 	}
	// }
	//
	// type Handlers struct {
	// 	HandleError                                        func(w http.ResponseWriter, r *http.Request, err error)
	// 	DoingSomethingWithOutputAndActorUsecaseFactory     usecasefactory.DoingSomethingWithOutputAndActorUsecaseFactory
	// 	DoingSomethingWithOutputWithoutActorUsecaseFactory usecasefactory.DoingSomethingWithOutputWithoutActorUsecaseFactory
	// 	DoingSomethingWithoutOutputAndActorUsecaseFactory  usecasefactory.DoingSomethingWithoutOutputAndActorUsecaseFactory
	// 	DoingSomethingWithoutOutputWithActorUsecaseFactory usecasefactory.DoingSomethingWithoutOutputWithActorUsecaseFactory
	// }
	//
	// func (h Handlers) DoingSomethingWithOutputAndActorUsecaseDoSomethingWithOutputAndActorHandler(w http.ResponseWriter, r *http.Request) {
	// 	body, err := ioutil.ReadAll(r.Body)
	// 	if err != nil {
	// 		h.HandleError(w, r, err)
	// 		return
	// 	}
	//
	// 	inputProtoType := protobuf.DoingSomethingWithOutputAndActorUsecaseInput{}
	// 	if err := proto.Unmarshal(body, &inputProtoType); err != nil {
	// 		h.HandleError(w, r, err)
	// 		return
	// 	}
	// 	input := func() application.DoingSomethingWithOutputAndActorUsecaseInput {
	// 		m := application.NewDoingSomethingWithOutputAndActorUsecaseInput(
	// 			inputProtoType.GetStringParam(),
	// 			int(inputProtoType.GetIntParam()),
	// 			inputProtoType.GetInt64Param(),
	// 			uint(inputProtoType.GetUintParam()),
	// 			inputProtoType.GetUint64Param(),
	// 			inputProtoType.GetFloat32Param(),
	// 			inputProtoType.GetFloat64Param(),
	// 			struct{ StringParam string }{
	// 				StringParam: inputProtoType.GetAnonymousNestedStructParam().GetStringParam(),
	// 			},
	// 			func() application.NamedSomeType {
	// 				m := application.NewNamedSomeType(
	// 					inputProtoType.GetNamedNestedStructParam().GetStringParam(),
	// 				)
	//
	// 				return m
	// 			}(),
	// 			func(s protobuf.DoingSomethingWithOutputAndActorUsecaseInput_StringEnum) domain.StringEnum {
	// 				switch s {
	// 				case protobuf.DoingSomethingWithOutputAndActorUsecaseInput_STRING_A:
	// 					return domain.StringA
	// 				case protobuf.DoingSomethingWithOutputAndActorUsecaseInput_STRING_B:
	// 					return domain.StringB
	// 				case protobuf.DoingSomethingWithOutputAndActorUsecaseInput_STRING_C:
	// 					return domain.StringC
	//
	// 				default:
	// 					var t domain.StringEnum
	// 					return t
	// 				}
	// 			}(inputProtoType.GetStringEnumParam()),
	// 			func(s protobuf.DoingSomethingWithOutputAndActorUsecaseInput_IntEnum) domain.IntEnum {
	// 				switch s {
	// 				case protobuf.DoingSomethingWithOutputAndActorUsecaseInput_INT_ONE:
	// 					return domain.IntOne
	// 				case protobuf.DoingSomethingWithOutputAndActorUsecaseInput_INT_THREE:
	// 					return domain.IntThree
	// 				case protobuf.DoingSomethingWithOutputAndActorUsecaseInput_INT_TWO:
	// 					return domain.IntTwo
	//
	// 				default:
	// 					var t domain.IntEnum
	// 					return t
	// 				}
	// 			}(inputProtoType.GetIntEnumParam()),
	// 			inputProtoType.GetStringSliceParam(),
	// 			func() []int {
	// 				t := make([]int, len(inputProtoType.GetIntSliceParam()))
	// 				for i := range t {
	// 					t[i] = int(inputProtoType.GetIntSliceParam()[i])
	// 				}
	// 				return t
	// 			}(),
	// 			inputProtoType.GetInt64SliceParam(),
	// 			func() []uint {
	// 				t := make([]uint, len(inputProtoType.GetUintSliceParam()))
	// 				for i := range t {
	// 					t[i] = uint(inputProtoType.GetUintSliceParam()[i])
	// 				}
	// 				return t
	// 			}(),
	// 			inputProtoType.GetUint64SliceParam(),
	// 			inputProtoType.GetFloat32SliceParam(),
	// 			inputProtoType.GetFloat64SliceParam(),
	// 			func() []struct{ StringParam string } {
	// 				t := make([]struct{ StringParam string }, len(inputProtoType.GetAnonymousNestedStructSliceParam()))
	// 				for i := range t {
	// 					t[i] = struct{ StringParam string }{
	// 						StringParam: inputProtoType.GetAnonymousNestedStructSliceParam()[i].GetStringParam(),
	// 					}
	// 				}
	// 				return t
	// 			}(),
	// 			func() []application.NamedSomeType {
	// 				t := make([]application.NamedSomeType, len(inputProtoType.GetNamedNestedStructSliceParam()))
	// 				for i := range t {
	// 					t[i] = func() application.NamedSomeType {
	// 						m := application.NewNamedSomeType(
	// 							inputProtoType.GetNamedNestedStructSliceParam()[i].GetStringParam(),
	// 						)
	//
	// 						return m
	// 					}()
	// 				}
	// 				return t
	// 			}(),
	// 			func() []domain.StringEnum {
	// 				t := make([]domain.StringEnum, len(inputProtoType.GetStringEnumSliceParam()))
	// 				for i := range t {
	// 					t[i] = func(s protobuf.DoingSomethingWithOutputAndActorUsecaseInput_StringEnum) domain.StringEnum {
	// 						switch s {
	// 						case protobuf.DoingSomethingWithOutputAndActorUsecaseInput_STRING_A:
	// 							return domain.StringA
	// 						case protobuf.DoingSomethingWithOutputAndActorUsecaseInput_STRING_B:
	// 							return domain.StringB
	// 						case protobuf.DoingSomethingWithOutputAndActorUsecaseInput_STRING_C:
	// 							return domain.StringC
	//
	// 						default:
	// 							var t domain.StringEnum
	// 							return t
	// 						}
	// 					}(inputProtoType.GetStringEnumSliceParam()[i])
	// 				}
	// 				return t
	// 			}(),
	// 			func() []domain.IntEnum {
	// 				t := make([]domain.IntEnum, len(inputProtoType.GetIntEnumSliceParam()))
	// 				for i := range t {
	// 					t[i] = func(s protobuf.DoingSomethingWithOutputAndActorUsecaseInput_IntEnum) domain.IntEnum {
	// 						switch s {
	// 						case protobuf.DoingSomethingWithOutputAndActorUsecaseInput_INT_ONE:
	// 							return domain.IntOne
	// 						case protobuf.DoingSomethingWithOutputAndActorUsecaseInput_INT_THREE:
	// 							return domain.IntThree
	// 						case protobuf.DoingSomethingWithOutputAndActorUsecaseInput_INT_TWO:
	// 							return domain.IntTwo
	//
	// 						default:
	// 							var t domain.IntEnum
	// 							return t
	// 						}
	// 					}(inputProtoType.GetIntEnumSliceParam()[i])
	// 				}
	// 				return t
	// 			}(),
	// 		)
	//
	// 		return m
	// 	}()
	//
	// 	usecase, err := h.DoingSomethingWithOutputAndActorUsecaseFactory.Generate(r.Context())
	// 	if err != nil {
	// 		h.HandleError(w, r, err)
	// 		return
	// 	}
	// 	actor, err := ParseSomeActorToApplicationSomeActorDescription(r)
	// 	if err != nil {
	// 		h.HandleError(w, r, err)
	// 		return
	// 	}
	// 	outputType, err := usecase.DoSomethingWithOutputAndActor(
	// 		r.Context(), input, actor,
	// 	)
	// 	if err != nil {
	// 		h.HandleError(w, r, err)
	// 		return
	// 	}
	// 	outputProtoType := protobuf.DoingSomethingWithOutputAndActorUsecaseOutput{
	// 		StringParam: outputType.StringParam(),
	// 	}
	// 	b, err := proto.Marshal(&outputProtoType)
	// 	if err != nil {
	// 		h.HandleError(w, r, err)
	// 		return
	// 	}
	// 	if _, err := w.Write(b); err != nil {
	// 		h.HandleError(w, r, err)
	// 		return
	// 	}
	// 	return
	// }
	//
	// func (h Handlers) DoingSomethingWithOutputWithoutActorUsecaseDoSomethingWithOutputWithoutActorHandler(w http.ResponseWriter, r *http.Request) {
	// 	body, err := ioutil.ReadAll(r.Body)
	// 	if err != nil {
	// 		h.HandleError(w, r, err)
	// 		return
	// 	}
	//
	// 	inputProtoType := protobuf.DoingSomethingWithOutputWithoutActorUsecaseInput{}
	// 	if err := proto.Unmarshal(body, &inputProtoType); err != nil {
	// 		h.HandleError(w, r, err)
	// 		return
	// 	}
	// 	input := func() application.DoingSomethingWithOutputWithoutActorUsecaseInput {
	// 		m := application.NewDoingSomethingWithOutputWithoutActorUsecaseInput(
	// 			inputProtoType.GetStringParam(),
	// 		)
	//
	// 		return m
	// 	}()
	//
	// 	usecase, err := h.DoingSomethingWithOutputWithoutActorUsecaseFactory.Generate(r.Context())
	// 	if err != nil {
	// 		h.HandleError(w, r, err)
	// 		return
	// 	}
	// 	outputType, err := usecase.DoSomethingWithOutputWithoutActor(
	// 		r.Context(), input,
	// 	)
	// 	if err != nil {
	// 		h.HandleError(w, r, err)
	// 		return
	// 	}
	// 	outputProtoType := protobuf.DoingSomethingWithOutputWithoutActorUsecaseOutput{
	// 		StringParam: outputType.StringParam(),
	// 	}
	// 	b, err := proto.Marshal(&outputProtoType)
	// 	if err != nil {
	// 		h.HandleError(w, r, err)
	// 		return
	// 	}
	// 	if _, err := w.Write(b); err != nil {
	// 		h.HandleError(w, r, err)
	// 		return
	// 	}
	// 	return
	// }
	//
	// func (h Handlers) DoingSomethingWithoutOutputAndActorUsecaseDoSomethingWithoutOutputAndActorHandler(w http.ResponseWriter, r *http.Request) {
	// 	body, err := ioutil.ReadAll(r.Body)
	// 	if err != nil {
	// 		h.HandleError(w, r, err)
	// 		return
	// 	}
	//
	// 	inputProtoType := protobuf.DoingSomethingWithoutOutputAndActorUsecaseInput{}
	// 	if err := proto.Unmarshal(body, &inputProtoType); err != nil {
	// 		h.HandleError(w, r, err)
	// 		return
	// 	}
	// 	input := func() application.DoingSomethingWithoutOutputAndActorUsecaseInput {
	// 		m := application.NewDoingSomethingWithoutOutputAndActorUsecaseInput(
	// 			inputProtoType.GetStringParam(),
	// 		)
	//
	// 		return m
	// 	}()
	//
	// 	usecase, err := h.DoingSomethingWithoutOutputAndActorUsecaseFactory.Generate(r.Context())
	// 	if err != nil {
	// 		h.HandleError(w, r, err)
	// 		return
	// 	}
	// 	if err := usecase.DoSomethingWithoutOutputAndActor(
	// 		r.Context(), input,
	// 	); err != nil {
	// 		h.HandleError(w, r, err)
	// 		return
	// 	}
	// 	return
	// }
	//
	// func (h Handlers) DoingSomethingWithoutOutputWithActorUsecaseDoSomethingWithoutOutputWithActorHandler(w http.ResponseWriter, r *http.Request) {
	// 	body, err := ioutil.ReadAll(r.Body)
	// 	if err != nil {
	// 		h.HandleError(w, r, err)
	// 		return
	// 	}
	//
	// 	inputProtoType := protobuf.DoingSomethingWithoutOutputWithActorUsecaseInput{}
	// 	if err := proto.Unmarshal(body, &inputProtoType); err != nil {
	// 		h.HandleError(w, r, err)
	// 		return
	// 	}
	// 	input := func() application.DoingSomethingWithoutOutputWithActorUsecaseInput {
	// 		m := application.NewDoingSomethingWithoutOutputWithActorUsecaseInput(
	// 			inputProtoType.GetStringParam(),
	// 		)
	//
	// 		return m
	// 	}()
	//
	// 	usecase, err := h.DoingSomethingWithoutOutputWithActorUsecaseFactory.Generate(r.Context())
	// 	if err != nil {
	// 		h.HandleError(w, r, err)
	// 		return
	// 	}
	// 	actor, err := ParseSomeActorToApplicationSomeActorDescription(r)
	// 	if err != nil {
	// 		h.HandleError(w, r, err)
	// 		return
	// 	}
	// 	if err := usecase.DoSomethingWithoutOutputWithActor(
	// 		r.Context(), input, actor,
	// 	); err != nil {
	// 		h.HandleError(w, r, err)
	// 		return
	// 	}
	// 	return
	// }
	//
	// func NewMux(handler Handlers, middlewares ...func(http.Handler) http.Handler) *http.ServeMux {
	// 	mux := http.NewServeMux()
	// 	ApplyMux(mux, handler, middlewares...)
	// 	return mux
	// }
	//
	// func ApplyMux(mux *http.ServeMux, handler Handlers, middlewares ...func(http.Handler) http.Handler) {
	//
	// 	mux.Handle(
	// 		"DoingSomethingWithOutputAndActorUsecase/DoSomethingWithOutputAndActor",
	// 		applyMiddleware(http.HandlerFunc(handler.DoingSomethingWithOutputAndActorUsecaseDoSomethingWithOutputAndActorHandler), middlewares...),
	// 	)
	//
	// 	mux.Handle(
	// 		"DoingSomethingWithOutputWithoutActorUsecase/DoSomethingWithOutputWithoutActor",
	// 		applyMiddleware(http.HandlerFunc(handler.DoingSomethingWithOutputWithoutActorUsecaseDoSomethingWithOutputWithoutActorHandler), middlewares...),
	// 	)
	//
	// 	mux.Handle(
	// 		"DoingSomethingWithoutOutputAndActorUsecase/DoSomethingWithoutOutputAndActor",
	// 		applyMiddleware(http.HandlerFunc(handler.DoingSomethingWithoutOutputAndActorUsecaseDoSomethingWithoutOutputAndActorHandler), middlewares...),
	// 	)
	//
	// 	mux.Handle(
	// 		"DoingSomethingWithoutOutputWithActorUsecase/DoSomethingWithoutOutputWithActor",
	// 		applyMiddleware(http.HandlerFunc(handler.DoingSomethingWithoutOutputWithActorUsecaseDoSomethingWithoutOutputWithActorHandler), middlewares...),
	// 	)
	// }
	//
	// func applyMiddleware(h http.Handler, middlewares ...func(http.Handler) http.Handler) http.Handler {
	// 	for i := range middlewares {
	// 		h = middlewares[len(middlewares)-i](h)
	// 	}
	// 	return h
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

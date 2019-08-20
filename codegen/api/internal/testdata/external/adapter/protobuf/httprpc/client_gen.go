// Code generated by go-codegen api protobuf go_server httprpc; DO NOT EDIT

package httprpc

import (
	bytes "bytes"
	context "context"
	io "io"
	ioutil "io/ioutil"
	http "net/http"
	url "net/url"
	path "path"

	proto "github.com/golang/protobuf/proto"
	protobuf "github.com/hori-ryota/go-codegen/codegen/api/internal/testdata/external/adapter/protobuf"
	zaperr "github.com/hori-ryota/zaperr"
	zap "go.uber.org/zap"
)

func NewClient(
	httpClient *http.Client,
	urlBase url.URL,
	errorResponseParser ErrorResponseParser,
) Client {
	return Client{
		httpClient:          httpClient,
		urlBase:             urlBase,
		errorResponseParser: errorResponseParser,
	}
}

type Client struct {
	httpClient          *http.Client
	urlBase             url.URL
	errorResponseParser ErrorResponseParser
}

type ErrorResponseParser interface {
	ParseError(resp *http.Response) error
}

func (c Client) DoSomethingWithOutputAndActor(ctx context.Context, input DoingSomethingWithOutputAndActorUsecaseInput) (output DoingSomethingWithOutputAndActorUsecaseOutput, err error) {
	u := c.urlBase
	u.Path = path.Join(u.Path, "DoingSomethingWithOutputAndActorUsecase/DoSomethingWithOutputAndActor")

	inputProtoType := protobuf.DoingSomethingWithOutputAndActorUsecaseInput{
		StringParam:  input.StringParam,
		IntParam:     int64(input.IntParam),
		Int64Param:   input.Int64Param,
		UintParam:    uint64(input.UintParam),
		Uint64Param:  input.Uint64Param,
		Float32Param: input.Float32Param,
		Float64Param: input.Float64Param,
		BytesParam:   input.BytesParam,
		AnonymousNestedStructParam: func() *protobuf.DoingSomethingWithOutputAndActorUsecaseInput_AnonymousNestedStructParam {
			t := protobuf.DoingSomethingWithOutputAndActorUsecaseInput_AnonymousNestedStructParam{
				StringParam: input.AnonymousNestedStructParam.StringParam,
			}
			return &t
		}(),
		NamedNestedStructParam: func() *protobuf.NamedSomeType {
			t := protobuf.NamedSomeType{
				StringParam: input.NamedNestedStructParam.StringParam,
			}
			return &t
		}(),
		StringEnumParam: func(s StringEnum) protobuf.StringEnum {
			switch s {
			case StringA:
				return protobuf.StringEnum_STRING_A
			case StringB:
				return protobuf.StringEnum_STRING_B
			case StringC:
				return protobuf.StringEnum_STRING_C

			default:
				var t protobuf.StringEnum
				return t
			}
		}(input.StringEnumParam),
		IntEnumParam: func(s IntEnum) protobuf.IntEnum {
			switch s {
			case IntOne:
				return protobuf.IntEnum_INT_ONE
			case IntThree:
				return protobuf.IntEnum_INT_THREE
			case IntTwo:
				return protobuf.IntEnum_INT_TWO

			default:
				var t protobuf.IntEnum
				return t
			}
		}(input.IntEnumParam),
		StringSliceParam: input.StringSliceParam,
		IntSliceParam: func() []int64 {
			t := make([]int64, len(input.IntSliceParam))
			for i := range t {
				t[i] = int64(input.IntSliceParam[i])
			}
			return t
		}(),
		Int64SliceParam: input.Int64SliceParam,
		UintSliceParam: func() []uint64 {
			t := make([]uint64, len(input.UintSliceParam))
			for i := range t {
				t[i] = uint64(input.UintSliceParam[i])
			}
			return t
		}(),
		Uint64SliceParam:  input.Uint64SliceParam,
		Float32SliceParam: input.Float32SliceParam,
		Float64SliceParam: input.Float64SliceParam,
		BytesSliceParam:   input.BytesSliceParam,
		AnonymousNestedStructSliceParam: func() []*protobuf.DoingSomethingWithOutputAndActorUsecaseInput_AnonymousNestedStructSliceParam {
			t := make([]*protobuf.DoingSomethingWithOutputAndActorUsecaseInput_AnonymousNestedStructSliceParam, len(input.AnonymousNestedStructSliceParam))
			for i := range t {
				t[i] = func() *protobuf.DoingSomethingWithOutputAndActorUsecaseInput_AnonymousNestedStructSliceParam {
					t := protobuf.DoingSomethingWithOutputAndActorUsecaseInput_AnonymousNestedStructSliceParam{
						StringParam: input.AnonymousNestedStructSliceParam[i].StringParam,
					}
					return &t
				}()
			}
			return t
		}(),
		NamedNestedStructSliceParam: func() []*protobuf.NamedSomeType {
			t := make([]*protobuf.NamedSomeType, len(input.NamedNestedStructSliceParam))
			for i := range t {
				t[i] = func() *protobuf.NamedSomeType {
					t := protobuf.NamedSomeType{
						StringParam: input.NamedNestedStructSliceParam[i].StringParam,
					}
					return &t
				}()
			}
			return t
		}(),
		StringEnumSliceParam: func() []protobuf.StringEnum {
			t := make([]protobuf.StringEnum, len(input.StringEnumSliceParam))
			for i := range t {
				t[i] = func(s StringEnum) protobuf.StringEnum {
					switch s {
					case StringA:
						return protobuf.StringEnum_STRING_A
					case StringB:
						return protobuf.StringEnum_STRING_B
					case StringC:
						return protobuf.StringEnum_STRING_C

					default:
						var t protobuf.StringEnum
						return t
					}
				}(input.StringEnumSliceParam[i])
			}
			return t
		}(),
		IntEnumSliceParam: func() []protobuf.IntEnum {
			t := make([]protobuf.IntEnum, len(input.IntEnumSliceParam))
			for i := range t {
				t[i] = func(s IntEnum) protobuf.IntEnum {
					switch s {
					case IntOne:
						return protobuf.IntEnum_INT_ONE
					case IntThree:
						return protobuf.IntEnum_INT_THREE
					case IntTwo:
						return protobuf.IntEnum_INT_TWO

					default:
						var t protobuf.IntEnum
						return t
					}
				}(input.IntEnumSliceParam[i])
			}
			return t
		}(),
	}
	b, err := proto.Marshal(&inputProtoType)
	if err != nil {
		return output, err
	}
	r, err := http.NewRequest("POST", u.String(), bytes.NewReader(b))
	if err != nil {
		return output, err
	}
	r = r.WithContext(ctx)
	r.Header.Add("Content-Type", "application/protobuf")

	resp, err := c.httpClient.Do(r)
	if err != nil {
		return output, err
	}
	defer func() {
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
	}()

	if resp.StatusCode >= 400 {
		err := c.errorResponseParser.ParseError(resp)
		return output, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err = zaperr.Wrap(err, "failed to read response body", zap.Int("statusCode", resp.StatusCode))
		return output, err
	}
	outputProtoType := protobuf.DoingSomethingWithOutputAndActorUsecaseOutput{}
	if err := proto.Unmarshal(body, &outputProtoType); err != nil {
		err = zaperr.Wrap(err, "failed to parse response body", zap.String("body", string(body)))
		return output, err
	}
	return func() DoingSomethingWithOutputAndActorUsecaseOutput {
		m := NewDoingSomethingWithOutputAndActorUsecaseOutput(
			outputProtoType.GetStringParam(),
		)

		return m
	}(), nil
}

func (c Client) DoSomethingWithOutputWithoutActor(ctx context.Context, input DoingSomethingWithOutputWithoutActorUsecaseInput) (output DoingSomethingWithOutputWithoutActorUsecaseOutput, err error) {
	u := c.urlBase
	u.Path = path.Join(u.Path, "DoingSomethingWithOutputWithoutActorUsecase/DoSomethingWithOutputWithoutActor")

	inputProtoType := protobuf.DoingSomethingWithOutputWithoutActorUsecaseInput{
		StringParam: input.StringParam,
	}
	b, err := proto.Marshal(&inputProtoType)
	if err != nil {
		return output, err
	}
	r, err := http.NewRequest("POST", u.String(), bytes.NewReader(b))
	if err != nil {
		return output, err
	}
	r = r.WithContext(ctx)
	r.Header.Add("Content-Type", "application/protobuf")

	resp, err := c.httpClient.Do(r)
	if err != nil {
		return output, err
	}
	defer func() {
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
	}()

	if resp.StatusCode >= 400 {
		err := c.errorResponseParser.ParseError(resp)
		return output, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err = zaperr.Wrap(err, "failed to read response body", zap.Int("statusCode", resp.StatusCode))
		return output, err
	}
	outputProtoType := protobuf.DoingSomethingWithOutputWithoutActorUsecaseOutput{}
	if err := proto.Unmarshal(body, &outputProtoType); err != nil {
		err = zaperr.Wrap(err, "failed to parse response body", zap.String("body", string(body)))
		return output, err
	}
	return func() DoingSomethingWithOutputWithoutActorUsecaseOutput {
		m := NewDoingSomethingWithOutputWithoutActorUsecaseOutput(
			outputProtoType.GetStringParam(),
		)

		return m
	}(), nil
}

func (c Client) DoSomethingWithoutOutputAndActor(ctx context.Context, input DoingSomethingWithoutOutputAndActorUsecaseInput) error {
	u := c.urlBase
	u.Path = path.Join(u.Path, "DoingSomethingWithoutOutputAndActorUsecase/DoSomethingWithoutOutputAndActor")

	inputProtoType := protobuf.DoingSomethingWithoutOutputAndActorUsecaseInput{
		StringParam: input.StringParam,
	}
	b, err := proto.Marshal(&inputProtoType)
	if err != nil {
		return err
	}
	r, err := http.NewRequest("POST", u.String(), bytes.NewReader(b))
	if err != nil {
		return err
	}
	r = r.WithContext(ctx)
	r.Header.Add("Content-Type", "application/protobuf")

	resp, err := c.httpClient.Do(r)
	if err != nil {
		return err
	}
	defer func() {
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
	}()

	if resp.StatusCode >= 400 {
		err := c.errorResponseParser.ParseError(resp)
		return err
	}
	return nil
}

func (c Client) DoSomethingWithoutOutputWithActor(ctx context.Context, input DoingSomethingWithoutOutputWithActorUsecaseInput) error {
	u := c.urlBase
	u.Path = path.Join(u.Path, "DoingSomethingWithoutOutputWithActorUsecase/DoSomethingWithoutOutputWithActor")

	inputProtoType := protobuf.DoingSomethingWithoutOutputWithActorUsecaseInput{
		StringParam: input.StringParam,
	}
	b, err := proto.Marshal(&inputProtoType)
	if err != nil {
		return err
	}
	r, err := http.NewRequest("POST", u.String(), bytes.NewReader(b))
	if err != nil {
		return err
	}
	r = r.WithContext(ctx)
	r.Header.Add("Content-Type", "application/protobuf")

	resp, err := c.httpClient.Do(r)
	if err != nil {
		return err
	}
	defer func() {
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
	}()

	if resp.StatusCode >= 400 {
		err := c.errorResponseParser.ParseError(resp)
		return err
	}
	return nil
}

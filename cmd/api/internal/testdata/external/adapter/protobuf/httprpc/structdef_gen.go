// Code generated by go-codegen api protobuf go_client httprpc structdef; DO NOT EDIT

package httprpc

type DoingSomethingWithOutputAndActorUsecaseInput struct {
	StringParam                     string
	IntParam                        int
	Int64Param                      int64
	UintParam                       uint
	Uint64Param                     uint64
	Float32Param                    float32
	Float64Param                    float64
	BytesParam                      []byte
	AnonymousNestedStructParam      struct{ StringParam string }
	NamedNestedStructParam          NamedSomeType
	StringEnumParam                 StringEnum
	IntEnumParam                    IntEnum
	StringSliceParam                []string
	IntSliceParam                   []int
	Int64SliceParam                 []int64
	UintSliceParam                  []uint
	Uint64SliceParam                []uint64
	Float32SliceParam               []float32
	Float64SliceParam               []float64
	BytesSliceParam                 [][]byte
	AnonymousNestedStructSliceParam []struct{ StringParam string }
	NamedNestedStructSliceParam     []NamedSomeType
	StringEnumSliceParam            []StringEnum
	IntEnumSliceParam               []IntEnum
}

func NewDoingSomethingWithOutputAndActorUsecaseInput(
	stringParam string,
	intParam int,
	int64Param int64,
	uintParam uint,
	uint64Param uint64,
	float32Param float32,
	float64Param float64,
	bytesParam []byte,
	anonymousNestedStructParam struct{ StringParam string },
	namedNestedStructParam NamedSomeType,
	stringEnumParam StringEnum,
	intEnumParam IntEnum,
	stringSliceParam []string,
	intSliceParam []int,
	int64SliceParam []int64,
	uintSliceParam []uint,
	uint64SliceParam []uint64,
	float32SliceParam []float32,
	float64SliceParam []float64,
	bytesSliceParam [][]byte,
	anonymousNestedStructSliceParam []struct{ StringParam string },
	namedNestedStructSliceParam []NamedSomeType,
	stringEnumSliceParam []StringEnum,
	intEnumSliceParam []IntEnum,
) DoingSomethingWithOutputAndActorUsecaseInput {
	return DoingSomethingWithOutputAndActorUsecaseInput{
		StringParam:                     stringParam,
		IntParam:                        intParam,
		Int64Param:                      int64Param,
		UintParam:                       uintParam,
		Uint64Param:                     uint64Param,
		Float32Param:                    float32Param,
		Float64Param:                    float64Param,
		BytesParam:                      bytesParam,
		AnonymousNestedStructParam:      anonymousNestedStructParam,
		NamedNestedStructParam:          namedNestedStructParam,
		StringEnumParam:                 stringEnumParam,
		IntEnumParam:                    intEnumParam,
		StringSliceParam:                stringSliceParam,
		IntSliceParam:                   intSliceParam,
		Int64SliceParam:                 int64SliceParam,
		UintSliceParam:                  uintSliceParam,
		Uint64SliceParam:                uint64SliceParam,
		Float32SliceParam:               float32SliceParam,
		Float64SliceParam:               float64SliceParam,
		BytesSliceParam:                 bytesSliceParam,
		AnonymousNestedStructSliceParam: anonymousNestedStructSliceParam,
		NamedNestedStructSliceParam:     namedNestedStructSliceParam,
		StringEnumSliceParam:            stringEnumSliceParam,
		IntEnumSliceParam:               intEnumSliceParam,
	}
}

type NamedSomeType struct {
	StringParam string
}

func NewNamedSomeType(
	stringParam string,
) NamedSomeType {
	return NamedSomeType{
		StringParam: stringParam,
	}
}

type StringEnum string

const (
	StringA StringEnum = "A"
	StringB StringEnum = "B"
	StringC StringEnum = "C"
)

type IntEnum int

const (
	IntOne   IntEnum = 1
	IntThree IntEnum = 3
	IntTwo   IntEnum = 2
)

type DoingSomethingWithOutputAndActorUsecaseOutput struct {
	StringParam string
}

func NewDoingSomethingWithOutputAndActorUsecaseOutput(
	stringParam string,
) DoingSomethingWithOutputAndActorUsecaseOutput {
	return DoingSomethingWithOutputAndActorUsecaseOutput{
		StringParam: stringParam,
	}
}

type DoingSomethingWithOutputWithoutActorUsecaseInput struct {
	StringParam string
}

func NewDoingSomethingWithOutputWithoutActorUsecaseInput(
	stringParam string,
) DoingSomethingWithOutputWithoutActorUsecaseInput {
	return DoingSomethingWithOutputWithoutActorUsecaseInput{
		StringParam: stringParam,
	}
}

type DoingSomethingWithOutputWithoutActorUsecaseOutput struct {
	StringParam string
}

func NewDoingSomethingWithOutputWithoutActorUsecaseOutput(
	stringParam string,
) DoingSomethingWithOutputWithoutActorUsecaseOutput {
	return DoingSomethingWithOutputWithoutActorUsecaseOutput{
		StringParam: stringParam,
	}
}

type DoingSomethingWithoutOutputAndActorUsecaseInput struct {
	StringParam string
}

func NewDoingSomethingWithoutOutputAndActorUsecaseInput(
	stringParam string,
) DoingSomethingWithoutOutputAndActorUsecaseInput {
	return DoingSomethingWithoutOutputAndActorUsecaseInput{
		StringParam: stringParam,
	}
}

type DoingSomethingWithoutOutputWithActorUsecaseInput struct {
	StringParam string
}

func NewDoingSomethingWithoutOutputWithActorUsecaseInput(
	stringParam string,
) DoingSomethingWithoutOutputWithActorUsecaseInput {
	return DoingSomethingWithoutOutputWithActorUsecaseInput{
		StringParam: stringParam,
	}
}

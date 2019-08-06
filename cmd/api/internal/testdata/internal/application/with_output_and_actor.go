package application

import (
	"context"

	"github.com/hori-ryota/go-codegen/cmd/api/internal/testdata/internal/domain"
)

//genconstructor
type DoingSomethingWithOutputAndActorUsecase struct {
}

//genconstructor
type DoingSomethingWithOutputAndActorUsecaseInput struct {
	stringParam                string  `required:"" getter:""`
	intParam                   int     `required:"" getter:""`
	int64Param                 int64   `required:"" getter:""`
	uintParam                  uint    `required:"" getter:""`
	uint64Param                uint64  `required:"" getter:""`
	float32Param               float32 `required:"" getter:""`
	float64Param               float64 `required:"" getter:""`
	anonymousNestedStructParam struct {
		StringParam string
	} `required:"" getter:""`
	namedNestedStructParam NamedSomeType     `required:"" getter:""`
	stringEnumParam        domain.StringEnum `required:"" getter:""`
	intEnumParam           domain.IntEnum    `required:"" getter:""`

	stringSliceParam                []string  `required:"" getter:""`
	intSliceParam                   []int     `required:"" getter:""`
	int64SliceParam                 []int64   `required:"" getter:""`
	uintSliceParam                  []uint    `required:"" getter:""`
	uint64SliceParam                []uint64  `required:"" getter:""`
	float32SliceParam               []float32 `required:"" getter:""`
	float64SliceParam               []float64 `required:"" getter:""`
	anonymousNestedStructSliceParam []struct {
		StringParam string
	} `required:"" getter:""`
	namedNestedStructSliceParam []NamedSomeType     `required:"" getter:""`
	stringEnumSliceParam        []domain.StringEnum `required:"" getter:""`
	intEnumSliceParam           []domain.IntEnum    `required:"" getter:""`
}

//genconstructor
type NamedSomeType struct {
	stringParam string `required:"" getter:""`
}

//genconstructor
type DoingSomethingWithOutputAndActorUsecaseOutput struct {
	stringParam string `required:"" getter:""`
}

func (u DoingSomethingWithOutputAndActorUsecase) DoSomethingWithOutputAndActor(
	ctx context.Context,
	input DoingSomethingWithOutputAndActorUsecaseInput,
	someActor SomeActorDescription,
) (DoingSomethingWithOutputAndActorUsecaseOutput, domain.Error) {
	return DoingSomethingWithOutputAndActorUsecaseOutput{}, nil
}

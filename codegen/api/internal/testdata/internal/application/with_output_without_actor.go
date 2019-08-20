package application

import (
	"context"

	"github.com/hori-ryota/go-codegen/codegen/api/internal/testdata/internal/domain"
)

//genconstructor
type DoingSomethingWithOutputWithoutActorUsecase struct {
}

//genconstructor
type DoingSomethingWithOutputWithoutActorUsecaseInput struct {
	stringParam string `required:"" getter:""`
}

//genconstructor
type DoingSomethingWithOutputWithoutActorUsecaseOutput struct {
	stringParam string `required:"" getter:""`
}

func (u DoingSomethingWithOutputWithoutActorUsecase) DoSomethingWithOutputWithoutActor(
	ctx context.Context,
	input DoingSomethingWithOutputWithoutActorUsecaseInput,
) (DoingSomethingWithOutputWithoutActorUsecaseOutput, domain.Error) {
	return DoingSomethingWithOutputWithoutActorUsecaseOutput{}, nil
}

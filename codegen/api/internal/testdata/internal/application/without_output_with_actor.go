package application

import (
	"context"

	"github.com/hori-ryota/go-codegen/codegen/api/internal/testdata/internal/domain"
)

//genconstructor
type DoingSomethingWithoutOutputWithActorUsecase struct {
}

//genconstructor
type DoingSomethingWithoutOutputWithActorUsecaseInput struct {
	stringParam string `required:"" getter:""`
}

func (u DoingSomethingWithoutOutputWithActorUsecase) DoSomethingWithoutOutputWithActor(
	ctx context.Context,
	input DoingSomethingWithoutOutputWithActorUsecaseInput,
	someActor SomeActorDescription,
) domain.Error {
	return nil
}

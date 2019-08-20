package application

import (
	"context"

	"github.com/hori-ryota/go-codegen/codegen/api/internal/testdata/internal/domain"
)

//genconstructor
type DoingSomethingWithoutOutputAndActorUsecase struct {
}

//genconstructor
type DoingSomethingWithoutOutputAndActorUsecaseInput struct {
	stringParam string `required:"" getter:""`
}

func (u DoingSomethingWithoutOutputAndActorUsecase) DoSomethingWithoutOutputAndActor(
	ctx context.Context,
	input DoingSomethingWithoutOutputAndActorUsecaseInput,
) domain.Error {
	return nil
}

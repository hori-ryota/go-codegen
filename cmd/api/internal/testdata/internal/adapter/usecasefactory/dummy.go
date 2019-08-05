package usecasefactory

import (
	"context"

	"github.com/hori-ryota/go-codegen/cmd/api/internal/testdata/internal/application"
	"github.com/hori-ryota/go-codegen/cmd/api/internal/testdata/internal/domain"
)

type DoingSomethingWithOutputAndActorUsecaseFactory struct {
}

func (f DoingSomethingWithOutputAndActorUsecaseFactory) Generate(ctx context.Context) (application.DoingSomethingWithOutputAndActorUsecase, domain.Error) {
	panic("not implemented")
}

type DoingSomethingWithOutputWithoutActorUsecaseFactory struct {
}

func (f DoingSomethingWithOutputWithoutActorUsecaseFactory) Generate(ctx context.Context) (application.DoingSomethingWithOutputWithoutActorUsecase, domain.Error) {
	panic("not implemented")
}

type DoingSomethingWithoutOutputAndActorUsecaseFactory struct {
}

func (f DoingSomethingWithoutOutputAndActorUsecaseFactory) Generate(ctx context.Context) (application.DoingSomethingWithoutOutputAndActorUsecase, domain.Error) {
	panic("not implemented")
}

type DoingSomethingWithoutOutputWithActorUsecaseFactory struct {
}

func (f DoingSomethingWithoutOutputWithActorUsecaseFactory) Generate(ctx context.Context) (application.DoingSomethingWithoutOutputWithActorUsecase, domain.Error) {
	panic("not implemented")
}

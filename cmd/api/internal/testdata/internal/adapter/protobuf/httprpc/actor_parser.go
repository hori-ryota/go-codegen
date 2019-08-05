package httprpc

import (
	http "net/http"

	"github.com/hori-ryota/go-codegen/cmd/api/internal/testdata/internal/application"
	"github.com/hori-ryota/go-codegen/cmd/api/internal/testdata/internal/domain"
)

func ParseSomeActorToApplicationSomeActorDescription(r *http.Request) (application.SomeActorDescription, domain.Error) {
	panic("not implemented")
}

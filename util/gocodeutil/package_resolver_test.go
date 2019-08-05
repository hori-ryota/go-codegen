package gocodeutil_test

import (
	"testing"

	"github.com/hori-ryota/go-codegen/util/gocodeutil"
	"github.com/stretchr/testify/assert"
)

func TestLocalPathToGoPackagePath(t *testing.T) {
	for _, tt := range []struct {
		s string
		t string
	}{
		{
			s: ".",
			t: "github.com/hori-ryota/go-codegen/util/gocodeutil",
		},
		{
			s: "..",
			t: "github.com/hori-ryota/go-codegen/util",
		},
		{
			s: "../gocodeutil",
			t: "github.com/hori-ryota/go-codegen/util/gocodeutil",
		},
	} {
		tt := tt
		t.Run(tt.s, func(t *testing.T) {
			got, err := gocodeutil.LocalPathToGoPackagePath(tt.s)
			assert.NoError(t, err)
			assert.Equal(t, tt.t, got)
		})
	}
}

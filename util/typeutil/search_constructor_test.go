package typeutil_test

import (
	"go/token"
	"go/types"
	"testing"

	"github.com/hori-ryota/go-codegen/util/typeutil"
	"github.com/stretchr/testify/assert"
)

func TestSearchConstructor(t *testing.T) {
	pkg := types.NewPackage("example.com/foo", "foo")

	correctConstructor := types.NewFunc(
		token.NoPos,
		pkg,
		"NewFoo",
		types.NewSignature(nil, nil, types.NewTuple(types.NewVar(token.NoPos, pkg, "", types.Universe.Lookup("string").Type())), false),
	)

	pkg.Scope().Insert(correctConstructor)

	t.Run("found", func(t *testing.T) {
		assert.Equal(
			t,
			correctConstructor,
			typeutil.SearchConstructor(types.NewNamed(
				types.NewTypeName(token.NoPos, pkg, "Foo", nil),
				types.NewStruct(nil, nil),
				nil,
			)),
		)
	})

	t.Run("not found", func(t *testing.T) {
		assert.Nil(
			t,
			typeutil.SearchConstructor(types.NewNamed(
				types.NewTypeName(token.NoPos, pkg, "unknown", nil),
				types.NewStruct(nil, nil),
				nil,
			)),
		)
	})
}

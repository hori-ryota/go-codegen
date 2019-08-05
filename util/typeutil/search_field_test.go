package typeutil_test

import (
	"go/token"
	"go/types"
	"testing"

	"github.com/hori-ryota/go-codegen/util/typeutil"
	"github.com/stretchr/testify/assert"
)

func TestSearchField(t *testing.T) {
	fooPackage := types.NewPackage("example.com/foo", "foo")
	for _, tt := range []struct {
		name         string
		typ          *types.Named
		paramName    string
		onlyExported bool
		want         *types.Var
	}{
		{
			name: "Named: has field",
			typ: types.NewNamed(
				types.NewTypeName(token.NoPos, fooPackage, "NamedStruct", nil),
				types.NewStruct([]*types.Var{
					types.NewField(
						token.NoPos,
						fooPackage,
						"Param",
						types.Universe.Lookup("string").Type(),
						false,
					),
				}, nil),
				nil,
			),
			paramName: "param",
			want: types.NewField(
				token.NoPos,
				fooPackage,
				"Param",
				types.Universe.Lookup("string").Type(),
				false,
			),
		},
		{
			name: "Named: has notfield: nil",
			typ: types.NewNamed(
				types.NewTypeName(token.NoPos, fooPackage, "NamedStruct", nil),
				types.NewStruct([]*types.Var{
					types.NewField(
						token.NoPos,
						fooPackage,
						"Param",
						types.Universe.Lookup("string").Type(),
						false,
					),
				}, nil),
				nil,
			),
			paramName: "unknown",
			want:      nil,
		},
		{
			name: "Named: has internal field",
			typ: types.NewNamed(
				types.NewTypeName(token.NoPos, fooPackage, "NamedStruct", nil),
				types.NewStruct([]*types.Var{
					types.NewField(
						token.NoPos,
						fooPackage,
						"param",
						types.Universe.Lookup("string").Type(),
						false,
					),
				}, nil),
				nil,
			),
			paramName: "param",
			want: types.NewField(
				token.NoPos,
				fooPackage,
				"param",
				types.Universe.Lookup("string").Type(),
				false,
			),
		},
		{
			name: "Named: only exported: has exported field: nil",
			typ: types.NewNamed(
				types.NewTypeName(token.NoPos, fooPackage, "NamedStruct", nil),
				types.NewStruct([]*types.Var{
					types.NewField(
						token.NoPos,
						fooPackage,
						"Param",
						types.Universe.Lookup("string").Type(),
						false,
					),
				}, nil),
				nil,
			),
			paramName:    "param",
			onlyExported: true,
			want: types.NewField(
				token.NoPos,
				fooPackage,
				"Param",
				types.Universe.Lookup("string").Type(),
				false,
			),
		},
		{
			name: "Named: only exported: has internal field: nil",
			typ: types.NewNamed(
				types.NewTypeName(token.NoPos, fooPackage, "NamedStruct", nil),
				types.NewStruct([]*types.Var{
					types.NewField(
						token.NoPos,
						fooPackage,
						"param",
						types.Universe.Lookup("string").Type(),
						false,
					),
				}, nil),
				nil,
			),
			paramName:    "param",
			onlyExported: true,
			want:         nil,
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := typeutil.SearchField(tt.typ, tt.paramName, tt.onlyExported)
			assert.Equal(t, tt.want, got)
		})
	}
}

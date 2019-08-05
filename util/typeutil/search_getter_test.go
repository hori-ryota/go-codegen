package typeutil_test

import (
	"go/token"
	"go/types"
	"testing"

	"github.com/hori-ryota/go-codegen/util/typeutil"
	"github.com/stretchr/testify/assert"
)

func TestSearchGetter(t *testing.T) {
	fooPackage := types.NewPackage("example.com/foo", "foo")
	for _, tt := range []struct {
		name         string
		typ          *types.Named
		paramName    string
		onlyExported bool
		want         *types.Func
	}{
		{
			name: "Named: has getter",
			typ: types.NewNamed(
				types.NewTypeName(token.NoPos, fooPackage, "NamedStruct", nil),
				types.NewStruct(nil, nil),
				[]*types.Func{
					types.NewFunc(
						token.NoPos,
						fooPackage,
						"Param",
						types.NewSignature(nil, nil, types.NewTuple(types.NewVar(token.NoPos, fooPackage, "", types.Universe.Lookup("string").Type())), false),
					),
				},
			),
			paramName: "param",
			want: types.NewFunc(
				token.NoPos,
				fooPackage,
				"Param",
				types.NewSignature(nil, nil, types.NewTuple(types.NewVar(token.NoPos, fooPackage, "", types.Universe.Lookup("string").Type())), false),
			),
		},
		{
			name: "Named: has getter with prefix 'Get'",
			typ: types.NewNamed(
				types.NewTypeName(token.NoPos, fooPackage, "NamedStruct", nil),
				types.NewStruct(nil, nil),
				[]*types.Func{
					types.NewFunc(
						token.NoPos,
						fooPackage,
						"GetParam",
						types.NewSignature(nil, nil, types.NewTuple(types.NewVar(token.NoPos, fooPackage, "", types.Universe.Lookup("string").Type())), false),
					),
				},
			),
			paramName: "param",
			want: types.NewFunc(
				token.NoPos,
				fooPackage,
				"GetParam",
				types.NewSignature(nil, nil, types.NewTuple(types.NewVar(token.NoPos, fooPackage, "", types.Universe.Lookup("string").Type())), false),
			),
		},
		{
			name: "Named: notfound: nil",
			typ: types.NewNamed(
				types.NewTypeName(token.NoPos, fooPackage, "NamedStruct", nil),
				types.NewStruct(nil, nil),
				[]*types.Func{
					types.NewFunc(
						token.NoPos,
						fooPackage,
						"GetParam",
						types.NewSignature(nil, nil, types.NewTuple(types.NewVar(token.NoPos, fooPackage, "", types.Universe.Lookup("string").Type())), false),
					),
				},
			),
			paramName: "unknown",
			want:      nil,
		},
		{
			name: "Named: only exported: has exported getter",
			typ: types.NewNamed(
				types.NewTypeName(token.NoPos, fooPackage, "NamedStruct", nil),
				types.NewStruct(nil, nil),
				[]*types.Func{
					types.NewFunc(
						token.NoPos,
						fooPackage,
						"Param",
						types.NewSignature(nil, nil, types.NewTuple(types.NewVar(token.NoPos, fooPackage, "", types.Universe.Lookup("string").Type())), false),
					),
				},
			),
			paramName:    "param",
			onlyExported: true,
			want: types.NewFunc(
				token.NoPos,
				fooPackage,
				"Param",
				types.NewSignature(nil, nil, types.NewTuple(types.NewVar(token.NoPos, fooPackage, "", types.Universe.Lookup("string").Type())), false),
			),
		},
		{
			name: "Named: only exported: has internal getter: nil",
			typ: types.NewNamed(
				types.NewTypeName(token.NoPos, fooPackage, "NamedStruct", nil),
				types.NewStruct(nil, nil),
				[]*types.Func{
					types.NewFunc(
						token.NoPos,
						fooPackage,
						"param",
						types.NewSignature(nil, nil, types.NewTuple(types.NewVar(token.NoPos, fooPackage, "", types.Universe.Lookup("string").Type())), false),
					),
				},
			),
			paramName:    "param",
			onlyExported: true,
			want:         nil,
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := typeutil.SearchGetter(tt.typ, tt.paramName, tt.onlyExported)
			assert.Equal(t, tt.want, got)
		})
	}
}

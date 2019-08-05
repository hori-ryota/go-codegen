package typeutil_test

import (
	"go/token"
	"go/types"
	"testing"

	"github.com/hori-ryota/go-codegen/util/typeutil"
	"github.com/stretchr/testify/assert"
)

func TestSearchSetter(t *testing.T) {
	fooPackage := types.NewPackage("example.com/foo", "foo")
	for _, tt := range []struct {
		name         string
		typ          *types.Named
		paramName    string
		onlyExported bool
		want         *types.Func
	}{
		{
			name: "Named: has setter with prefix 'Set'",
			typ: types.NewNamed(
				types.NewTypeName(token.NoPos, fooPackage, "NamedStruct", nil),
				types.NewStruct(nil, nil),
				[]*types.Func{
					types.NewFunc(
						token.NoPos,
						fooPackage,
						"SetParam",
						types.NewSignature(nil, nil, types.NewTuple(types.NewVar(token.NoPos, fooPackage, "", types.Universe.Lookup("string").Type())), false),
					),
				},
			),
			paramName: "param",
			want: types.NewFunc(
				token.NoPos,
				fooPackage,
				"SetParam",
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
						"SetParam",
						types.NewSignature(nil, nil, types.NewTuple(types.NewVar(token.NoPos, fooPackage, "", types.Universe.Lookup("string").Type())), false),
					),
				},
			),
			paramName: "unknown",
			want:      nil,
		},
		{
			name: "Named: not exported: notfound: nil",
			typ: types.NewNamed(
				types.NewTypeName(token.NoPos, fooPackage, "NamedStruct", nil),
				types.NewStruct(nil, nil),
				[]*types.Func{
					types.NewFunc(
						token.NoPos,
						fooPackage,
						"setParam",
						types.NewSignature(nil, nil, types.NewTuple(types.NewVar(token.NoPos, fooPackage, "", types.Universe.Lookup("string").Type())), false),
					),
				},
			),
			paramName:    "param",
			onlyExported: true,
			want:         nil,
		},
		{
			name: "Named: has internal setter with prefix 'set'",
			typ: types.NewNamed(
				types.NewTypeName(token.NoPos, fooPackage, "NamedStruct", nil),
				types.NewStruct(nil, nil),
				[]*types.Func{
					types.NewFunc(
						token.NoPos,
						fooPackage,
						"setParam",
						types.NewSignature(nil, nil, types.NewTuple(types.NewVar(token.NoPos, fooPackage, "", types.Universe.Lookup("string").Type())), false),
					),
				},
			),
			paramName:    "param",
			onlyExported: false,
			want: types.NewFunc(
				token.NoPos,
				fooPackage,
				"setParam",
				types.NewSignature(nil, nil, types.NewTuple(types.NewVar(token.NoPos, fooPackage, "", types.Universe.Lookup("string").Type())), false),
			),
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := typeutil.SearchSetter(tt.typ, tt.paramName, tt.onlyExported)
			assert.Equal(t, tt.want, got)
		})
	}
}

package typeutil_test

import (
	"go/token"
	"go/types"
	"testing"

	"github.com/hori-ryota/go-codegen/util/typeutil"
	"github.com/stretchr/testify/assert"
)

func TestPrinter_PrintGetValue(t *testing.T) {
	fooPackage := types.NewPackage("example.com/foo", "foo")
	for _, tt := range []struct {
		name       string
		parentName string
		parentType types.Type
		targetName string
		dstPackage *types.Package
		want       string
		failed     bool
	}{
		{
			name:       "Named: has public field",
			parentName: "p",
			parentType: types.NewNamed(
				types.NewTypeName(token.NoPos, fooPackage, "NamedStruct", nil),
				types.NewStruct([]*types.Var{
					types.NewField(
						token.NoPos,
						fooPackage,
						"Value",
						types.Universe.Lookup("string").Type(),
						false,
					),
				}, nil),
				nil,
			),
			targetName: "value",
			want:       "p.Value",
		},
		{
			name:       "Named: only private field: failed",
			parentName: "p",
			parentType: types.NewNamed(
				types.NewTypeName(token.NoPos, fooPackage, "NamedStruct", nil),
				types.NewStruct([]*types.Var{
					types.NewField(
						token.NoPos,
						fooPackage,
						"value",
						types.Universe.Lookup("string").Type(),
						false,
					),
				}, nil),
				nil,
			),
			targetName: "value",
			failed:     true,
		},
		{
			name:       "Named: has getter",
			parentName: "p",
			parentType: types.NewNamed(
				types.NewTypeName(token.NoPos, fooPackage, "NamedStruct", nil),
				types.NewStruct(nil, nil),
				[]*types.Func{
					types.NewFunc(
						token.NoPos,
						fooPackage,
						"Value",
						types.NewSignature(nil, nil, types.NewTuple(types.NewVar(token.NoPos, fooPackage, "", types.Universe.Lookup("string").Type())), false),
					),
				},
			),
			targetName: "value",
			want:       "p.Value()",
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			p := typeutil.NewPrinter(tt.dstPackage)
			got, err := p.PrintGetValue(tt.parentName, tt.parentType, tt.targetName)
			if tt.failed {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

package typeutil_test

import (
	"go/token"
	"go/types"
	"testing"

	"github.com/hori-ryota/go-codegen/util/typeutil"
	"github.com/stretchr/testify/assert"
)

func TestPrinter_PrintConverterWitoutErrorCheck(t *testing.T) {
	fooPackage := types.NewPackage("example.com/foo", "foo")
	for _, tt := range []struct {
		name       string
		sn         string
		s          types.Type
		t          types.Type
		dstPackage *types.Package
		want       string
	}{
		{
			name: "assignable: string to string",
			sn:   "p",
			s:    types.Universe.Lookup("string").Type(),
			t:    types.Universe.Lookup("string").Type(),
			want: "p",
		},
		{
			name: "convertible: int to int64",
			sn:   "p",
			s:    types.Universe.Lookup("int").Type(),
			t:    types.Universe.Lookup("int64").Type(),
			want: "int64(p)",
		},
		{
			name: "assignable: slice: []string to []string",
			sn:   "p",
			s:    types.NewSlice(types.Universe.Lookup("string").Type()),
			t:    types.NewSlice(types.Universe.Lookup("string").Type()),
			want: "p",
		},
		{
			name: "convertible: slice: []NamedString to []string",
			sn:   "p",
			s: types.NewSlice(types.NewNamed(
				types.NewTypeName(token.NoPos, fooPackage, "NamedString", nil),
				types.Universe.Lookup("string").Type(),
				nil,
			)),
			t: types.NewSlice(types.Universe.Lookup("string").Type()),
			want: `func() []string {
				t0 := make([]string, len(p))
				for i0 := range t0 {
					t0[i0] = string(p[i0])
				}
				return t0
			}()`,
		},
		{
			name: "convertible: slice: []string to []NamedString",
			sn:   "p",
			s:    types.NewSlice(types.Universe.Lookup("string").Type()),
			t: types.NewSlice(types.NewNamed(
				types.NewTypeName(token.NoPos, fooPackage, "NamedString", nil),
				types.Universe.Lookup("string").Type(),
				nil,
			)),
			want: `func() []foo.NamedString {
				t0 := make([]foo.NamedString, len(p))
				for i0 := range t0 {
					t0[i0] = foo.NamedString(p[i0])
				}
				return t0
			}()`,
		},
		{
			name: "pointer: string to *string",
			sn:   "p",
			s:    types.Universe.Lookup("string").Type(),
			t:    types.NewPointer(types.Universe.Lookup("string").Type()),
			want: `func() *string {
				t0 := p
				return &t0
			}()`,
		},
		{
			name: "pointer: *string to string",
			sn:   "p",
			s:    types.NewPointer(types.Universe.Lookup("string").Type()),
			t:    types.Universe.Lookup("string").Type(),
			want: `func(s *string) string {
				if s == nil {
					var t0 string
					return t0
				}
				return *s
			}(p)`,
		},
		{
			name: "slice: pointer []string to []*string",
			sn:   "p",
			s:    types.NewSlice(types.Universe.Lookup("string").Type()),
			t:    types.NewSlice(types.NewPointer(types.Universe.Lookup("string").Type())),
			want: `func() []*string {
				t0 := make([]*string, len(p))
				for i0 := range t0 {
					t0[i0] = func() *string {
				t1 := p[i0]
				return &t1
			}()
				}
				return t0
			}()`,
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			p := typeutil.NewPrinter(tt.dstPackage)
			got, err := p.PrintConverterWitoutErrorCheck(tt.sn, tt.s, tt.t)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

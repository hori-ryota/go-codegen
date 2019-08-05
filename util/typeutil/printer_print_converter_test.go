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
				t := make([]string, len(p))
				for i := range t {
					t[i] = string(p[i])
				}
				return t
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
				t := make([]foo.NamedString, len(p))
				for i := range t {
					t[i] = foo.NamedString(p[i])
				}
				return t
			}()`,
		},
		{
			name: "pointer: string to *string",
			sn:   "p",
			s:    types.Universe.Lookup("string").Type(),
			t:    types.NewPointer(types.Universe.Lookup("string").Type()),
			want: `func() *string {
				t := p
				return &t
			}()`,
		},
		{
			name: "pointer: *string to string",
			sn:   "p",
			s:    types.NewPointer(types.Universe.Lookup("string").Type()),
			t:    types.Universe.Lookup("string").Type(),
			want: `func(s *string) string {
				if s == nil {
					var t string
					return t
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
				t := make([]*string, len(p))
				for i := range t {
					t[i] = func() *string {
				t := p[i]
				return &t
			}()
				}
				return t
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

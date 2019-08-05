package gocodeutil_test

import (
	"testing"

	"github.com/hori-ryota/go-codegen/util/gocodeutil"
	"github.com/stretchr/testify/assert"
)

func TestGoFmtImports(t *testing.T) {
	for _, tt := range []struct {
		name string
		src  map[string]string
		want string
	}{
		{
			name: "empty",
			src:  map[string]string{},
			want: ``,
		},
		{
			name: "only standard package",
			src: map[string]string{
				"fmt":   "fmt",
				"bytes": "bytes",
				"json":  "encoding/json",
			},
			want: `import (
bytes "bytes"
json "encoding/json"
fmt "fmt"
)`,
		},
		{
			name: "with alias",
			src: map[string]string{
				"afmt":  "fmt",
				"bytes": "bytes",
				"ajson": "encoding/json",
			},
			want: `import (
bytes "bytes"
ajson "encoding/json"
afmt "fmt"
)`,
		},
		{
			name: "with non-standard package",
			src: map[string]string{
				"fmt":        "fmt",
				"bytes":      "bytes",
				"json":       "encoding/json",
				"go-codegen": "github.com/hori-ryota/go-codegen",
			},
			want: `import (
bytes "bytes"
json "encoding/json"
fmt "fmt"

go-codegen "github.com/hori-ryota/go-codegen"
)`,
		},
		{
			name: "with aliased non-standard package",
			src: map[string]string{
				"fmt":     "fmt",
				"codegen": "github.com/hori-ryota/go-codegen",
			},
			want: `import (
fmt "fmt"

codegen "github.com/hori-ryota/go-codegen"
)`,
		},
		{
			name: "only non-standard package",
			src: map[string]string{
				"go-codegen": "github.com/hori-ryota/go-codegen",
			},
			want: `import (
go-codegen "github.com/hori-ryota/go-codegen"
)`,
		},
	} {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, gocodeutil.FmtImports(tt.src))
		})
	}
}

/*
Copyright © 2019 hori-ryota <hori.ryota@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package go_accessor

import (
	"fmt"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/hori-ryota/go-codegen/codegen/go_accessor"
	"github.com/hori-ryota/go-codegen/util/gocodeutil"
	"github.com/spf13/cobra"
	"golang.org/x/tools/go/loader"
)

func NewGoAccessorCmd() *cobra.Command {
	var targetDir string
	var useStdOut bool

	goAccessorCmd := &cobra.Command{
		Use:   "go_accessor",
		Short: "accessor genreator for Go",
		Long:  `go_accessor is a generator command for Go accessor.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return Run(targetDir, useStdOut)
		},
	}

	goAccessorCmd.Flags().StringVarP(&targetDir, "targetDir", "t", ".", "target directory")
	if err := goAccessorCmd.MarkFlagDirname("targetDir"); err != nil {
		panic(err)
	}
	goAccessorCmd.Flags().BoolVar(&useStdOut, "useStdOut", false, "use stdout")

	return goAccessorCmd
}

func Run(targetDir string, useStdOut bool) error {
	targetDir = filepath.FromSlash(targetDir)

	lconf := loader.Config{
		ParserMode: parser.ParseComments,
		TypeChecker: types.Config{
			Importer: importer.ForCompiler(token.NewFileSet(), "source", nil),
		},
		AllowErrors: true,
	}
	targetPkgPath, err := gocodeutil.LocalPathToGoPackagePath(targetDir)
	if err != nil {
		return err
	}
	lconf.Import(targetPkgPath)
	lprog, err := lconf.Load()
	if err != nil {
		return err
	}
	generated, err := go_accessor.Generate(lprog.Package(targetPkgPath))
	if err != nil {
		return err
	}
	if strings.TrimSpace(generated) == "" {
		return nil
	}

	if useStdOut {
		fmt.Print(generated)
		return nil
	}

	return ioutil.WriteFile(filepath.Join(targetDir, "accessor_gen.go"), []byte(generated), 0644)
}

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
package go_definition

import (
	"fmt"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/hori-ryota/go-codegen/util/gocodeutil"
	"github.com/spf13/cobra"
	"golang.org/x/tools/go/loader"
)

func NewGoDefinitionCmd(
	targetDir *string,
	codes *[]string,
) *cobra.Command {
	var outputDir string
	var useStdOut bool

	goDefinitionCmd := &cobra.Command{
		Use:   "go_definition",
		Short: "definition genreator for Go",
		Long:  `go_definition is a generator command for Go definition.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return Run(*targetDir, *codes, outputDir, useStdOut)
		},
	}

	goDefinitionCmd.Flags().StringVarP(&outputDir, "outputDir", "o", ".", "output directory")
	if err := goDefinitionCmd.MarkFlagDirname("outputDir"); err != nil {
		panic(err)
	}
	goDefinitionCmd.Flags().BoolVar(&useStdOut, "useStdOut", false, "use stdout")

	return goDefinitionCmd
}

func Run(targetDir string, codes []string, outputDir string, useStdOut bool) error {
	targetDir = filepath.FromSlash(targetDir)
	outputDir = filepath.FromSlash(outputDir)

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

	outputPkgPath, err := gocodeutil.LocalPathToGoPackagePath(outputDir)
	if err != nil {
		return err
	}
	generated, err := Generate(lprog.Package(targetPkgPath), codes, lprog.Package(outputPkgPath).Pkg)
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

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return err
	}

	return ioutil.WriteFile(filepath.Join(outputDir, "error_definition_gen.go"), []byte(generated), 0644)
}

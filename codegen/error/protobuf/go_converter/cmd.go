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
package go_converter

import (
	"fmt"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/hori-ryota/go-codegen/util/gocodeutil"
	"github.com/spf13/cobra"
	"golang.org/x/tools/go/loader"
)

func NewGoConverterCmd() *cobra.Command {
	var errorDefinitionDir string
	var protoDir string
	var outputDir string
	var useStdOut bool

	goConverterCmd := &cobra.Command{
		Use:   "go_converter",
		Short: "converter generator of go for error",
		Long:  `go_converter is a generator command of converter for error in Go.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return Run(
				errorDefinitionDir,
				protoDir,
				outputDir,
				useStdOut,
			)
		},
	}

	goConverterCmd.Flags().StringVarP(&errorDefinitionDir, "errorDefinitionDir", "t", ".", "target error definition directory")
	if err := goConverterCmd.MarkFlagDirname("errorDefinitionDir"); err != nil {
		panic(err)
	}
	goConverterCmd.Flags().StringVarP(&protoDir, "protoDir", "p", ".", "target proto types of Go directory")
	if err := goConverterCmd.MarkFlagDirname("protoDir"); err != nil {
		panic(err)
	}
	goConverterCmd.Flags().StringVarP(&outputDir, "outputDir", "o", "", "output directory")
	if err := goConverterCmd.MarkFlagRequired("outputDir"); err != nil {
		panic(err)
	}
	if err := goConverterCmd.MarkFlagDirname("outputDir"); err != nil {
		panic(err)
	}
	goConverterCmd.Flags().BoolVar(&useStdOut, "useStdOut", false, "use stdout")

	return goConverterCmd
}

func Run(
	errorDefinitionDir string,
	protoDir string,
	outputDir string,
	useStdOut bool,
) error {
	errorDefinitionDir = filepath.FromSlash(errorDefinitionDir)
	protoDir = filepath.FromSlash(protoDir)
	outputDir = filepath.FromSlash(outputDir)

	lconf := loader.Config{
		ParserMode: parser.ParseComments,
		TypeChecker: types.Config{
			Importer: importer.ForCompiler(token.NewFileSet(), "source", nil),
		},
		AllowErrors: true,
	}
	errorDefinitionPkgPath, err := gocodeutil.LocalPathToGoPackagePath(errorDefinitionDir)
	if err != nil {
		return err
	}
	lconf.Import(errorDefinitionPkgPath)
	protoPkgPath, err := gocodeutil.LocalPathToGoPackagePath(protoDir)
	if err != nil {
		return err
	}
	lconf.Import(protoPkgPath)
	lprog, err := lconf.Load()
	if err != nil {
		return err
	}

	outputPackage, err := gocodeutil.LocalPathToGoPackagePath(outputDir)
	if err != nil {
		return err
	}

	generated, err := Generate(
		lprog.Package(errorDefinitionPkgPath).Pkg,
		lprog.Package(protoPkgPath).Pkg,
		types.NewPackage(outputPackage, path.Base(outputPackage)),
	)
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

	return ioutil.WriteFile(filepath.Join(outputDir, "error_converter_gen.go"), []byte(generated), 0644)
}

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
package structdef

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

func NewStructdefCmd() *cobra.Command {
	var usecaseDir string
	var outputDir string
	var useStdOut bool

	structdefCmd := &cobra.Command{
		Use:   "structdef",
		Short: "structdef handler genreator",
		Long:  `structdef is a generator command of structdef for client.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return Run(
				usecaseDir,
				outputDir,
				useStdOut,
			)
		},
	}

	structdefCmd.Flags().StringVarP(&usecaseDir, "usecaseDir", "t", ".", "target usecase directory")
	if err := structdefCmd.MarkFlagDirname("usecaseDir"); err != nil {
		panic(err)
	}
	structdefCmd.Flags().StringVarP(&outputDir, "outputDir", "o", "", "output directory")
	if err := structdefCmd.MarkFlagRequired("outputDir"); err != nil {
		panic(err)
	}
	if err := structdefCmd.MarkFlagDirname("outputDir"); err != nil {
		panic(err)
	}
	structdefCmd.Flags().BoolVar(&useStdOut, "useStdOut", false, "use stdout")

	return structdefCmd
}

func Run(
	usecaseDir string,
	outputDir string,
	useStdOut bool,
) error {
	usecaseDir = filepath.FromSlash(usecaseDir)
	outputDir = filepath.FromSlash(outputDir)

	lconf := loader.Config{
		ParserMode: parser.ParseComments,
		TypeChecker: types.Config{
			Importer: importer.ForCompiler(token.NewFileSet(), "source", nil),
		},
		AllowErrors: true,
	}
	usecasePkgPath, err := gocodeutil.LocalPathToGoPackagePath(usecaseDir)
	if err != nil {
		return err
	}
	lconf.Import(usecasePkgPath)
	lprog, err := lconf.Load()
	if err != nil {
		return err
	}

	outputPackage, err := gocodeutil.LocalPathToGoPackagePath(outputDir)
	if err != nil {
		return err
	}

	generated, err := Generate(
		lprog.Package(usecasePkgPath),
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

	return ioutil.WriteFile(filepath.Join(outputDir, "structdef_gen.go"), []byte(generated), 0644)
}

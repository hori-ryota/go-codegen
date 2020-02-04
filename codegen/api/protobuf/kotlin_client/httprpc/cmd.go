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
package httprpc

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
	"github.com/hori-ryota/go-strcase"
	"github.com/spf13/cobra"
	"golang.org/x/tools/go/loader"
)

func NewHttprpcCmd() *cobra.Command {
	var usecaseDir string
	var outputDir string
	var kotlinPackage string
	var kotlinClassName string
	var structdefPackage string
	var useStdOut bool

	httprpcCmd := &cobra.Command{
		Use:   "httprpc",
		Short: "httprpc client generator",
		Long:  `httprpc is a generator command of httprpc client.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return Run(
				usecaseDir,
				outputDir,
				kotlinPackage,
				kotlinClassName,
				structdefPackage,
				useStdOut,
			)
		},
	}

	httprpcCmd.Flags().StringVarP(&usecaseDir, "usecaseDir", "t", ".", "target usecase directory")
	if err := httprpcCmd.MarkFlagDirname("usecaseDir"); err != nil {
		panic(err)
	}
	httprpcCmd.Flags().StringVarP(&outputDir, "outputDir", "o", "", "output directory")
	if err := httprpcCmd.MarkFlagRequired("outputDir"); err != nil {
		panic(err)
	}
	if err := httprpcCmd.MarkFlagDirname("outputDir"); err != nil {
		panic(err)
	}
	httprpcCmd.Flags().StringVar(&kotlinPackage, "kotlinPackage", "", "package of kotlin")
	if err := httprpcCmd.MarkFlagRequired("kotlinPackage"); err != nil {
		panic(err)
	}
	httprpcCmd.Flags().StringVar(&kotlinClassName, "kotlinClassName", "", "className of kotlin")
	if err := httprpcCmd.MarkFlagRequired("kotlinClassName"); err != nil {
		panic(err)
	}
	httprpcCmd.Flags().StringVar(&structdefPackage, "structdefPackage", "", "package of structdef")
	httprpcCmd.Flags().BoolVar(&useStdOut, "useStdOut", false, "use stdout")

	return httprpcCmd
}

func Run(
	usecaseDir string,
	outputDir string,
	kotlinPackage string,
	kotlinClassName string,
	structdefPackage string,
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

	if structdefPackage == "" {
		structdefPackage = kotlinPackage
	}

	generated, err := Generate(
		lprog.Package(usecasePkgPath),
		kotlinPackage,
		kotlinClassName,
		structdefPackage,
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

	outputFileDir := filepath.Join(outputDir, filepath.FromSlash(strings.ReplaceAll(kotlinPackage, ".", "/")))
	if err := os.MkdirAll(outputFileDir, 0755); err != nil {
		return err
	}

	return ioutil.WriteFile(filepath.Join(outputFileDir, strcase.ToUpperCamel(kotlinClassName)+".kt"), []byte(generated), 0644)
}

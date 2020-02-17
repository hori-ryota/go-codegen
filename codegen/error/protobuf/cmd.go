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
package protobuf

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

	"github.com/hori-ryota/go-codegen/codegen/error/protobuf/go_converter"
	"github.com/hori-ryota/go-codegen/util/gocodeutil"
	"github.com/spf13/cobra"
	"golang.org/x/tools/go/loader"
)

func NewProtobufCmd() *cobra.Command {
	var targetDir string
	var outputFile string
	var protoPackage string
	var goPackage string
	var javaPackage string
	var javaOuterClassName string
	var useStdOut bool

	protobufCmd := &cobra.Command{
		Use:   "protobuf",
		Short: "proto generator for error",
		Long:  `protobuf is a generator command for proto.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return Run(
				targetDir,
				outputFile,
				protoPackage,
				goPackage,
				javaPackage,
				javaOuterClassName,
				useStdOut,
			)
		},
	}

	protobufCmd.Flags().StringVarP(&targetDir, "targetDir", "t", ".", "target directory")
	if err := protobufCmd.MarkFlagDirname("targetDir"); err != nil {
		panic(err)
	}
	protobufCmd.Flags().StringVarP(&outputFile, "outputFile", "o", "", "output file")
	if err := protobufCmd.MarkFlagRequired("outputFile"); err != nil {
		panic(err)
	}
	if err := protobufCmd.MarkFlagFilename("outputFile"); err != nil {
		panic(err)
	}
	protobufCmd.Flags().StringVar(&protoPackage, "protoPackage", "", "package of proto")
	if err := protobufCmd.MarkFlagRequired("protoPackage"); err != nil {
		panic(err)
	}
	protobufCmd.Flags().StringVar(&goPackage, "goPackage", "", "package of Go")
	protobufCmd.Flags().StringVar(&javaPackage, "javaPackage", "", "package of java")
	protobufCmd.Flags().StringVar(&javaOuterClassName, "javaOuterClassName", "", "outer class name of java")
	protobufCmd.Flags().BoolVar(&useStdOut, "useStdOut", false, "use stdout")

	protobufCmd.AddCommand(go_converter.NewGoConverterCmd())
	return protobufCmd
}

func Run(
	targetDir string,
	outputFile string,
	protoPackage string,
	goPackage string,
	javaPackage string,
	javaOuterClassName string,
	useStdOut bool,
) error {
	targetDir = filepath.FromSlash(targetDir)
	outputFile = filepath.FromSlash(outputFile)

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

	if goPackage == "" {
		goPackage, err = gocodeutil.LocalPathToGoPackagePath(filepath.Dir(outputFile))
		if err != nil {
			return err
		}
	}

	generated, err := Generate(lprog.Package(targetPkgPath), protoPackage, goPackage, javaPackage, javaOuterClassName)
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

	if err := os.MkdirAll(filepath.Dir(outputFile), 0755); err != nil {
		return err
	}

	return ioutil.WriteFile(outputFile, []byte(generated), 0644)
}

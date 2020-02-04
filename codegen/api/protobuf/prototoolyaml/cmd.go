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
package prototoolyaml

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

func NewPrototoolYamlCmd() *cobra.Command {
	var outputDir string
	var javaOutputDir string
	var useStdOut bool

	prototoolYamlCmd := &cobra.Command{
		Use:   "prototoolyaml",
		Short: "yaml generator for prototool",
		Long:  `prototoolyaml is a yaml generator command for prototool.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return Run(
				outputDir,
				javaOutputDir,
				useStdOut,
			)
		},
	}

	prototoolYamlCmd.Flags().StringVarP(&outputDir, "outputDir", "o", "", "output directory")
	if err := prototoolYamlCmd.MarkFlagRequired("outputDir"); err != nil {
		panic(err)
	}
	if err := prototoolYamlCmd.MarkFlagDirname("outputDir"); err != nil {
		panic(err)
	}
	prototoolYamlCmd.Flags().StringVar(&javaOutputDir, "javaOutputDir", "", "output directory of java")
	prototoolYamlCmd.Flags().BoolVar(&useStdOut, "useStdOut", false, "use stdout")

	return prototoolYamlCmd
}

func Run(
	outputDir string,
	javaOutputDir string,
	useStdOut bool,
) error {
	outputDir = filepath.FromSlash(outputDir)
	javaOutputDir = filepath.FromSlash(javaOutputDir)

	generated, err := Generate(outputDir, javaOutputDir)
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

	return ioutil.WriteFile(filepath.Join(outputDir, "prototool.yaml"), []byte(generated), 0644)
}

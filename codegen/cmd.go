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
package codegen

import (
	"github.com/hori-ryota/go-codegen/codegen/api"
	"github.com/hori-ryota/go-codegen/codegen/error"
	"github.com/hori-ryota/go-codegen/codegen/go_accessor"
	"github.com/hori-ryota/go-codegen/codegen/go_constructor"
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "go-codegen",
		Short: "code generator for Go",
		Long: `go-codegen is a generator library for Go.

go-codegen mainly generates codes for Domain Driven Development.`,
	}
	rootCmd.AddCommand(go_accessor.NewGoAccessorCmd())
	rootCmd.AddCommand(go_constructor.NewGoConstructorCmd())
	rootCmd.AddCommand(error.NewErrorCmd())
	rootCmd.AddCommand(api.NewAPICmd())
	return rootCmd
}

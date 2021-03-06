package go_accessor_test

import (
	cmd "github.com/hori-ryota/go-codegen/codegen"
	"github.com/mattn/go-shellwords"
	"github.com/spf13/cobra"
)

func Example() {
	input := "go_accessor -t internal/testdata"

	cmd := prepareCmdWithStdout(input)
	err := cmd.Execute()
	if err != nil {
		panic(err)
	}
	// Output:
	// // Code generated by go-codegen go_accessor; DO NOT EDIT.
	//
	// package example
	//
	// import (
	// 	"encoding"
	// )
	//
	// func (m Person) ID() string {
	// 	return m.id
	// }
	//
	// func (m Person) Name() string {
	// 	return m.name
	// }
	//
	// func (m *Person) Rename(s string) {
	// 	m.name = s
	// }
	//
	// func (m Person) Tags() []string {
	// 	return m.tags
	// }
	//
	// func (m *Person) SetTags(s []string) {
	// 	m.tags = s
	// }
	//
	// func (m Person) Text() encoding.TextMarshaler {
	// 	return m.text
	// }
	//
	// func (m *Person) SetText(s encoding.TextMarshaler) {
	// 	m.text = s
	// }
}

func prepareCmdWithStdout(input string) *cobra.Command {
	input += " --useStdOut"
	args, err := shellwords.Parse(input)
	if err != nil {
		panic(err)
	}
	cmd := cmd.NewRootCmd()
	cmd.SetArgs(args)
	return cmd
}

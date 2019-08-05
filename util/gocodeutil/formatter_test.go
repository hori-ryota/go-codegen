package gocodeutil_test

import (
	"fmt"

	"github.com/hori-ryota/go-codegen/util/gocodeutil"
)

func ExampleFormatGoFileFromString_removeUnusedImports() {
	src := `package example

import (
    "usedA"
    used "usedB"
    "unusedA"
    unused "unusedB"
)

func foo() {
    usedA.Bar()
    used.Baz()
}`

	dst, err := gocodeutil.FormatGoFileFromString(src)
	if err != nil {
		panic(err)
	}
	fmt.Print(dst)
	// Output:
	// package example
	//
	// import (
	// 	"usedA"
	// 	used "usedB"
	// )
	//
	// func foo() {
	// 	usedA.Bar()
	// 	used.Baz()
	// }
}

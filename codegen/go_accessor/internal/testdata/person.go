package example

import "encoding"

type Person struct {
	id   string                 `getter:""`
	name string                 `getter:"" setter:"Rename"`
	tags []string               `getter:"" setter:""`
	text encoding.TextMarshaler `getter:"" setter:""`
}

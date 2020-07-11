package api

import (
	"github.com/majintao/swag/testdata/nested2"
)

type Foo struct {
	Field1      string `validate:"required"`
	OutsideData *nested2.Body
	InsideData  Bar      `validate:"required"`
	ArrayField1 []string `validate:"required"`
	ArrayField2 []Bar    `validate:"required"`
}

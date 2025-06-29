package annotate

import "reflect"

type CustomAnnotation interface {
	Tag() string
	Enabled() bool
	KeyValue() bool
	Apply(reflect.Value) error
}


package annotate

import "reflect"

type CustomAnnotation interface {
	Metadata() AnnotationMetadata
	Apply(reflect.Value) error
}

type AnnotationMetadata struct {
	Tag       string
	Enabled   bool
	KeyValue  bool
}

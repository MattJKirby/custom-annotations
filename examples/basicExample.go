package examples

import (
	"custom-annotations/annotate"
	"reflect"
)

type MyCustomAnnotation struct {}

func (mca *MyCustomAnnotation) Metadata() annotate.AnnotationMetadata {
	return annotate.AnnotationMetadata {
		Tag: "mca",
		Enabled: true,
	}
}

func (mca *MyCustomAnnotation) Apply(value reflect.Value) error {
	return nil
}

func (mca *MyCustomAnnotation) ApplyValues(value reflect.Value) error {
	return nil
}

type AnnotatedStruct struct {
	myField string `a:"mca,2"`
}

func init() {
	annotator := annotate.NewAnnotator(annotate.WithName("a"))
	annotator.Register(&MyCustomAnnotation{})

	annotator.Annotate(&AnnotatedStruct{})

}
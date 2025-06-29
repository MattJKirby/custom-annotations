package examples

import (
	"custom-annotations/annotate"
	"reflect"
)

type MyCustomAnnotation struct {}

func (mca *MyCustomAnnotation) Tag() string {
	return "mca"
}

func (mca *MyCustomAnnotation) Enabled() bool {
	return true
}

func (mca *MyCustomAnnotation) KeyValue() bool {
	return false
}

func (mca* MyCustomAnnotation) Apply(value reflect.Value) error {
	return nil
}

type AnnotatedStruct struct {
	myField string `ca:"mca,2"`
}

func init() {
	annotator := annotate.NewAnnotator()
	annotator.Register(&MyCustomAnnotation{})

	annotator.Annotate(&AnnotatedStruct{})

}
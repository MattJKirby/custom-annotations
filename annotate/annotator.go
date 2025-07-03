package annotate

import (
	"fmt"
	"reflect"
)

type Annotator struct {
	name string
	annotations *registry[CustomAnnotation]
}

var defaultAnnotator = &Annotator{
	name: "annotator",
	annotations: NewRegistry(func(c CustomAnnotation) string {		
		return c.Metadata().Tag
	}),
}

func NewAnnotator(opts ...AnnotateOption) *Annotator {
	annotator := defaultAnnotator
	for _,opt := range opts {
		opt(annotator)
	}
	return annotator
}

func (a *Annotator) Register(cAnn CustomAnnotation) error {
	return a.annotations.Register(cAnn)
}

func (a *Annotator) annotateField(field reflect.StructField) error {
	if val, ok := field.Tag.Lookup(a.name); ok {
		fmt.Println(val)
	}
	return nil
}

func (a *Annotator) Annotate(val any) error {
	t := reflect.TypeOf(val)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		return fmt.Errorf("error annotating: '%s' is not a struct", t.Name())
	}

	for i := range t.NumField() {
		field := t.Field(i)
		a.annotateField(field)
	}

	return nil
}
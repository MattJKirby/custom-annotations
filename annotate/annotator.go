package annotate

import (
	"fmt"
	"reflect"
)

type Annotator struct {
	singleAnnotations map[string]CustomAnnotation
	kvAnnotations map[string]CustomAnnotation
}

func NewAnnotator() *Annotator {
	return &Annotator{
		singleAnnotations: make(map[string]CustomAnnotation),
		kvAnnotations: make(map[string]CustomAnnotation),
	}
}

func (a *Annotator) register(registered map[string]CustomAnnotation, cAnn CustomAnnotation) error {
	if _, exists := registered[cAnn.Tag()]; exists {
	return fmt.Errorf("error registering annotation: annoatation with tag '%s' already exists", cAnn.Tag())
 }
 registered[cAnn.Tag()] = cAnn
 return nil
}

func (a *Annotator) Register(cAnn CustomAnnotation) error {
	if cAnn.KeyValue() {
		return a.register(a.kvAnnotations, cAnn)
	}
	return a.register(a.singleAnnotations, cAnn)
}

func (a *Annotator) annotateField(field reflect.StructField) error {
	if val, ok := field.Tag.Lookup("ca"); ok {
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
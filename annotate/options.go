package annotate

type AnnotateOption func(*Annotator)

func WithName(name string) AnnotateOption {
	return func(a *Annotator) {
		a.name = name
	}
}
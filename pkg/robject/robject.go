package robject

type RObject struct {
	object any
}

func New(object any) *RObject {
	return &RObject{
		object: object,
	}
}

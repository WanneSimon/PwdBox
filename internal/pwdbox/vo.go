package pwdbox

type EntityType interface {
	Account | Platform
}

type PageData[T EntityType] struct {
	Page  int
	Size  int
	Total int
	Data  []T
}

func NewPageData[T EntityType](page int, size int, total int, list []T) PageData[T] {
	return PageData[T]{
		Page:  page,
		Size:  size,
		Total: total,
		Data:  list,
	}
}
func EmptyPageData[T EntityType]() PageData[T] {
	return PageData[T]{
		Page:  0,
		Size:  0,
		Total: 0,
		Data:  []T{},
	}
}

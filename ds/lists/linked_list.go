package lists

type Node[T comparable] struct {
	value T
}

type LinkedList[T comparable] interface {
	Insert()
	DeleteByKey()
	DeleteByIndex()
}

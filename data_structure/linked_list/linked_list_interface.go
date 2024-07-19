package linked_list

type LinkedListInterface interface {
	IsEmpty() bool
	Size() uint64
	Get(index uint64) interface{}
	Find(value interface{}) (uint64, bool)
	Set(index uint64, value interface{}) error
	AddLast(value interface{})
	RemoveAt(index uint64) error
}

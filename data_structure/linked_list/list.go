package linked_list

import (
	"errors"
	"reflect"
)

type Node struct {
	Value interface{}
	Next  *Node
}

type List struct {
	head *Node
	size uint64
}

func NewList() *List {
	return &List{}
}

func (l *List) IsEmpty() bool {
	return l.head == nil
}

func (l *List) Size() uint64 {
	return l.size
}

func (l *List) Get(index uint64) interface{} {
	if index < 0 || index >= l.size {
		return nil
	}
	current := l.head
	var i uint64
	for i = 0; i < index; i++ {
		current = current.Next
	}
	return current.Value
}

func (l *List) Find(value interface{}) (uint64, bool) {
	current := l.head
	var index uint64 = 0
	for current != nil {
		if reflect.DeepEqual(current.Value, value) {
			return index, true
		}
		current = current.Next
		index++
	}
	return 0, false
}
func (l *List) Set(index uint64, value interface{}) error {
	if index < 0 || index >= l.size {
		return errors.New("index out of range")
	}
	head := l.head
	var i uint64
	for i = 0; i < index; i++ {
		head = head.Next
	}
	head.Value = value
	return nil
}

func (l *List) AddLast(value interface{}) {
	newNode := &Node{
		Value: value,
	}
	if l.head == nil {
		l.head = newNode
	} else {
		current := l.head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
	l.size++
}

func (l *List) RemoveAt(index uint64) error {
	if index < 0 || index >= l.size {
		return errors.New("index out of range")
	}
	if index == 0 {
		l.head = l.head.Next
		l.size--
		return nil
	}
	current := l.head
	var i uint64
	for i = 0; current != nil && i < index-1; i++ {
		current = current.Next
	}
	if current == nil || current.Next == nil {
		return errors.New("index out of range")
	}
	current.Next = current.Next.Next
	l.size--
	return nil
}

package set

import (
	"errors"
	"go-dsa/list"
)

type ListSet struct {
	list *list.DoubleLinkedList
}

func NewListSet() *ListSet {
	return &ListSet{
		list: list.NewDoubleLinkedList().(*list.DoubleLinkedList),
	}
}

func (l *ListSet) Size() int {
	return l.list.Len()
}

func (l *ListSet) IsEmpty() bool {
	return l.list.IsEmpty()
}

func (l *ListSet) Clear() {
	l.list.Clear()
}

func (l *ListSet) IsContain(element interface{}) bool {
	return l.list.IsContain(element)
}

func (l *ListSet) Add(element interface{}) (interface{}, error) {
	if element == nil {
		return nil, errors.New("element is nil")
	}

	idx, err := l.list.IndexOf(element)
	if err == nil {
		return l.list.Set(idx, element)
	} else {
		l.list.Append(element)
		return element, nil
	}
}

func (l *ListSet) Remove(element interface{}) (interface{}, error) {
	if element == nil {
		return nil, errors.New("element is nil")
	}

	idx, err := l.list.IndexOf(element)
	if err == nil {
		return l.list.Remove(idx)
	} else {
		return nil, err
	}
}

func (l *ListSet) Traversal(visitor Visitor) {
	if visitor == nil {
		return
	}

	size := l.list.Len()
	for i := 0; i < size; i++ {
		ele, _ := l.list.Get(i)
		if visitor.Visit(ele) {
			return
		}
	}
}

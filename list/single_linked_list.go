package list

import (
	"encoding/json"
	"errors"
	"reflect"
)

type singleLinkedListNode struct {
	ele  interface{}
	next *singleLinkedListNode
}

func newSingleLinkedNode(ele interface{}, next *singleLinkedListNode) *singleLinkedListNode {
	return &singleLinkedListNode{
		ele,
		next,
	}
}

type SingleLinkedList struct {
	listCommon
	first *singleLinkedListNode
}

func NewSingleLinkedList() List {
	return &SingleLinkedList{
		first: newSingleLinkedNode(nil, nil),
	}
}

func (l *SingleLinkedList) getNodeByIdx(idx int) (*singleLinkedListNode, error) {
	chkErr := l.indexCheck(idx, l.length)
	if chkErr != nil {
		return nil, chkErr
	}

	cursor := l.first
	for i := 0; i < idx; i++ {
		cursor = cursor.next
	}
	return cursor, nil
}

func (l *SingleLinkedList) IsContain(ele interface{}) bool {
	_, idxErr := l.IndexOf(ele)
	if idxErr != nil {
		return false
	}
	return true
}

func (l *SingleLinkedList) Append(ele interface{}) {
	_ = l.Insert(l.length, ele)
}

func (l *SingleLinkedList) Get(idx int) (interface{}, error) {
	cursor, nodeErr := l.getNodeByIdx(idx)
	if nodeErr != nil {
		return nil, nodeErr
	}

	return cursor.ele, nil
}

func (l *SingleLinkedList) Set(idx int, ele interface{}) (interface{}, error) {
	cursor, nodeErr := l.getNodeByIdx(idx)
	if nodeErr != nil {
		return nil, nodeErr
	}

	old := cursor.ele
	cursor.ele = ele
	return old, nil
}

func (l *SingleLinkedList) Insert(idx int, ele interface{}) error {
	if idx == 0 {
		l.first.ele = ele
	} else {
		cursor, nodeErr := l.getNodeByIdx(idx - 1)
		if nodeErr != nil {
			return nodeErr
		}

		cursor.next = newSingleLinkedNode(ele, cursor.next)
	}
	l.length++
	return nil
}

func (l *SingleLinkedList) Remove(idx int) (interface{}, error) {
	chkErr := l.indexCheck(idx, l.length)
	if chkErr != nil {
		return nil, chkErr
	}
	var removed *singleLinkedListNode
	if idx == 0 {
		removed = l.first
		l.first = l.first.next
	} else {
		cursor, nodeErr := l.getNodeByIdx(idx - 1)
		if nodeErr != nil {
			return nil, nodeErr
		}
		removed = cursor.next
		cursor.next = cursor.next.next
	}
	l.length--

	return removed.ele, nil
}

func (l *SingleLinkedList) IndexOf(ele interface{}) (int, error) {
	cursor := l.first
	for i := 0; i < l.length; i++ {
		if reflect.DeepEqual(cursor.ele, ele) {
			return i, nil
		}
		cursor = cursor.next
	}
	return invalidIdx, errors.New(elementNotInListError)
}

func (l *SingleLinkedList) Clear() {
	l.first.ele = nil
	l.first.next = nil
	l.length = 0
}

func (l *SingleLinkedList) String() string {
	eles := "["
	cursor := l.first
	if cursor.ele != nil {
		for i := 0; i < l.length-1; i++ {
			eleBytes, _ := json.Marshal(cursor.ele)
			eles += string(eleBytes) + ","
			cursor = cursor.next
		}
		eleBytes, _ := json.Marshal(cursor.ele)
		eles += string(eleBytes)
	}
	eles += "]"
	return eles
}

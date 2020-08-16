package list

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"runtime"
)

// TODO: 析构函数打印节点是否消除

type doubleLinkedListNode struct {
	ele  interface{}
	prev *doubleLinkedListNode
	next *doubleLinkedListNode
}

func newDoubleLinkedNode(ele interface{}, prev, next *doubleLinkedListNode) *doubleLinkedListNode {
	return &doubleLinkedListNode{
		ele,
		prev,
		next,
	}
}

type DoubleLinkedList struct {
	listCommon
	headGuard *doubleLinkedListNode
	tailGuard *doubleLinkedListNode
}

func NewDoubleLinkedList() List {
	front := newDoubleLinkedNode(nil, nil, nil)
	back := newDoubleLinkedNode(nil, nil, nil)
	front.next = back
	back.prev = front
	return &DoubleLinkedList{
		headGuard: front,
		tailGuard: back,
	}
}

func (l *DoubleLinkedList) getNodeByIdx(idx int) (*doubleLinkedListNode, error) {
	chkErr := l.indexCheck(idx, l.length)
	if chkErr != nil {
		return nil, chkErr
	}

	if idx < l.length>>1 {
		cursor := l.headGuard.next
		for i := 0; i < idx; i++ {
			cursor = cursor.next
		}
		return cursor, nil
	} else {
		cursor := l.tailGuard.prev
		for i := l.length; i > idx+1; i-- {
			cursor = cursor.prev
		}
		return cursor, nil
	}
}

func (l *DoubleLinkedList) IsContain(ele interface{}) bool {
	_, idxErr := l.IndexOf(ele)
	if idxErr != nil {
		return false
	}
	return true
}

func (l *DoubleLinkedList) Append(ele interface{}) {
	_ = l.Insert(l.length, ele)
}

func (l *DoubleLinkedList) Get(idx int) (interface{}, error) {
	cursor, nodeErr := l.getNodeByIdx(idx)
	if nodeErr != nil {
		return nil, nodeErr
	}

	return cursor.ele, nil
}

func (l *DoubleLinkedList) Set(idx int, ele interface{}) (interface{}, error) {
	cursor, nodeErr := l.getNodeByIdx(idx)
	if nodeErr != nil {
		return nil, nodeErr
	}

	old := cursor.ele
	cursor.ele = ele
	return old, nil
}

func (l *DoubleLinkedList) Insert(idx int, ele interface{}) error {
	if l.length == 0 {
		insertedNode := newDoubleLinkedNode(ele, l.headGuard, l.tailGuard)
		runtime.SetFinalizer(insertedNode, func(i *doubleLinkedListNode) {
			fmt.Println(i.ele)
			fmt.Println("I'm deleting")
		})
		l.headGuard.next = insertedNode
		l.tailGuard.prev = insertedNode
	} else {
		if idx == l.length {
			insertedNode := newDoubleLinkedNode(ele, l.tailGuard.prev, l.tailGuard)
			runtime.SetFinalizer(insertedNode, func(i *doubleLinkedListNode) {
				fmt.Println(i.ele)
				fmt.Println("I'm deleting")
			})
			l.tailGuard.prev.next = insertedNode
			l.tailGuard.prev = insertedNode
		} else {
			cursor, nodeErr := l.getNodeByIdx(idx)
			if nodeErr != nil {
				return nodeErr
			}

			insertedNode := newDoubleLinkedNode(ele, cursor.prev, cursor)
			runtime.SetFinalizer(insertedNode, func(i *doubleLinkedListNode) {
				fmt.Println(i.ele)
				fmt.Println("I'm deleting")
			})
			cursor.prev.next = insertedNode
			cursor.prev = insertedNode
		}
	}
	l.length++
	return nil
}

func (l *DoubleLinkedList) Remove(idx int) (interface{}, error) {
	chkErr := l.indexCheck(idx, l.length)
	if chkErr != nil {
		return nil, chkErr
	}
	removed, nodeErr := l.getNodeByIdx(idx)
	if nodeErr != nil {
		return nil, nodeErr
	}
	removed.prev.next = removed.next
	removed.next.prev = removed.prev
	l.length--

	return removed.ele, nil
}

func (l *DoubleLinkedList) IndexOf(ele interface{}) (int, error) {
	cursor := l.headGuard.next
	for i := 0; i < l.length; i++ {
		if reflect.DeepEqual(cursor.ele, ele) {
			return i, nil
		}
		cursor = cursor.next
	}
	return invalidIdx, errors.New(elementNotInListError)
}

func (l *DoubleLinkedList) Clear() {
	l.headGuard = nil
	l.tailGuard = nil
	l.length = 0
}

func (l *DoubleLinkedList) String() string {
	eles := "["
	cursor := l.headGuard.next
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

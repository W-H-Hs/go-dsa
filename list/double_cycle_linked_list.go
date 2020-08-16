package list

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"time"
)

type doubleCycleLinkedListNode struct {
	ele  interface{}
	prev *doubleCycleLinkedListNode
	next *doubleCycleLinkedListNode
}

func newDoubleCycleLinkedNode(ele interface{}, prev *doubleCycleLinkedListNode, next *doubleCycleLinkedListNode) *doubleCycleLinkedListNode {
	return &doubleCycleLinkedListNode{
		ele,
		prev,
		next,
	}
}

type DoubleCycleLinkedList struct {
	listCommon
	headGuard *doubleCycleLinkedListNode
	tailGuard *doubleCycleLinkedListNode
	current   *doubleCycleLinkedListNode
}

func NewDoubleCycleLinkedList() *DoubleCycleLinkedList {
	head := newDoubleCycleLinkedNode(nil, nil, nil)
	tail := newDoubleCycleLinkedNode(nil, nil, nil)
	return &DoubleCycleLinkedList{
		headGuard: head,
		tailGuard: tail,
		current:   head.next,
	}
}

func (d *DoubleCycleLinkedList) Current() interface{} {
	return d.current.ele
}

func (d *DoubleCycleLinkedList) Next() {
	d.current = d.current.next
}
func (d *DoubleCycleLinkedList) Prev() {
	d.current = d.current.prev
}

func (d *DoubleCycleLinkedList) Reset() {
	d.current = d.headGuard.next
}

func (d *DoubleCycleLinkedList) RemoveCurrent() {
	d.current.prev.next = d.current.next
	d.current.next.prev = d.current.prev
	d.current = d.current.next
	d.length--
}

func (d *DoubleCycleLinkedList) getNodeByIdx(idx int) (*doubleCycleLinkedListNode, error) {
	chkErr := d.indexCheck(idx, d.length)
	if chkErr != nil {
		return nil, chkErr
	}

	if idx < d.length>>1 {
		cursor := d.headGuard.next
		for i := 0; i < idx; i++ {
			cursor = cursor.next
		}
		return cursor, nil
	} else {
		cursor := d.tailGuard.prev
		for i := d.length - 1; i > idx; i-- {
			cursor = cursor.prev
		}
		return cursor, nil
	}
}

func (d *DoubleCycleLinkedList) IsContain(ele interface{}) bool {
	_, idxErr := d.IndexOf(ele)
	if idxErr != nil {
		return false
	}
	return true
}

func (d *DoubleCycleLinkedList) Append(ele interface{}) {
	_ = d.Insert(d.length, ele)
}

func (d *DoubleCycleLinkedList) Get(idx int) (interface{}, error) {
	cursor, nodeErr := d.getNodeByIdx(idx)
	if nodeErr != nil {
		return nil, nodeErr
	}

	return cursor.ele, nil
}

func (d *DoubleCycleLinkedList) Set(idx int, ele interface{}) (interface{}, error) {
	cursor, nodeErr := d.getNodeByIdx(idx)
	if nodeErr != nil {
		return nil, nodeErr
	}

	old := cursor.ele
	cursor.ele = ele
	return old, nil
}

func (d *DoubleCycleLinkedList) R() {
	c := d.headGuard.next
	for {
		fmt.Println(c.ele)
		c = c.prev
		time.Sleep(time.Second >> 1)
	}
}

func (d *DoubleCycleLinkedList) Insert(idx int, ele interface{}) error {
	switch {
	case d.length == 0:
		insertedNode := newDoubleCycleLinkedNode(ele, nil, nil)
		d.headGuard.next = insertedNode
		d.tailGuard.prev = insertedNode
		insertedNode.next = insertedNode
		insertedNode.prev = insertedNode
		d.current = d.headGuard.next
	case idx == 0:
		insertedNode := newDoubleCycleLinkedNode(ele, d.tailGuard.prev, d.headGuard.next)
		insertedNode.prev.next = insertedNode
		insertedNode.next.prev = insertedNode
		d.headGuard.next = insertedNode
		d.current = d.headGuard.next
	case idx == d.length:
		insertedNode := newDoubleCycleLinkedNode(ele, d.tailGuard.prev, d.headGuard.next)
		insertedNode.prev.next = insertedNode
		insertedNode.next.prev = insertedNode
		d.tailGuard.prev = insertedNode
	default:
		cursor, nodeErr := d.getNodeByIdx(idx)
		if nodeErr != nil {
			return nodeErr
		}
		inserted := newDoubleCycleLinkedNode(ele, cursor, cursor.next)
		inserted.next.prev = inserted
		inserted.prev.next = inserted
	}
	d.length++
	return nil
}

func (d *DoubleCycleLinkedList) Remove(idx int) (interface{}, error) {
	chkErr := d.indexCheck(idx, d.length)
	if chkErr != nil {
		return nil, chkErr
	}

	removed, nodeErr := d.getNodeByIdx(idx)
	if nodeErr != nil {
		return nil, nodeErr
	}
	if idx == 0 {
		d.headGuard.next = d.headGuard.next.next
	}
	if idx == d.length {
		d.tailGuard.prev = d.tailGuard.prev.prev
	}
	removed.prev.next = removed.next
	removed.next.prev = removed.prev
	d.length--

	return removed.ele, nil
}

func (d *DoubleCycleLinkedList) IndexOf(ele interface{}) (int, error) {
	cursor := d.headGuard.next
	for i := 0; i < d.length; i++ {
		if reflect.DeepEqual(cursor.ele, ele) {
			return i, nil
		}
		cursor = cursor.next
	}
	return invalidIdx, errors.New(elementNotInListError)
}

func (d *DoubleCycleLinkedList) Clear() {
	d.headGuard.ele = nil
	d.headGuard.next = nil
	d.length = 0
}

func (d *DoubleCycleLinkedList) String() string {
	eles := "["
	cursor := d.headGuard.next
	if cursor.ele != nil {
		for i := 0; i < d.length-1; i++ {
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

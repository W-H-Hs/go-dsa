package list

import (
	"encoding/json"
	"errors"
	"reflect"
)

type singleCycleLinkedListNode struct {
	ele  interface{}
	next *singleCycleLinkedListNode
}

func newSingleCycleLinkedNode(ele interface{}, next *singleCycleLinkedListNode) *singleCycleLinkedListNode {
	return &singleCycleLinkedListNode{
		ele,
		next,
	}
}

type SingleCycleLinkedList struct {
	listCommon
	headGuard *singleCycleLinkedListNode
}

func NewSingleCycleLinkedList() *SingleCycleLinkedList {
	return &SingleCycleLinkedList{
		headGuard: newSingleCycleLinkedNode(nil, nil),
	}
}

func (c *SingleCycleLinkedList) getNodeByIdx(idx int) (*singleCycleLinkedListNode, error) {
	chkErr := c.indexCheck(idx, c.length)
	if chkErr != nil {
		return nil, chkErr
	}

	cursor := c.headGuard.next
	for i := 0; i < idx; i++ {
		cursor = cursor.next
	}
	return cursor, nil
}

func (c *SingleCycleLinkedList) IsContain(ele interface{}) bool {
	_, idxErr := c.IndexOf(ele)
	if idxErr != nil {
		return false
	}
	return true
}

func (c *SingleCycleLinkedList) Append(ele interface{}) {
	_ = c.Insert(c.length, ele)
}

func (c *SingleCycleLinkedList) Get(idx int) (interface{}, error) {
	cursor, nodeErr := c.getNodeByIdx(idx)
	if nodeErr != nil {
		return nil, nodeErr
	}

	return cursor.ele, nil
}

func (c *SingleCycleLinkedList) Set(idx int, ele interface{}) (interface{}, error) {
	cursor, nodeErr := c.getNodeByIdx(idx)
	if nodeErr != nil {
		return nil, nodeErr
	}

	old := cursor.ele
	cursor.ele = ele
	return old, nil
}

func (c *SingleCycleLinkedList) Insert(idx int, ele interface{}) error {
	if c.length == 0 {
		insertedNode := newSingleCycleLinkedNode(ele, nil)
		c.headGuard.next = insertedNode
		insertedNode.next = insertedNode
	} else if idx == 0 {
		last, _ := c.getNodeByIdx(c.length - 1)
		insertedNode := newSingleCycleLinkedNode(ele, c.headGuard.next)
		c.headGuard.next = insertedNode
		last.next = insertedNode
	} else {
		cursor, nodeErr := c.getNodeByIdx(idx - 1)
		if nodeErr != nil {
			return nodeErr
		}
		inserted := newSingleCycleLinkedNode(ele, cursor.next)
		cursor.next = inserted
	}
	c.length++
	return nil
}

func (c *SingleCycleLinkedList) Remove(idx int) (interface{}, error) {
	chkErr := c.indexCheck(idx, c.length)
	if chkErr != nil {
		return nil, chkErr
	}
	var removed *singleCycleLinkedListNode
	if idx == 0 {
		removed = c.headGuard.next
		c.headGuard.next = c.headGuard.next.next
	} else {
		cursor, nodeErr := c.getNodeByIdx(idx - 1)
		if nodeErr != nil {
			return nil, nodeErr
		}
		removed = cursor.next
		cursor.next = cursor.next.next
	}
	c.length--

	return removed.ele, nil
}

func (c *SingleCycleLinkedList) IndexOf(ele interface{}) (int, error) {
	cursor := c.headGuard.next
	for i := 0; i < c.length; i++ {
		if reflect.DeepEqual(cursor.ele, ele) {
			return i, nil
		}
		cursor = cursor.next
	}
	return invalidIdx, errors.New(elementNotInListError)
}

func (c *SingleCycleLinkedList) Clear() {
	c.headGuard.next = nil
	c.length = 0
}

func (c *SingleCycleLinkedList) String() string {
	eles := "["
	cursor := c.headGuard.next
	if cursor.ele != nil {
		for i := 0; i < c.length-1; i++ {
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

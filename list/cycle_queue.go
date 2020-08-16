package list

import (
	"reflect"
	"strconv"
)

type CycleQueue struct {
	store  []interface{}
	front  int
	rear   int
	length int
}

func NewCycleQueue(cap int) *CycleQueue {
	return &CycleQueue{
		store: make([]interface{}, cap),
	}
}

func (c *CycleQueue) Len() int {
	return c.length
}

func (c *CycleQueue) Clear() {
	c.front = 0
	c.rear = 0
}

func (c *CycleQueue) grow(cap int) {
	oldCap := len(c.store)
	if cap > oldCap {
		var newCap int
		if oldCap < 1024 {
			newCap = 2 * oldCap
		} else {
			newCap = oldCap + oldCap/4
		}

		// 根据类型大小来调整newLen
		typeSize := int(reflect.TypeOf(c.store[0]).Size())
		switch {
		case typeSize == 1:
			newCap = getMallocSize(newCap)
		case typeSize == ptrSize:
			tmp := getMallocSize(newCap * ptrSize)
			newCap = tmp / ptrSize
		case isPowOfTwo(typeSize):
			var shift int
			if ptrSize == 8 {
				shift = calcLog(typeSize) & 63
			} else {
				shift = calcLog(typeSize) & 31
			}
			tmp := getMallocSize(newCap << shift)
			newCap = tmp >> shift
		default:
			tmp := getMallocSize(newCap * typeSize)
			newCap = tmp / typeSize
		}

		newArr := make([]interface{}, newCap)
		for i := 0; i <= c.length; i++ {
			newArr[i] = c.store[(i+c.front)%len(c.store)]
		}
		c.store = newArr
		c.front = 0
		c.rear = c.length
	}
}

func (c *CycleQueue) Front() int {
	return c.front
}

func (c *CycleQueue) Rear() int {
	return c.rear
}

func (c *CycleQueue) Store() []interface{} {
	return c.store
}

func (c *CycleQueue) EnQueue(ele interface{}) {
	if c.length+1 == len(c.store) {
		c.grow(c.length + 2)
	}

	c.store[c.rear] = ele
	c.rear = (c.rear + 1) % len(c.store)
	c.length++
}

func (c *CycleQueue) DeQueue() interface{} {
	ele := c.store[c.front]
	c.front = (c.front + 1) % len(c.store)
	c.length--
	return ele
}

func (c *CycleQueue) IsEmpty() bool {
	return c.front == c.rear
}

func (c *CycleQueue) String() string {
	tmp := *c
	e := ""
	for !tmp.IsEmpty() {
		e += strconv.Itoa(tmp.DeQueue().(int))
	}
	return e
}

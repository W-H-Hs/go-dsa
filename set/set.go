package set

import (
	"fmt"
	"go-dsa/common"
)

type Visitor interface {
	GetStop() bool
	Visit(interface{}) bool
}

type Printer struct {
	stop bool
}

func (p Printer) GetStop() bool {
	return p.stop
}

func (p Printer) Visit(ele interface{}) bool {
	fmt.Println(ele)
	return false
}

type Set interface {
	Size() int
	IsEmpty() bool
	Clear()
	IsContain(common.ComparableElement) bool
	Add(interface{}) (interface{}, error)
	Remove(interface{}) (interface{}, error)
	Traversal(Visitor)
}

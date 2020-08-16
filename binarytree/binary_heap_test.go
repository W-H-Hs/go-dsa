package binarytree

import (
	"fmt"
	"testing"
)

var bh = NewBinaryHeap()

func TestBinaryHeap_Add(t *testing.T) {
	bh.Add(68)
	bh.Add(72)
	bh.Add(43)
	bh.Add(50)
	bh.Add(38)
	fmt.Println(bh.elements)
}

func TestBinaryHeap_Remove(t *testing.T) {
	arr := []int{27, 15, 34, 20, 45, 55, 95, 81, 49, 4, 29, 86, 98}
	for _, value := range arr {
		bh.Add(value)
	}
	fmt.Println(bh.elements)
	bh.Remove()
	bh.Remove()
	bh.Remove()
	bh.Remove()
	bh.Remove()
	bh.Remove()
	fmt.Println(bh.elements)
}

func TestBinaryHeap_Replace(t *testing.T) {
	bh.Add(68)
	bh.Add(72)
	bh.Add(43)
	bh.Add(50)
	bh.Add(38)
	bh.Add(10)
	bh.Add(90)
	bh.Add(65)
	bh.Replace(70)
	fmt.Println(bh.elements)
}

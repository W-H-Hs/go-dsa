package sorting

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	var a []int
	for i := 1; i < 10; i++ {
		a = append(a, i)
	}
	fmt.Println(BinarySearch(1, a))
	fmt.Println(BinarySearch(2, a))
	fmt.Println(BinarySearch(3, a))
	fmt.Println(BinarySearch(4, a))
	fmt.Println(BinarySearch(5, a))
	fmt.Println(BinarySearch(6, a))
	fmt.Println(BinarySearch(7, a))
	fmt.Println(BinarySearch(8, a))
	fmt.Println(BinarySearch(9, a))
}

func TestBinarySearchOfInsertIndex(t *testing.T) {
	a := []int{1, 3, 3, 5, 7, 9}
	fmt.Println(BinarySearchOfInsertIndex(1, a))
	fmt.Println(BinarySearchOfInsertIndex(2, a))
	fmt.Println(BinarySearchOfInsertIndex(3, a))
	fmt.Println(BinarySearchOfInsertIndex(4, a))
	fmt.Println(BinarySearchOfInsertIndex(5, a))
	fmt.Println(BinarySearchOfInsertIndex(6, a))
	fmt.Println(BinarySearchOfInsertIndex(7, a))
	fmt.Println(BinarySearchOfInsertIndex(8, a))
	fmt.Println(BinarySearchOfInsertIndex(9, a))
}

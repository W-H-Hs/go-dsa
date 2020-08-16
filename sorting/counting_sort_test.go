package sorting

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCountingSort(t *testing.T) {
	testArr = []int{-10, 4, 2, 6, -10, -8, -8, -9, -7, -5, -1, 2, 3, 1, 4}
	fmt.Println(testArr)
	oldArr := make([]int, len(testArr))
	reflect.Copy(reflect.ValueOf(oldArr), reflect.ValueOf(testArr))
	CountingSort(testArr, -10, 6)
	fmt.Println(testArr)
	fmt.Println(isSortedArray(oldArr, testArr))
}

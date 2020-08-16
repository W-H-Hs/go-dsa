package sorting

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRadixSort(t *testing.T) {
	fmt.Println(testArr)
	oldArr := make([]int, len(testArr))
	reflect.Copy(reflect.ValueOf(oldArr), reflect.ValueOf(testArr))
	RadixSort(testArr)
	fmt.Println(testArr)
	fmt.Println(isSortedArray(oldArr, testArr))
}

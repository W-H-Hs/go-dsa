package sorting

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	oldArr := make([]int, len(testArr))
	reflect.Copy(reflect.ValueOf(oldArr), reflect.ValueOf(testArr))
	fmt.Println(testArr)
	MergeSort(testArr)
	fmt.Println(testArr)
	fmt.Println(isSortedArray(oldArr, testArr))
}

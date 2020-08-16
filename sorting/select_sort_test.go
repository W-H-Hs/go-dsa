package sorting

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSelectSort(t *testing.T) {
	oldArr := make([]int, len(testArr))
	reflect.Copy(reflect.ValueOf(oldArr), reflect.ValueOf(testArr))
	SelectSort(testArr)
	fmt.Println(testArr)
	fmt.Println(isSortedArray(oldArr, testArr))
}

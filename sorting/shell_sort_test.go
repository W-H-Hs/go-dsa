package sorting

import (
	"fmt"
	"reflect"
	"testing"
)

func TestShellSort(t *testing.T) {
	//testArr = []int{44, 46, 7, 86, 83, 72, 21, 28, 91, 46}
	oldArr := make([]int, len(testArr))
	reflect.Copy(reflect.ValueOf(oldArr), reflect.ValueOf(testArr))
	fmt.Println(testArr)
	ShellSort(testArr)
	fmt.Println(testArr)
	fmt.Println(isSortedArray(oldArr, testArr))
}

package sorting

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{19, 9, 16, 7, 26, 25, 11, 33, 41, 48, 52}
	oldArr := make([]int, len(arr))
	reflect.Copy(reflect.ValueOf(oldArr), reflect.ValueOf(arr))
	// 切片在函数参数的传递上为引用传递，数组为值传递。
	BubbleSort(arr)
	fmt.Println(arr)
	fmt.Println(isSortedArray(oldArr, arr))
}

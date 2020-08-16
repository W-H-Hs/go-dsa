package list

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"reflect"
)

const (
	defaultCap = 5

	// 指针大小。
	ptrSize = 4 << (^uintptr(0) >> 63)
	// 最大小对象的大小，参见golang内存管理。
	maxSmallSize = 32678
	// 页大小，参见golang内存管理。
	pageSize = 8192
)

// 分配空间大小，参见golang内存管理或/$goroot$/src/runtime/sizeclass.go。
var mallocSize = []int{
	0, 8, 16, 32, 48, 64, 80, 96,
	112, 128, 144, 160, 176, 192, 208, 224,
	240, 256, 288, 320, 352, 384, 416, 448,
	480, 512, 576, 640, 704, 768, 896, 1024,
	1152, 1280, 1408, 1536, 1792, 2048, 2304,
	2688, 3072, 3200, 3456, 4096, 4864, 5376,
	6144, 6528, 6784, 6912, 8192, 9472, 9728,
	10240, 10880, 12288, 13568, 14336, 16384,
	18432, 19072, 20480, 21760, 24576, 27264,
	28672, 32768,
}

// 将n提升为最小大于n的a的倍数
func round(n, a int) int {
	return (n + a - 1) &^ (a - 1)
}

// 传入申请空间大小，返回实际分配空间大小
func getMallocSize(applySize int) int {
	if applySize < maxSmallSize {
		for i := 0; i < len(mallocSize)-1; i++ {
			if mallocSize[i] < applySize && applySize <= mallocSize[i+1] {
				return mallocSize[i+1]
			}
		}
	}
	return round(applySize, pageSize)
}

func isPowOfTwo(num int) bool {
	return num&(num-1) == 0
}

// 计算log2(num)。
func calcLog(num int) int {
	return int(math.Log(float64(num)) / math.Log(2))
}

type ArrayList struct {
	listCommon
	elements []interface{}
}

func NewArrayList() List {
	return &ArrayList{
		elements: make([]interface{}, defaultCap),
	}
}

func NewArrayListWithCapacity(capacity int) List {
	return &ArrayList{
		elements: make([]interface{}, capacity),
	}
}

func (a *ArrayList) IsContain(ele interface{}) bool {
	_, idxErr := a.IndexOf(ele)
	if idxErr != nil {
		return false
	}
	return true
}

func (a *ArrayList) grow(cap int) {
	oldCap := len(a.elements)
	if cap > oldCap {
		var newCap int
		if oldCap < 1024 {
			newCap = 2 * oldCap
		} else {
			newCap = oldCap + oldCap/4
		}

		// 根据类型大小来调整newLen
		typeSize := int(reflect.TypeOf(a.elements[0]).Size())
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
		copy(newArr, a.elements)
		a.elements = newArr
	}
}

func (a *ArrayList) Append(ele interface{}) {
	a.grow(a.length + 1)
	a.elements[a.length] = ele
	a.length++
}

func (a *ArrayList) Get(idx int) (interface{}, error) {
	chkErr := a.indexCheck(idx, a.length)
	if chkErr != nil {
		return nil, chkErr
	}
	return a.elements[idx], nil
}

func (a *ArrayList) Set(idx int, ele interface{}) (interface{}, error) {
	chkErr := a.indexCheck(idx, a.length)
	if chkErr != nil {
		return nil, chkErr
	}
	old := a.elements[idx]
	a.elements[idx] = ele
	return old, nil
}

func (a *ArrayList) Insert(idx int, ele interface{}) error {
	chkErr := a.indexCheck(idx, a.length+1)
	if chkErr != nil {
		return chkErr
	}
	a.grow(a.length + 1)
	for i := a.length; i > idx; i-- {
		a.elements[i] = a.elements[i-1]
	}
	a.elements[idx] = ele
	a.length++
	return nil
}

func (a *ArrayList) Remove(idx int) (interface{}, error) {
	chkErr := a.indexCheck(idx, a.length)
	if chkErr != nil {
		return nil, chkErr
	}
	rmEle := a.elements[idx]
	for ; idx < a.length-1; idx++ {
		a.elements[idx] = a.elements[idx+1]
	}
	a.length--
	a.elements[a.length] = nil
	return rmEle, nil
}

func (a *ArrayList) IndexOf(ele interface{}) (int, error) {
	fmt.Println(&a.listCommon)
	for i := 0; i < a.length; i++ {
		if reflect.DeepEqual(ele, a.elements[i]) {
			return i, nil
		}
	}
	return invalidIdx, errors.New(elementNotInListError)
}

func (a *ArrayList) Clear() {
	for i := 0; i < len(a.elements); i++ {
		a.elements[i] = nil
	}
	a.length = 0
}

func (a *ArrayList) String() string {
	eles, _ := json.Marshal(a.elements[:a.length])
	return string(eles)
}

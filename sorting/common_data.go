package sorting

import (
	"math/rand"
	"reflect"
	"time"
)

const invalidIndex = -1

var testArr []int

func init() {
	rand.Seed(time.Now().Unix())
	for i := 0; i < 50; i++ {
		testArr = append(testArr, rand.Intn(1000000))
	}
}

func isSortedArray(oldArr, arr []int) bool {
	m1 := make(map[int][]int)
	m2 := make(map[int][]int)
	for i := 0; i < len(oldArr); i++ {
		m1[oldArr[i]] = nil
		m2[oldArr[i]] = nil
	}
	if !reflect.DeepEqual(m1, m2) {
		return false
	}

	isSorted := true
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			isSorted = false
		}
	}

	return isSorted
}

package sorting

import (
	"fmt"
	"testing"
)

func TestHeapSort(t *testing.T) {
	HeapSort(testArr)
	fmt.Println(testArr)
}

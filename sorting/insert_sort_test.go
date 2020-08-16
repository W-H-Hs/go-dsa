package sorting

import (
	"fmt"
	"testing"
)

func TestInsertSort(t *testing.T) {
	InsertSort(testArr)
	fmt.Println(testArr)
}

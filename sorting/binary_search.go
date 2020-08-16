package sorting

// 寻找指定元素的索引。
func BinarySearch(ele int, arr []int) int {
	if arr == nil || len(arr) == 0 {
		return invalidIndex
	}

	begin := 0
	end := len(arr)
	for {
		// 例如begin为3，end为6时，这个区间的中点索引为4，而
		// (end+begin)>>1就是4。
		midIdx := (end + begin) >> 1
		if ele == arr[midIdx] {
			return midIdx
		} else if ele > arr[midIdx] {
			begin = midIdx + 1
		} else {
			end = midIdx
		}
	}
}

// 寻找指定元素的插入位置。arr[i]<ele<arr[i+1]时返回i+1，即第一个
// 大于ele的元素索引。
func BinarySearchOfInsertIndex(ele int, arr []int) int {
	if arr == nil || len(arr) == 0 {
		return invalidIndex
	}

	begin := 0
	end := len(arr)
	for begin < end {
		midIdx := (begin + end) >> 1
		if ele >= arr[midIdx] {
			begin = midIdx + 1
		} else {
			end = midIdx
		}
	}
	return end
}

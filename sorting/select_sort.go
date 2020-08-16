package sorting

// 交换次数远远少于冒泡排序，性能优于冒泡排序。

// 遍历数组找出最大的元素，将其与数组最后一个元素交换位置。
// 将将数组除开最后一个（最大元素）进行递归。
// 时间复杂度：
//    最好、最坏、平均均为O(n^2)（一次递归执行中遍历arr需要n，总共需要n-2次递归执行）。
// 空间复杂度：
//    O(1)。
func SelectSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}

	maxEleIdx := 0
	for i := 1; i < len(arr); i++ {
		// 大于等于保证稳定性。
		// 只是大于则无稳定性。
		if arr[i] >= arr[maxEleIdx] {
			maxEleIdx = i
		}
	}
	lastEleIdx := len(arr) - 1
	if maxEleIdx != lastEleIdx {
		arr[maxEleIdx] ^= arr[lastEleIdx]
		arr[lastEleIdx] ^= arr[maxEleIdx]
		arr[maxEleIdx] ^= arr[lastEleIdx]
	}

	SelectSort(arr[:len(arr)-1])
}

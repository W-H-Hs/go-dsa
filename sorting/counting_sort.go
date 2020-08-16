package sorting

func CountingSort(arr []int, minBoundary, maxBoundary int) {
	// 计数数组。
	counting := make([]int, maxBoundary-minBoundary+1)
	// 统计待排序数组中元素的出现次数。
	for _, ele := range arr {
		// minBoundary看做偏移量。
		counting[ele-minBoundary]++
	}

	arrIdx := 0
	for i := 0; i < len(counting); i++ {
		for counting[i] != 0 {
			arr[arrIdx] = i + minBoundary
			counting[i]--
			arrIdx++
		}
	}
}

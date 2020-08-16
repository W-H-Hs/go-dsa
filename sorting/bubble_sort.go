package sorting

// 两根指针一前一后，如果前面指针指向的元素小于后面的，则交换两个指针指向的元素。
// 遍历到最后一个元素时，将数组[:n]（除开最后一个元素）作为参数进行递归。
// 时间复杂度：
//    最好情况O(n)，即arr已经是有序的O(n)时间花在一次遍历上。
//    最坏情况O(n^2)，即arr所有元素均需交换（一次递归执行中遍历arr需要n，总共需要n-2次递归执行）。
// 	  平均O(n^2)。
// 空间复杂度：
//    O(1)。
func BubbleSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		return
	}

	var first, second, lastExchange int
	for i := 0; i < len(arr)-1; i++ {
		first = arr[i+1]
		second = arr[i]
		// 小于等于保证稳定性。
		// 只是小于则无稳定性。
		if first < second {
			arr[i] ^= arr[i+1]
			arr[i+1] ^= arr[i]
			arr[i] ^= arr[i+1]
			lastExchange = i + 1
		}
	}

	// 当一次BubbleSort()函数执行完时，arr最后一个元素一定arr中最大的元素。
	// 如果arr从某个元素起，之后的所有元素都是有序的，那么该部分就不用再递
	// 归。该元素的索引即lastExchange - 1。
	BubbleSort(arr[:lastExchange])
}

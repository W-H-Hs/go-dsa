package sorting

// TODO: 算法描述，时间复杂度分析。

func sortOfMerge(arr []int, begin, end int) {
	midIdx := (begin + end) >> 1

	// divide操作。

	if end-begin < 2 {
		return
	}

	// midIdx计算结果已经是子数组最后一个元素的索引加一。
	// 比如数组arr的元素索引为0-5，左半部分的元素索引为
	// [0, 3)，右半部分的元素索引为[3-6)，midIdx计算为
	// (0+5)>>1=3。
	sortOfMerge(arr, begin, midIdx)
	sortOfMerge(arr, midIdx, end)

	// merge操作。

	//备份左半部分。
	tmpArr := make([]int, midIdx-begin)
	for i := begin; i < midIdx; i++ {
		tmpArr[i-begin] = arr[i]
	}

	leftPointer := 0
	rightPointer := midIdx
	insertPointer := begin
	for leftPointer < len(tmpArr) {
		// leftPointer==len(tmpArr)，时，左半部分元素（tmpArr中元素）先取完，则子数组的排序已完成，因为右
		// 半部分元素是已经排好序且存在于arr右半部分，直接跳出循环即可。
		if rightPointer < end && tmpArr[leftPointer] > arr[rightPointer] {
			arr[insertPointer] = arr[rightPointer]
			rightPointer++
		} else {
			// rightPointer==end时，右半部分元素先取完，则将左半部分（tmpArr中元素）挪到arr右半部分，此时
			// 只会执行该else部分代码，即直接将tmpArr中剩下的元素挪到arr右半部分剩下的位置中。
			arr[insertPointer] = tmpArr[leftPointer]
			leftPointer++
		}
		insertPointer++
	}
}

func MergeSort(arr []int) {
	if arr == nil {
		return
	}

	sortOfMerge(arr, 0, len(arr))
}

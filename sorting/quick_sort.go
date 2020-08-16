package sorting

// TODO: 算法描述，时间复杂度分析。

func sortOfQuick(arr []int, begin, end int) {
	if end-begin < 2 {
		return
	}

	oldBegin := begin
	oldEnd := end

	pivotEle := arr[begin]
	isBeginShift := false
	for {
		if isBeginShift {
			if arr[begin] > pivotEle {
				arr[end-1] = arr[begin]
				end--
				isBeginShift = false
			} else {
				begin++
			}
		} else {
			if arr[end-1] < pivotEle {
				arr[begin] = arr[end-1]
				begin++
				isBeginShift = true
			} else {
				end--
			}
		}
		if begin == end-1 {
			arr[begin] = pivotEle
			break
		}
	}

	sortOfQuick(arr, oldBegin, begin)
	sortOfQuick(arr, begin+1, oldEnd)
}

func QuickSort(arr []int) {
	if arr == nil {
		return
	}

	sortOfQuick(arr, 0, len(arr))
}

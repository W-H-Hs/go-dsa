package sorting

func ShellSort(arr []int) {
	// 增量递减中的增量。
	k := 1
	length := len(arr)
	for {
		// 步长，希尔给出的公式，n/(2^k)，也是增量递减中的递减。
		step := length >> k
		if step == 0 {
			break
		}
		for col := 0; col < step; col++ {
			// 对每一列进行插入排序。
			// 可以看做模step剩余类，每一列的所有元素就是[a](0<a<stop)。
			sortedIdx := col
			for sortedIdx < length-step {
				// 插入排序。
				insertEle := arr[sortedIdx+step]
				// 初始化插入位置为起始位置，如果插入元素小于起始位置元素，那插入位置就是起始位置。。
				insertIdx := col
				if insertEle >= arr[col] {
					// 如果插入元素大于等于起始位置元素，则需要找到插入位置。等于是为了保证稳定性。
					for idx := sortedIdx + step; idx > col; idx -= step {
						// 从已排序元素的最后一个开始往前遍历，寻找插入元素的位置，当出现第一个小于插入元素时，
						// 其后一个位置即为插入位置。
						// idx+step为插入位置的索引。
						// 等于保证稳定性。
						if insertEle >= arr[idx-step] {
							// 跳出当前for循环，因为插入位置已经找到，没必要再往前遍历。
							insertIdx = idx
							break
						}
					}
				}
				for i := sortedIdx; i >= insertIdx; i -= step {
					// 后挪元素，为插入元素挪出位置。
					arr[i+step] = arr[i]
				}
				// 插入元素。
				arr[insertIdx] = insertEle
				// 更新已排序元素的位置。
				sortedIdx += step
			}
		}
		k++
	}
}

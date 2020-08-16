package sorting

// 1.分为两部分，一部分是已经排好序的，另一部分是未排序的。
// 2.从未排序部分中选出一个元素，遍历已排序部分，插入合适的位置。
// 类比于打扑克牌，手里面抓的牌是已经排好序的，洗好的牌是未排序的，
// 从洗好的牌里面拿出一张来放进手里抓的牌里面。
func InsertSort(arr []int) {
	sortedIdx := 0

	for sortedIdx < len(arr)-1 {
		insertedEle := arr[sortedIdx+1]
		// 通过二分搜索法查找出插入位置。
		insertedIdx := BinarySearchOfInsertIndex(insertedEle, arr[:sortedIdx+1])

		// 后挪元素，为插入元素腾出位置。
		for i := sortedIdx; i >= insertedIdx; i-- {
			arr[i+1] = arr[i]
		}
		// 插入元素。
		arr[insertedIdx] = insertedEle
		// 更新已排序元素索引。
		sortedIdx++
	}
}

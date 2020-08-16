package sorting

import (
	"errors"
)

type binaryHeap struct {
	elements []int
	size     int
}

func newBinaryHeap(arr []int) *binaryHeap {
	bh := &binaryHeap{
		elements: arr,
		size:     len(arr),
	}

	// 原地建堆，自下而上的下滤。
	for i := (len(arr) >> 1) - 1; i >= 0; i-- {
		_ = bh.siftDown(i)
	}

	return bh
}

func (b *binaryHeap) getIndexOfLeftChild(idx int) (int, error) {
	if (idx<<1)+1 > b.size-1 {
		return invalidIndex, errors.New("this element doesn't have left child")
	}

	return (idx << 1) + 1, nil
}

func (b *binaryHeap) getIndexOfRightChild(idx int) (int, error) {
	if (idx<<1)+2 > b.size-1 {
		return invalidIndex, errors.New("this element doesn't have right child")
	}

	return (idx << 1) + 2, nil
}

func (b *binaryHeap) getIndexOfParent(idx int) (int, error) {
	if idx == 0 {
		return invalidIndex, errors.New("root don't have parent")
	}

	return (idx - 1) >> 1, nil
}

func (b *binaryHeap) siftUp(index int) error {
	ele := b.elements[index]

	for index > 0 {
		parentIdx, parentErr := b.getIndexOfParent(index)
		if parentErr != nil {
			return parentErr
		}
		parent := b.elements[parentIdx]
		if ele > parent {
			b.elements[index] = parent
			b.elements[parentIdx] = ele
		}
		index = parentIdx
	}

	return nil
}

func (b *binaryHeap) add(element int) (int, error) {
	b.elements = append(b.elements, element)
	b.size++
	err := b.siftUp(b.size - 1)
	if err != nil {
		return invalidIndex, err
	}

	return element, nil
}

func (b *binaryHeap) siftDown(idx int) error {
	ele := b.elements[idx]

	firstLeafIdx := b.size >> 1
	for idx < firstLeafIdx {
		rightChildIdx, rightChildIdxErr := b.getIndexOfRightChild(idx)
		if rightChildIdxErr != nil {
			leftChildIdx, _ := b.getIndexOfLeftChild(idx)
			leftChild := b.elements[leftChildIdx]
			if ele < leftChild {
				b.elements[idx] = leftChild
				b.elements[leftChildIdx] = ele
				idx = leftChildIdx
			} else {
				return nil
			}
		} else {
			rightChild := b.elements[rightChildIdx]
			leftChildIdx, _ := b.getIndexOfLeftChild(idx)
			leftChild := b.elements[leftChildIdx]
			if leftChild > rightChild {
				if ele < leftChild {
					b.elements[idx] = leftChild
					b.elements[leftChildIdx] = ele
					idx = leftChildIdx
				} else {
					return nil
				}
			} else {
				if ele < rightChild {
					b.elements[idx] = rightChild
					b.elements[rightChildIdx] = ele
					idx = rightChildIdx
				} else {
					return nil
				}
			}
		}
	}

	return nil
}

// 对选择排序的一种优化，选最大值的时候不用遍历arr，而是使用堆。
func HeapSort(arr []int) {
	// 最大堆。
	heap := newBinaryHeap(arr)

	for heap.size > 1 {
		// 交换堆顶元素和尾部元素，最大堆的堆顶元素为最大值。
		arr[0] ^= arr[heap.size-1]
		arr[heap.size-1] ^= arr[0]
		arr[0] ^= arr[heap.size-1]

		//
		heap.size--

		// 对根元素执行下滤。
		_ = heap.siftDown(0)
	}
}

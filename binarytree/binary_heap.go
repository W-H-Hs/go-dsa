package binarytree

import (
	"errors"
	"go-dsa/list"
)

const (
	defaultCapacity = 10
	invalidIndex    = -1
)

// 获取最大值：O(1)
// 删除最大值：O(logn)
// 添加元素：O(logn)
type BinaryHeap struct {
	elements *list.ArrayList
	size     int
}

func NewBinaryHeap() *BinaryHeap {
	return &BinaryHeap{
		elements: list.NewArrayListWithCapacity(defaultCapacity).(*list.ArrayList),
	}
}

func (b *BinaryHeap) Size() int {
	return b.size
}

func (b *BinaryHeap) IsEmpty() bool {
	return b.size == 0
}

func (b *BinaryHeap) Clear() {
	b.size = 0
}

func (b *BinaryHeap) GetRoot() (int, error) {
	if b.size == 0 {
		return 0, errors.New("heap is empty")
	}

	root, err := b.elements.Get(0)
	return root.(int), err
}

// 左子节点索引计算表达式：2 * index + 1。
func (b *BinaryHeap) getIndexOfLeftChild(idx int) (int, error) {
	if (idx<<1)+1 > b.size-1 {
		return invalidIndex, errors.New("this element doesn't have left child")
	}

	return (idx << 1) + 1, nil
}

// 右子节点索引计算表达式：2 * index + 2。
func (b *BinaryHeap) getIndexOfRightChild(idx int) (int, error) {
	if (idx<<1)+2 > b.size-1 {
		return invalidIndex, errors.New("this element doesn't have right child")
	}

	return (idx << 1) + 2, nil
}

// 父节点索引结算表达式：2 * (index - 1)。
func (b *BinaryHeap) getIndexOfParent(idx int) (int, error) {
	if idx == 0 {
		return invalidIndex, errors.New("root don't have parent")
	}

	return (idx - 1) >> 1, nil
}

// 让index位置的元素上滤，如果index元素的值大于父节点，则交换
// 该元素与父节点的位置。
func (b *BinaryHeap) siftUp(index int) error {
	ele, err := b.elements.Get(index)
	if err != nil {
		return err
	}

	for index > 0 {
		parentIdx, parentErr := b.getIndexOfParent(index)
		if parentErr != nil {
			return parentErr
		}
		parent, _ := b.elements.Get(parentIdx)
		if ele.(int) > parent.(int) {
			// 如果ele的值大于parent的值，则将ele所在的位置
			// 设置为parent。
			_, _ = b.elements.Set(index, parent)
			_, _ = b.elements.Set(parentIdx, ele)
		}
		index = parentIdx
	}

	return nil
}

// 插入新元素时首先将其放在最末尾的位置，然后进行上滤。
func (b *BinaryHeap) Add(element int) (int, error) {
	b.elements.Append(element)
	b.size++
	err := b.siftUp(b.size - 1)
	if err != nil {
		return invalidIndex, err
	}

	return element, nil
}

// 让index位置的元素下滤，如果index元素存在左子节点且其值小于左子节点
// 元素的值则交换该元素与左子节点，否则查看右子节点，如果存在且小于，则
// 交换。
func (b *BinaryHeap) siftDown(idx int) error {
	ele, err := b.elements.Get(idx)
	if err != nil {
		return err
	}

	// 只有当index对应的元素为非叶节点时，才会执行下滤。
	// 即index的值必须小于非叶节点的个数即第一个叶节点
	// 的索引。
	// 非叶节点的个数为floor(n/2)。
	// 叶节点的个数为ceiling(n/2)。
	firstLeafIdx := b.size >> 1
	for idx < firstLeafIdx {
		rightChildIdx, rightChildIdxErr := b.getIndexOfRightChild(idx)
		if rightChildIdxErr != nil {
			// index对应的节点只有左子节点。
			leftChildIdx, _ := b.getIndexOfLeftChild(idx)
			leftChild, _ := b.elements.Get(leftChildIdx)
			if ele.(int) < leftChild.(int) {
				_, _ = b.elements.Set(idx, leftChild)
				_, _ = b.elements.Set(leftChildIdx, ele)
				idx = leftChildIdx
			} else {
				// index对应的节点已经大于它的左子节点，下滤结束。
				return nil
			}
		} else {
			// index对应的节点有左右子节点。
			rightChild, _ := b.elements.Get(rightChildIdx)
			leftChildIdx, _ := b.getIndexOfLeftChild(idx)
			leftChild, _ := b.elements.Get(leftChildIdx)
			if leftChild.(int) > rightChild.(int) {
				// 左子节点的元素值大于右子节点。
				if ele.(int) < leftChild.(int) {
					// index对应的元素值小于左子节点的值。
					_, _ = b.elements.Set(idx, leftChild)
					_, _ = b.elements.Set(leftChildIdx, ele)
					idx = leftChildIdx
				} else {
					// index对应的元素值大于左右子节点，下滤结束。
					return nil
				}
			} else {
				if ele.(int) < rightChild.(int) {
					// index对应的元素值小于右子节点的值。
					_, _ = b.elements.Set(idx, rightChild)
					_, _ = b.elements.Set(rightChildIdx, ele)
					idx = rightChildIdx
				} else {
					// index对应的元素值大于左右子节点，下滤结束。
					return nil
				}
			}
		}
	}

	return nil
}

// 删除根元素，用数组最后一个元素（堆的最后一个元素）替换根元素，
// 删除最后一个元素，然后对根元素进行下滤。
func (b *BinaryHeap) Remove() error {
	lastElement, err := b.elements.Get(b.size - 1)
	if err != nil {
		return err
	}

	_, err = b.elements.Set(0, lastElement)
	if err != nil {
		return err
	}

	_, err = b.elements.Remove(b.size - 1)
	if err != nil {
		return err
	}
	b.size--

	err = b.siftDown(0)
	if err != nil {
		return err
	}

	return nil
}

func (b *BinaryHeap) Replace(element int) (int, error) {
	if b.size == 0 {
		b.elements.Append(element)
		b.size++
		return invalidIndex, nil
	} else {
		root, _ := b.elements.Set(0, element)
		err := b.siftDown(0)
		if err != nil {
			return invalidIndex, err
		}

		return root.(int), nil
	}
}

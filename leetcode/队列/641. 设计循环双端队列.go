package 队列

// https://leetcode-cn.com/problems/design-circular-deque/

//type MyCircularDeque struct {
//	store    []int
//	front    int
//	rear     int
//	length   int
//	capacity int
//}
//
///** Initialize your data structure here. Set the size of the deque to be k. */
//func Constructor(k int) MyCircularDeque {
//	return MyCircularDeque{
//		store:    make([]int, k+1),
//		capacity: k,
//	}
//}
//
//func (this *MyCircularDeque) canInsert() bool {
//	if this.length+1 == len(this.store) {
//		return false
//	}
//	return true
//}
//
//func (this *MyCircularDeque) grow() {
//	capacity := len(this.store)
//	newArr := make([]int, 2*capacity)
//	for i := 0; i <= this.length; i++ {
//		newArr[i] = this.store[(i+this.front)%capacity]
//	}
//	this.store = newArr
//	this.front = 0
//	this.rear = this.length
//}
//
//func (this *MyCircularDeque) insert(fn func() bool) func(int) bool {
//	return func(value int) bool {
//		if value < 1 || value > 1000 || this.length == 1000 {
//			return false
//		}
//
//		if this.length == 0 {
//			if len(this.store) == 0 {
//				this.store = make([]int, 5)
//			}
//
//			this.store[0] = value
//			this.rear++
//			this.length++
//
//			return true
//		}
//
//		if this.canInsert() {
//			return fn()
//		} else {
//			this.grow()
//			return fn()
//		}
//	}
//}
//
///** Adds an item at the front of Deque. Return true if the operation is successful. */
//func (this *MyCircularDeque) InsertFront(value int) bool {
//	return this.insert(func() bool {
//		this.front--
//		if this.front < 0 {
//			this.front = len(this.store) + this.front
//		}
//		this.store[this.front] = value
//		this.length++
//		return true
//	})(value)
//}
//
///** Adds an item at the rear of Deque. Return true if the operation is successful. */
//func (this *MyCircularDeque) InsertLast(value int) bool {
//	return this.insert(func() bool {
//		this.store[this.rear] = value
//		this.rear = (this.rear + 1) % len(this.store)
//		this.length++
//		return true
//	})(value)
//}
//
///** Deletes an item from the front of Deque. Return true if the operation is successful. */
//func (this *MyCircularDeque) DeleteFront() bool {
//	capacity := len(this.store)
//	if capacity == 0 {
//		return false
//	}
//
//	this.front = (this.front + 1) % capacity
//	this.length--
//	return true
//}
//
///** Deletes an item from the rear of Deque. Return true if the operation is successful. */
//func (this *MyCircularDeque) DeleteLast() bool {
//	capacity := len(this.store)
//	if capacity == 0 {
//		return false
//	}
//
//	this.rear = (this.rear - 1) % capacity
//	this.length--
//	return true
//}
//
///** Get the front item from the deque. */
//func (this *MyCircularDeque) GetFront() int {
//	if this.length == 0 {
//		return -1
//	}
//
//	return this.front
//}
//
///** Get the rear item from the deque. */
//func (this *MyCircularDeque) GetRear() int {
//	if this.length == 0 {
//		return -1
//	}
//
//	return this.rear
//}
//
///** Checks whether the circular deque is empty or not. */
//func (this *MyCircularDeque) IsEmpty() bool {
//	return this.length == 0
//}
//
/////** Checks whether the circular deque is full or not. */
//func (this *MyCircularDeque) IsFull() bool {
//	return this.length == this.capacity
//}

type MyCircularDeque struct {
	store    []int
	front    int
	rear     int
	length   int
	capacity int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		store:    make([]int, k+1),
		capacity: k,
	}
}

func (this *MyCircularDeque) insert(fn func() bool) func(int) bool {
	return func(value int) bool {
		if this.length == this.capacity {
			return false
		}

		if this.length == 0 {
			this.store[this.front] = value
			this.rear++
			this.length++

			return true
		}

		return fn()
	}
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	return this.insert(func() bool {
		this.front--
		if this.front < 0 {
			this.front = this.capacity
		}
		this.store[this.front] = value
		this.length++
		return true
	})(value)
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	return this.insert(func() bool {
		this.store[this.rear] = value
		this.rear = (this.rear + 1) % (this.capacity + 1)
		this.length++
		return true
	})(value)
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.length == 0 {
		return false
	}

	this.front = (this.front + 1) % (this.capacity + 1)
	this.length--
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.length == 0 {
		return false
	}

	this.rear--
	if this.rear < 0 {
		this.rear = this.capacity
	}
	this.length--
	return true
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if this.length == 0 {
		return -1
	}

	return this.store[this.front]
}

/** Get the rear item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if this.length == 0 {
		return -1
	}

	if this.rear-1<0 {
		return this.store[this.capacity]
	}
	return this.store[this.rear-1]
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	return this.length == 0
}

///** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	return this.length == this.capacity
}

/**
* Your MyCircularDeque object will be instantiated and called as such:
* obj := Constructor(k);
* param_1 := obj.InsertFront(value);
* param_2 := obj.InsertLast(value);
* param_3 := obj.DeleteFront();
* param_4 := obj.DeleteLast();
* param_5 := obj.GetFront();
* param_6 := obj.GetRear();
* param_7 := obj.IsEmpty();
* param_8 := obj.IsFull();
 */

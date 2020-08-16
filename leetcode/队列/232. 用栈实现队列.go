package 队列

// https://leetcode-cn.com/problems/implement-queue-using-stacks/

//type stack struct {
//	array []int
//}
//
//func (s *stack) push(ele int) {
//	s.array = append(s.array, ele)
//}
//func (s *stack) top() int {
//	return s.array[s.len()-1]
//}
//func (s *stack) pop() int {
//	ele := s.array[len(s.array)-1]
//	newArray := make([]int, len(s.array)-1)
//	copy(newArray, s.array[:len(s.array)-1])
//	s.array = newArray
//	return ele
//}
//func (s *stack) len() int {
//	return len(s.array)
//}
//
//type MyQueue struct {
//	tmp   stack
//	store stack
//}
//
///** Initialize your data structure here. */
//func Constructor() MyQueue {
//	return MyQueue{}
//}
//
///** Push element x to the back of queue. */
//func (this *MyQueue) Push(x int) {
//	this.store.push(x)
//}
//
//func (this *MyQueue) reverse(isFromStore bool) {
//	if isFromStore {
//		length := this.store.len()
//		for i := 0; i < length; i++ {
//			this.tmp.push(this.store.pop())
//		}
//	} else {
//		length := this.tmp.len()
//		for i := 0; i < length; i++ {
//			this.store.push(this.tmp.pop())
//		}
//	}
//}
//
///** Removes the element from in front of queue and returns that element. */
//func (this *MyQueue) Pop() int {
//	this.reverse(true)
//	ele := this.tmp.pop()
//	this.reverse(false)
//	return ele
//}
//
///** Get the front element. */
//func (this *MyQueue) Peek() int {
//	this.reverse(true)
//	ele := this.tmp.top()
//	this.reverse(false)
//	return ele
//}
//
///** Returns whether the queue is empty. */
//func (this *MyQueue) Empty() bool {
//	return this.store.len() == 0
//}
//
///**
// * Your MyQueue object will be instantiated and called as such:
// * obj := Constructor();
// * obj.Push(x);
// * param_2 := obj.Pop();
// * param_3 := obj.Peek();
// * param_4 := obj.Empty();
// */

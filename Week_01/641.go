package main

/*
641

| | | ||||||
*/
type MyCircularDeque struct {
	q     []int
	front int
	last  int
	cap   int
	size  int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	deque := MyCircularDeque{}
	deque.front = -1
	deque.last = -1
	deque.cap = k
	deque.size = 0
	deque.q = make([]int, k)
	return deque
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if deque.size >= this.cap {
		return false
	}
	this.q[this.front] = value
	this.front++
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {

}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {

}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {

}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {

}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {

}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {

}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {

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

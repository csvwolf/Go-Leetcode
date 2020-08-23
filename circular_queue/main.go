package main

type MyCircularQueue struct {
	queue  []int
	head   int
	tail   int
	length int
}

/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		queue:  make([]int, k),
		head:   -1,
		tail:   -1,
		length: 0,
	}
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	capacity := cap(this.queue)
	if this.length == capacity {
		return false
	}
	if this.head == -1 {
		this.head = 0
	}
	this.tail++
	if this.tail == cap(this.queue) {
		this.tail = 0
	}
	this.queue[this.tail] = value

	this.length++
	return true
}

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if this.length == 0 {
		return false
	}
	this.head++
	if this.head == cap(this.queue) {
		this.head = 0
	}
	this.length--
	return true
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if this.length == 0 {
		return -1
	}
	return this.queue[this.head]
}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if this.length == 0 {
		return -1
	}
	return this.queue[this.tail]
}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	if this.length == 0 {
		return true
	}
	return false
}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	if this.length == cap(this.queue) {
		return true
	}
	return false
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */

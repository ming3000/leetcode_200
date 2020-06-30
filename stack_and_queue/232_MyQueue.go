package stack_and_queue

type MyQueue struct {
	input  []int
	output []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{input: make([]int, 0), output: make([]int, 0)}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.input = append(this.input, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	this.Shift()
	top := this.output[len(this.output)-1]
	this.output = this.output[:len(this.output)-1]
	return top
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	this.Shift()
	return this.output[len(this.output)-1]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return (len(this.input) + len(this.output)) == 0
}

func (this *MyQueue) Shift() {
	if len(this.output) != 0 {
		return
	} // if>
	for len(this.input) != 0 {
		tp := this.input[len(this.input)-1]
		this.input = this.input[:len(this.input)-1]
		this.output = append(this.output, tp)
	} // for>
}

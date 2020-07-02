package stack_and_queue

import (
	"leetcode_200/common"
	"math"
)

type MinStack struct {
	stack    []int
	minStack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	// notice the first element of min-stack
	return MinStack{stack: []int{}, minStack: []int{math.MaxInt64}}
}

func (this *MinStack) Push(x int) {
	this.stack = append(this.stack, x)
	theMinTop := this.minStack[len(this.minStack)-1]
	this.minStack = append(this.minStack, common.Min(theMinTop, x))
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}
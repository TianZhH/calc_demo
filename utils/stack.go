package utils

/**
 * @Author: tianzhenhua.nathan@bytedance.com
 * @Date: 2022-09-17
 * @Desc:
 */

type (
	Stack struct {
		first  *node
		top    *node
		length int
	}
	node struct {
		value interface{}
		prev  *node
	}
)

// NewStack Create a new stack
func NewStack() *Stack {
	return &Stack{nil, nil, 0}
}

// Len Return the number of items in the stack
func (this *Stack) Len() int {
	return this.length
}

// Peek View the top item on the stack
func (this *Stack) Peek() interface{} {
	if this.length == 0 {
		return nil
	}
	return this.top.value
}

func (this *Stack) First() interface{} {
	if this.length == 0 {
		return nil
	}
	return this.first.value
}

// Pop the top item of the stack and return it
func (this *Stack) Pop() interface{} {
	if this.length == 0 {
		return nil
	}

	n := this.top
	this.top = n.prev
	this.length--
	if this.length == 0 {
		this.first = nil
	}
	return n.value
}

// Push a value onto the top of the stack
func (this *Stack) Push(value interface{}) {
	n := &node{value, this.top}
	this.top = n
	this.length++
	if this.length == 1 {
		this.first = n
	}
}

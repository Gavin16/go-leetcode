package _55_MinStack

import "testing"

func TestMinStack(t *testing.T) {

	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-1)
	println(minStack.GetMin())
	println(minStack.Top())
	minStack.Pop()
	println(minStack.GetMin())

}

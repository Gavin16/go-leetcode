package _55_MinStack

// 155. 最小栈
// 设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。
//
// 实现 MinStack 类:
//
// MinStack() 初始化堆栈对象。
// void push(int val) 将元素val推入堆栈。
// void pop() 删除堆栈顶部的元素。
// int top() 获取堆栈顶部的元素。
// int getMin() 获取堆栈中的最小元素。
//
// 示例 1:
//
// 输入：
// ["MinStack","push","push","push","getMin","pop","top","getMin"]
// [[],[-2],[0],[-3],[],[],[],[]]
//
// 输出：
// [null,null,null,null,-3,null,0,-2]
//
// 解释：
// MinStack minStack = new MinStack();
// minStack.push(-2);
// minStack.push(0);
// minStack.push(-3);
// minStack.getMin();   --> 返回 -3.
// minStack.pop();
// minStack.top();      --> 返回 0.
// minStack.getMin();   --> 返回 -2.
//
// 提示：
//
// -231 <= val <= 231 - 1
// pop、top 和 getMin 操作总是在 非空栈 上调用
// push, pop, top, and getMin最多被调用 3 * 104 次

// MinStack 击败: 72.48%
type MinStack struct {
	items []Node
}

type Node struct {
	value  int
	minVal int
}

func Constructor() MinStack {
	return MinStack{}
}

func (minStack *MinStack) Push(val int) {
	if len(minStack.items) == 0 {
		minStack.items = append(minStack.items, Node{val, val})
	} else {
		minVal := min(minStack.items[len(minStack.items)-1].minVal, val)
		minStack.items = append(minStack.items, Node{val, minVal})
	}
}

func (minStack *MinStack) Pop() {
	currLen := len(minStack.items)
	minStack.items = minStack.items[:currLen-1]
}

func (minStack *MinStack) Top() int {
	return minStack.items[len(minStack.items)-1].value
}

func (minStack *MinStack) GetMin() int {
	return minStack.items[len(minStack.items)-1].minVal
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

package _336_SmallestInfiniteSet

import (
	"container/heap"
	"sort"
)

// 2336. 无限集中的最小数字
// 现有一个包含所有正整数的集合 [1, 2, 3, 4, 5, ...] 。
//
// 实现 SmallestInfiniteSet 类：
//
// SmallestInfiniteSet() 初始化 SmallestInfiniteSet 对象以包含 所有 正整数。
// int popSmallest() 移除 并返回该无限集中的最小整数。
// void addBack(int num) 如果正整数 num 不 存在于无限集中，则将一个 num 添加 到该无限集中。
//
// 示例：
//
// 输入
// ["SmallestInfiniteSet", "addBack", "popSmallest", "popSmallest", "popSmallest", "addBack", "popSmallest", "popSmallest", "popSmallest"]
// [[], [2], [], [], [], [1], [], [], []]
// 输出
// [null, null, 1, 2, 3, null, 1, 4, 5]
//
// 解释
// SmallestInfiniteSet smallestInfiniteSet = new SmallestInfiniteSet();
// smallestInfiniteSet.addBack(2);    // 2 已经在集合中，所以不做任何变更。
// smallestInfiniteSet.popSmallest(); // 返回 1 ，因为 1 是最小的整数，并将其从集合中移除。
// smallestInfiniteSet.popSmallest(); // 返回 2 ，并将其从集合中移除。
// smallestInfiniteSet.popSmallest(); // 返回 3 ，并将其从集合中移除。
// smallestInfiniteSet.addBack(1);    // 将 1 添加到该集合中。
// smallestInfiniteSet.popSmallest(); // 返回 1 ，因为 1 在上一步中被添加到集合中，
//
//	// 且 1 是最小的整数，并将其从集合中移除。
//
// smallestInfiniteSet.popSmallest(); // 返回 4 ，并将其从集合中移除。
// smallestInfiniteSet.popSmallest(); // 返回 5 ，并将其从集合中移除。
//
// 提示：
//
// 1 <= num <= 1000
// 最多调用 popSmallest 和 addBack 方法 共计 1000 次

type SmallestInfiniteSet struct {
	dict    []int
	minHeap *MinHeap
}

func Constructor() SmallestInfiniteSet {
	existDict := make([]int, 1000)
	minHeap := &MinHeap{}
	heap.Init(minHeap)
	for i := range existDict {
		existDict[i] = 1
		heap.Push(minHeap, i+1)
	}
	set := SmallestInfiniteSet{
		dict:    existDict,
		minHeap: minHeap,
	}
	return set
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	ele := heap.Pop(this.minHeap).(int)
	this.dict[ele-1] = 0
	return ele
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	if this.dict[num-1] == 0 {
		this.dict[num-1] = 1
		heap.Push(this.minHeap, num)
	}
}

type MinHeap struct {
	sort.IntSlice
}

func (h *MinHeap) Push(v any) {
	h.IntSlice = append(h.IntSlice, v.(int))
}

func (h *MinHeap) Pop() any {
	old := h.IntSlice
	pop := old[len(old)-1]
	old = old[:len(old)-1]
	h.IntSlice = old
	return pop
}

/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.PopSmallest();
 * obj.AddBack(num);
 */

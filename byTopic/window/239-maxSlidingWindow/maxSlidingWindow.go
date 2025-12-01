package _39_maxSlidingWindow

import "container/heap"

// 239. 滑动窗口最大值
// 给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。
// 你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
//
// 返回 滑动窗口中的最大值 。
//
// 示例 1：
//
// 输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
// 输出：[3,3,5,5,6,7]
// 解释：
// 滑动窗口的位置                最大值
// ---------------               -----
// [1  3  -1] -3  5  3  6  7       3
// 1 [3  -1  -3] 5  3  6  7       3
// 1  3 [-1  -3  5] 3  6  7       5
// 1  3  -1 [-3  5  3] 6  7       5
// 1  3  -1  -3 [5  3  6] 7       6
// 1  3  -1  -3  5 [3  6  7]      7
// 示例 2：
//
// 输入：nums = [1], k = 1
// 输出：[1]
//
// 提示：
//
// 1 <= nums.length <= 105
// -104 <= nums[i] <= 104
// 1 <= k <= nums.length

func maxSlidingWindow(nums []int, k int) []int {
	n, maxHeap := len(nums), MaxHeap{}

	ans := make([]int, n-k+1)
	heap.Init(&maxHeap)
	for i := 0; i < k; i++ {
		heap.Push(&maxHeap, &Node{pos: i, val: nums[i]})
	}
	ans[0] = maxHeap.Peek().val

	for i := k; i < n; i++ {
		for maxHeap.Len() > 0 && maxHeap.Peek().pos < i-k+1 {
			heap.Pop(&maxHeap)
		}
		heap.Push(&maxHeap, &Node{pos: i, val: nums[i]})
		ans[i-k+1] = maxHeap.Peek().val
	}
	return ans
}

type Node struct {
	pos int
	val int
}

type MaxHeap []*Node

func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Less(i, j int) bool {
	if h[i].val == h[j].val {
		return h[i].pos < h[j].pos
	}
	return h[i].val > h[j].val
}
func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(*Node))
}
func (h *MaxHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}
func (h MaxHeap) Peek() *Node {
	return h[0]
}

package _15_findKthLargest

import (
	"container/heap"
	"sort"
)

// 215. 数组中的第K个最大元素
// 给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
//
// 请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
//
// 你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。
//
// 示例 1:
//
// 输入: [3,2,1,5,6,4], k = 2
// 输出: 5
// 示例 2:
//
// 输入: [3,2,3,1,2,4,5,5,6], k = 4
// 输出: 4
//
// 提示：
//
// 1 <= k <= nums.length <= 105
// -104 <= nums[i] <= 104

// 经典实现方式
// MinHeap 定义使用 sort.IntSlice, 这样通过结构体嵌入实现了 Len, Less, Swap 方法的类型, 可以减少代码量
// 使用小顶堆每次遇到大于小顶堆的元素,就进行替换, 可以确保堆内一直存储着前K大的元素
func findKthLargest(nums []int, k int) int {
	h := &MinHeap{make([]int, 0)}
	heap.Init(h)

	for i := 0; i < len(nums); i++ {
		if i < k {
			heap.Push(h, nums[i])
		} else {
			if nums[i] > h.Peek() {
				h.Replace(0, nums[i])
			}
		}
	}
	return h.Peek()
}

type MinHeap struct {
	sort.IntSlice
}

func (h *MinHeap) Push(v any)   { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *MinHeap) Peek() int    { e := h.IntSlice[0]; return e }
func (h *MinHeap) Pop() (_ any) { return }
func (h *MinHeap) Replace(i, v int) int {
	pop := h.IntSlice[i]
	h.IntSlice[i] = v
	heap.Fix(h, i)
	return pop
}

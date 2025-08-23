package _73_kSmallestPairs

// 给定两个以 非递减顺序排列 的整数数组 nums1 和 nums2 , 以及一个整数 k 。
//
// 定义一对值 (u,v)，其中第一个元素来自 nums1，第二个元素来自 nums2 。
//
// 请找到和最小的 k 个数对 (u1,v1),  (u2,v2)  ...  (uk,vk) 。
//
// 示例 1:
//
// 输入: nums1 = [1,7,11], nums2 = [2,4,6], k = 3
// 输出: [1,2],[1,4],[1,6]
// 解释: 返回序列中的前 3 对数：
//
//	[1,2],[1,4],[1,6],[7,2],[7,4],[11,2],[7,6],[11,4],[11,6]
//
// 示例 2:
//
// 输入: nums1 = [1,1,2], nums2 = [1,2,3], k = 2
// 输出: [1,1],[1,1]
// 解释: 返回序列中的前 2 对数：
//
//	[1,1],[1,1],[1,2],[2,1],[1,2],[2,2],[1,3],[1,3],[2,3]
//
// 提示:
//
// 1 <= nums1.length, nums2.length <= 105
// -109 <= nums1[i], nums2[i] <= 109
// nums1 和 nums2 均为 升序排列
// 1 <= k <= 104
// k <= nums1.length * nums2.length

//func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
//	res := make([][]int, 0)
//
//	for i := 0; i < len(nums1); i++ {
//		for j := 0; j < len(nums2); j++ {
//			res = append(res, []int{nums1[i], nums2[j]})
//		}
//	}
//
//	sort.Slice(res, func(i, j int) bool {
//		return sum(res[i]) < sum(res[j])
//	})
//
//	return res[:k]
//}
//
//func sum(nums []int) int {
//	total := 0
//	for _, num := range nums {
//		total += num
//	}
//	return total
//}

import (
	"container/heap"
)

type Pair struct {
	sum  int // nums1[i] + nums2[j]
	i, j int // index in nums1 and nums2
}
type MinHeap []Pair

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i].sum < h[j].sum }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(Pair)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	p := old[n-1]
	*h = old[:n-1]
	return p
}

func kSmallestPairs(nums1, nums2 []int, k int) [][]int {
	res := [][]int{}
	if len(nums1) == 0 || len(nums2) == 0 || k == 0 {
		return res
	}

	h := &MinHeap{}
	heap.Init(h)

	// 初始化堆：只需要固定 nums1[i=0], 遍历前 min(k, len(nums2)) 个 nums2[j]
	for j := 0; j < len(nums2) && j < k; j++ {
		heap.Push(h, Pair{nums1[0] + nums2[j], 0, j})
	}

	for len(res) < k && h.Len() > 0 {
		p := heap.Pop(h).(Pair)
		i, j := p.i, p.j
		res = append(res, []int{nums1[i], nums2[j]})

		// 如果还能往下扩展 nums1 的下一个元素
		if i+1 < len(nums1) {
			heap.Push(h, Pair{nums1[i+1] + nums2[j], i + 1, j})
		}
	}
	return res
}

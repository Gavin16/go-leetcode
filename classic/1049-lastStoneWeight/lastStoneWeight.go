// package _049_lastStoneWeight
package _049_lastStoneWeight

import (
	"container/heap"
	"math"
)

// 1046. 最后一块石头的重量
// 有一堆石头，每块石头的重量都是正整数。
//
// 每一回合，从中选出两块 最重的 石头，然后将它们一起粉碎。假设石头的重量分别为 x 和 y，且 x <= y。那么粉碎的可能结果如下：
//
// 如果 x == y，那么两块石头都会被完全粉碎；
// 如果 x != y，那么重量为 x 的石头将会完全粉碎，而重量为 y 的石头新重量为 y-x。
// 最后，最多只会剩下一块石头。返回此石头的重量。如果没有石头剩下，就返回 0。
//
// 示例：
//
// 输入：[2,7,4,1,8,1]
// 输出：1
// 解释：
// 先选出 7 和 8，得到 1，所以数组转换为 [2,4,1,1,1]，
// 再选出 2 和 4，得到 2，所以数组转换为 [2,1,1,1]，
// 接着是 2 和 1，得到 1，所以数组转换为 [1,1,1]，
// 最后选出 1 和 1，得到 0，最终数组转换为 [1]，这就是最后剩下那块石头的重量。
//
//
// 提示：
//
// 1 <= stones.length <= 30
// 1 <= stones[i] <= 1000

func main() {
	stones := []int{2, 7, 4, 1, 8, 1}
	println(lastStoneWeight(stones))
}

func lastStoneWeight(stones []int) int {
	maxHeap := MaxHeap(stones)
	heap.Init(&maxHeap)
	for maxHeap.Len() > 1 {
		// 取出最大的两个石头
		stone1 := heap.Pop(&maxHeap).(int)
		stone2 := heap.Pop(&maxHeap).(int)
		if stone1 != stone2 {
			heap.Push(&maxHeap, int(math.Abs(float64(stone1-stone2))))
		}
	}
	if maxHeap.Len() == 0 {
		return 0
	} else {
		return heap.Pop(&maxHeap).(int)
	}
}

// 定义大顶堆结构体
type MaxHeap []int

// Len 方法返回堆的元素数量
func (h MaxHeap) Len() int {
	return len(h)
}

// Less 方法定义大顶堆的排序规则，i 和 j 对应堆中的元素
// 如果 h[i] > h[j]，返回 true 以保持大顶堆性质
func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

// Swap 方法交换堆中两个元素的位置
func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push 方法将一个元素添加到堆中
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop 方法从堆中移除并返回堆顶元素
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

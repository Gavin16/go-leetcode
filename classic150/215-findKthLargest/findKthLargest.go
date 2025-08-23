package _15_findKthLargest

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

// TODO 击败: 51.49%
func findKthLargest(nums []int, k int) int {
	heap := MinHeap{elements: make([]int, 0, k)}
	for i := 0; i < len(nums); i++ {
		if len(heap.elements) < k {
			heap.insert(nums[i])
		} else {
			if nums[i] > heap.elements[0] {
				heap.extractMin()
				heap.insert(nums[i])
			}
		}
	}
	return heap.elements[0]
}

type MinHeap struct {
	elements []int
}

// 插入元素并堆化
func (h *MinHeap) insert(num int) {
	h.elements = append(h.elements, num)
	h.heapUp(len(h.elements) - 1)
}

func (h *MinHeap) parent(index int) int {
	return (index - 1) / 2
}
func (h *MinHeap) leftChild(index int) int {
	return 2*index + 1
}
func (h *MinHeap) rightChild(index int) int {
	return 2*index + 2
}

// 末尾元素堆化--向上堆化
func (h *MinHeap) heapUp(index int) {
	for index > 0 && h.elements[h.parent(index)] > h.elements[index] {
		h.elements[h.parent(index)], h.elements[index] = h.elements[index], h.elements[h.parent(index)]
		index = h.parent(index)
	}
}

// 提取最小元素
func (h *MinHeap) extractMin() int {
	if len(h.elements) == 0 {
		return -1
	}
	minimum := h.elements[0]
	h.elements[0] = h.elements[len(h.elements)-1]
	h.elements = h.elements[:len(h.elements)-1]
	h.heapDown(0)
	return minimum
}

// 堆顶元素向下堆化
func (h *MinHeap) heapDown(index int) {
	smallest := index
	left := h.leftChild(index)
	right := h.rightChild(index)

	if left < len(h.elements) && h.elements[left] < h.elements[smallest] {
		smallest = left
	}
	if right < len(h.elements) && h.elements[right] < h.elements[smallest] {
		smallest = right
	}
	if smallest != index {
		h.elements[index], h.elements[smallest] = h.elements[smallest], h.elements[index]
		h.heapDown(smallest)
	}
}

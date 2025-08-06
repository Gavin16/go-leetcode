package _95_MedianFinder

import "container/heap"

// 295. 数据流的中位数
// 中位数是有序整数列表中的中间值。如果列表的大小是偶数，则没有中间值，中位数是两个中间值的平均值。
//
// 例如 arr = [2,3,4] 的中位数是 3 。
// 例如 arr = [2,3] 的中位数是 (2 + 3) / 2 = 2.5 。
// 实现 MedianFinder 类:
//
// MedianFinder() 初始化 MedianFinder 对象。
//
// void addNum(int num) 将数据流中的整数 num 添加到数据结构中。
//
// double findMedian() 返回到目前为止所有元素的中位数。与实际答案相差 10-5 以内的答案将被接受。
//
// 示例 1：
//
// 输入
// ["MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"]
// [[], [1], [2], [], [3], []]
// 输出
// [null, null, null, 1.5, null, 2.0]
//
// 解释
// MedianFinder medianFinder = new MedianFinder();
// medianFinder.addNum(1);    // arr = [1]
// medianFinder.addNum(2);    // arr = [1, 2]
// medianFinder.findMedian(); // 返回 1.5 ((1 + 2) / 2)
// medianFinder.addNum(3);    // arr[1, 2, 3]
// medianFinder.findMedian(); // return 2.0
// 提示:
//
// -105 <= num <= 105
// 在调用 findMedian 之前，数据结构中至少有一个元素
// 最多 5 * 104 次调用 addNum 和 findMedian

// 使用一个大顶堆 和  一个小顶堆来存储数组内容
// 其中大顶堆存有序切片中前半部分, 小顶堆保存切片中后半部分
// 当添加元素总长度为奇数时, 返回大顶堆的堆顶
// 当添加元素总长度为偶数时, 返回大顶堆和小顶堆两个堆顶和的一半

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() any          { old := *h; x := old[len(old)-1]; *h = old[:len(old)-1]; return x }
func (h MaxHeap) Peek() int          { return h[0] }

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() any          { old := *h; x := old[len(old)-1]; *h = old[:len(old)-1]; return x }
func (h MinHeap) Peek() int          { return h[0] }

type MedianFinder struct {
	maxHeap MaxHeap
	minHeap MinHeap
}

func Constructor() MedianFinder {
	maxH := MaxHeap{}
	//heap.Init(maxH)
	minH := MinHeap{}
	//heap.Init(minH)

	return MedianFinder{
		maxHeap: maxH,
		minHeap: minH,
	}
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(&this.maxHeap, num)
	heap.Push(&this.minHeap, heap.Pop(&this.maxHeap))

	//this.maxHeap.Push(num)
	//this.minHeap.Push(this.maxHeap.Pop())
	if this.minHeap.Len() > this.maxHeap.Len() {
		//this.maxHeap.Push(this.minHeap.Pop())
		heap.Push(&this.maxHeap, heap.Pop(&this.minHeap))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	l1, l2 := this.maxHeap.Len(), this.minHeap.Len()
	if l1 > l2 {
		return float64(this.maxHeap.Peek())
	} else {
		maxPeek := this.maxHeap.Peek()
		minPeek := this.minHeap.Peek()
		return float64(maxPeek+minPeek) / 2.0
	}
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

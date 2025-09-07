package _02_findMaximizedCapital

import (
	"container/heap"
	"sort"
)

// 502. IPO
// 假设 力扣（LeetCode）即将开始 IPO 。为了以更高的价格将股票卖给风险投资公司，力扣 希望在 IPO 之前开展一些项目以增加其资本。 由于资源有限，它只能在 IPO 之前完成最多 k 个不同的项目。帮助 力扣 设计完成最多 k 个不同项目后得到最大总资本的方式。
//
// 给你 n 个项目。对于每个项目 i ，它都有一个纯利润 profits[i] ，和启动该项目需要的最小资本 capital[i] 。
//
// 最初，你的资本为 w 。当你完成一个项目时，你将获得纯利润，且利润将被添加到你的总资本中。
//
// 总而言之，从给定项目中选择 最多 k 个不同项目的列表，以 最大化最终资本 ，并输出最终可获得的最多资本。
//
// 答案保证在 32 位有符号整数范围内。
//
// 示例 1：
//
// 输入：k = 2, w = 0, profits = [1,2,3], capital = [0,1,1]
// 输出：4
// 解释：
// 由于你的初始资本为 0，你仅可以从 0 号项目开始。
// 在完成后，你将获得 1 的利润，你的总资本将变为 1。
// 此时你可以选择开始 1 号或 2 号项目。
// 由于你最多可以选择两个项目，所以你需要完成 2 号项目以获得最大的资本。
// 因此，输出最后最大化的资本，为 0 + 1 + 3 = 4。
// 示例 2：
//
// 输入：k = 3, w = 0, profits = [1,2,3], capital = [0,1,2]
// 输出：6
//
// 提示：
//
// 1 <= k <= 105
// 0 <= w <= 109
// n == profits.length
// n == capital.length
// 1 <= n <= 105
// 0 <= profits[i] <= 104
// 0 <= capital[i] <= 109

// 使用大顶堆保存利润 profits[i] - capital[i]
// 对cap 做成本排序，仅取出w >= cap[i] 对应的项目加入大顶堆
// 结束一个项目后, k -= 1, w += profits[j] - capital[j]

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	n, projects := len(profits), make([]*Project, 0)
	for i := 0; i < n; i++ {
		projects = append(projects,
			&Project{id: i, pro: profits[i], cap: capital[i]})
	}
	sort.Slice(projects, func(i, j int) bool {
		return projects[i].cap < projects[j].cap
	})

	// 使用大顶堆计算
	maxHeap := MaxHeap{}
	heap.Init(&maxHeap)
	pos := 0
	for i := k; i > 0; i-- {
		for ; pos < n && projects[pos].cap <= w; pos++ {
			heap.Push(&maxHeap, projects[pos])
		}
		if len(maxHeap) == 0 {
			continue
		}
		maxEle := heap.Pop(&maxHeap)
		maxProject := maxEle.(*Project)
		w += maxProject.pro
	}
	return w
}

type Project struct {
	id  int
	cap int
	pro int
}

// 元素以[i,cap[i]] 形式保存各项目id 和 对应启动成本

type MaxHeap []*Project

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].pro > h[j].pro }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x any)        { *h = append(*h, x.(*Project)) }
func (h *MaxHeap) Pop() any          { old := *h; x := old[len(old)-1]; *h = old[:len(old)-1]; return x }
func (h MaxHeap) Peek() *Project     { return h[0] }

package _07_canFinish

import (
	"errors"
)

// 207. 课程表
// 你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。
//
// 在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites 给出，其中 prerequisites[i] = [ai, bi] ，
// 表示如果要学习课程 ai 则 必须 先学习课程  bi 。
//
// 例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
// 请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。
//
// 示例 1：
//
// 输入：numCourses = 2, prerequisites = [[1,0]]
// 输出：true
// 解释：总共有 2 门课程。学习课程 1 之前，你需要完成课程 0 。这是可能的。
// 示例 2：
//
// 输入：numCourses = 2, prerequisites = [[1,0],[0,1]]
// 输出：false
// 解释：总共有 2 门课程。学习课程 1 之前，你需要先完成课程 0 ；并且学习课程 0 之前，你还应先完成课程 1 。这是不可能的。
//
// 提示：
//
// 1 <= numCourses <= 2000
// 0 <= prerequisites.length <= 5000
// prerequisites[i].length == 2
// 0 <= ai, bi < numCourses
// prerequisites [i] 中的所有课程对 互不相同

// TODO: 图的拓补排序
// 根据题意, 课程的数量就是  numCourses, 因此要想能修完所有课程, 只要判断图由 prerequisites 组成的图中是否包含有环
// 若图中有环，那么环内的课程必然无法修， 最后也一定不能修完所有课程
// 击败:52.79%
func canFinish(numCourses int, prerequisites [][]int) bool {
	var (
		visited = make([]int, numCourses)
		edges   = make([][]int, numCourses)
		valid   = true
		dfs     func(n int)
	)
	dfs = func(n int) {
		visited[n] = 1
		for _, v := range edges[n] {
			if visited[v] == 0 {
				dfs(v)
				if !valid {
					return
				}
			} else if visited[v] == 1 {
				valid = false
				return
			}
		}
		visited[n] = 2
	}
	for _, v := range prerequisites {
		edges[v[1]] = append(edges[v[1]], v[0])
	}
	for i := 0; i < numCourses && valid; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}
	return valid
}

// 找出拓补排序的结果, 若图中存在环 则返回错误
func canFinish2(numCourses int, prerequisites [][]int) (order []int, err error) {
	var (
		visited = make([]int, numCourses)
		edges   = make([][]int, numCourses)
		valid   = true
		dfs     func(n int)
		result  = make([]int, 0)
	)
	dfs = func(n int) {
		visited[n] = 1
		for _, v := range edges[n] {
			if visited[v] == 0 {
				dfs(v)
				if !valid {
					return
				}
			} else if visited[v] == 1 {
				valid = false
				return
			}
		}
		visited[n] = 2
		result = append(result, n)
	}
	for _, v := range prerequisites {
		edges[v[1]] = append(edges[v[1]], v[0])
	}
	for i := 0; i < numCourses && valid; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}
	if !valid {
		return result, errors.New("can not finish")
	}

	var reverse func([]int)
	reverse = func(a []int) {
		for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
			a[i], a[j] = a[j], a[i]
		}
	}
	reverse(result)

	return result, nil
}

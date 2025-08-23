package _99_calcEquation

// 399. 除法求值
// 给你一个变量对数组 equations 和一个实数值数组 values 作为已知条件，
// 其中 equations[i] = [Ai, Bi] 和 values[i] 共同表示等式 Ai / Bi = values[i]。
// 每个 Ai 或 Bi 是一个表示单个变量的字符串。
//
// 另有一些以数组 queries 表示的问题，其中 queries[j] = [Cj, Dj] 表示第 j 个问题，请你根据已知条件找出 Cj / Dj = ? 的结果作为答案。
//
// 返回 所有问题的答案 。如果存在某个无法确定的答案，则用 -1.0 替代这个答案。
// 如果问题中出现了给定的已知条件中没有出现的字符串，也需要用 -1.0 替代这个答案。
//
// 注意：输入总是有效的。你可以假设除法运算中不会出现除数为 0 的情况，且不存在任何矛盾的结果。
// 注意：未在等式列表中出现的变量是未定义的，因此无法确定它们的答案。
//
// 示例 1：
// 输入：equations = [["a","b"],["b","c"]],
//      values = [2.0,3.0],
//      queries = [["a","c"],["b","a"],["a","e"],["a","a"],["x","x"]]
// 输出：[6.00000,0.50000,-1.00000,1.00000,-1.00000]
//
// 解释：
// 条件：a / b = 2.0, b / c = 3.0
// 问题：a / c = ?, b / a = ?, a / e = ?, a / a = ?, x / x = ?
// 结果：[6.0, 0.5, -1.0, 1.0, -1.0 ]
// 注意：x 是未定义的 => -1.0
//
// 示例 2：
// 输入：equations = [["a","b"],["b","c"],["bc","cd"]],
//      values = [1.5,2.5,5.0],
//      queries = [["a","c"],["c","b"],["bc","cd"],["cd","bc"]]
// 输出：[3.75000,0.40000,5.00000,0.20000]
//
// 示例 3：
// 输入：equations = [["a","b"]],
//      values = [0.5],
//      queries = [["a","b"],["b","a"],["a","c"],["x","y"]]
// 输出：[0.50000,2.00000,-1.00000,-1.00000]
//
// 提示：
//
// 1 <= equations.length <= 20
// equations[i].length == 2
// 1 <= Ai.length, Bi.length <= 5
// values.length == equations.length
// 0.0 < values[i] <= 20.0
// 1 <= queries.length <= 20
// queries[i].length == 2
// 1 <= Cj.length, Dj.length <= 5
// Ai, Bi, Cj, Dj 由小写英文字母与数字组成
//

// 使用图的邻接表表示
// 对于出现过且相互联通的两个变量
// 通过对图进行DFS 搜索可以得到计算结果

func calcEquation(equations [][]string, values []float64,
	queries [][]string) []float64 {

	type Node struct {
		name byte
		adj  map[byte]float64
	}
	adjList := make(map[byte]*Node)

	for i := 0; i < len(equations); i++ {
		eq := equations[i]
		a, b := removeDuplication(eq[0], eq[1])
		if node, ok := adjList[a]; ok {
			node.adj[b] = values[i]
		} else {
			adjList[a] = &Node{name: a, adj: map[byte]float64{b: values[i]}}
			adjList[b] = &Node{name: b, adj: map[byte]float64{a: 1 / values[i]}}
		}
	}
	var dfs func(c, d byte, prod float64) float64
	visited := make([]bool, 255)

	dfs = func(c, d byte, prod float64) float64 {
		visited[c] = true
		if node, ok := adjList[c]; !ok {
			return -1.0
		} else {
			adjMap := node.adj
			for k, v := range adjMap {
				if visited[k] {
					continue
				}
				if k == d {
					return prod * v
				} else {
					if ret := dfs(k, d, prod*v); ret > 0 {
						return ret
					}
				}
			}
			return -1.0
		}
	}
	ans := make([]float64, 0)
	for _, query := range queries {
		q0, q1 := query[0], query[1]
		c, d := removeDuplication(q0, q1)
		val := dfs(c, d, 1.0)
		ans = append(ans, val)
	}
	return ans
}

func removeDuplication(str1, str2 string) (byte, byte) {
	m1, m2 := map[byte]bool{}, map[byte]bool{}
	for i := 0; i < len(str1); i++ {
		m1[str1[i]], m2[str2[i]] = true, true
	}
	var s1, s2 byte
	for k, _ := range m1 {
		if !m2[k] {
			s1 = k
			break
		}
	}
	for k, _ := range m2 {
		if !m1[k] {
			s2 = k
			break
		}
	}
	return s1, s2
}

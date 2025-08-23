package _33_minMutation

// 433. 最小基因变化
// 基因序列可以表示为一条由 8 个字符组成的字符串，其中每个字符都是 'A'、'C'、'G' 和 'T' 之一。
// 假设我们需要调查从基因序列 start 变为 end 所发生的基因变化。一次基因变化就意味着这个基因序列中的一个字符发生了变化。
//
// 例如，"AACCGGTT" --> "AACCGGTA" 就是一次基因变化。
// 另有一个基因库 bank 记录了所有有效的基因变化，只有基因库中的基因才是有效的基因序列。（变化后的基因必须位于基因库 bank 中）
//
// 给你两个基因序列 start 和 end ，以及一个基因库 bank ，请你找出并返回能够使 start 变化为 end 所需的最少变化次数。
// 如果无法完成此基因变化，返回 -1 。
// 注意：起始基因序列 start 默认是有效的，但是它并不一定会出现在基因库中。
//
// 示例 1：
//
// 输入：start = "AACCGGTT", end = "AACCGGTA", bank = ["AACCGGTA"]
// 输出：1
// 示例 2：
//
// 输入：start = "AACCGGTT", end = "AAACGGTA", bank = ["AACCGGTA","AACCGCTA","AAACGGTA"]
// 输出：2
// 示例 3：
//
// 输入：start = "AAAAACCC", end = "AACCCCCC", bank = ["AAAACCCC","AAACCCCC","AACCCCCC"]
// 输出：3
//
// 提示：
//
// start.length == 8
// end.length == 8
// 0 <= bank.length <= 10
// bank[i].length == 8
// start、end 和 bank[i] 仅由字符 ['A', 'C', 'G', 'T'] 组成

// 击败: 23.02%
func minMutation(startGene string, endGene string, bank []string) int {
	type Node struct {
		code      string
		neighbour []*Node
	}
	m := map[string]*Node{}
	for _, v := range bank {
		m[v] = &Node{code: v}
	}
	if _, ok := m[endGene]; !ok {
		return -1
	} else if dist := distance(startGene, endGene); dist <= 1 {
		return dist
	}
	for k, node := range m {
		for i := 0; i < len(bank); i++ {
			code := bank[i]
			if distance(k, code) == 1 {
				nb := m[code]
				node.neighbour = append(node.neighbour, nb)
			}
		}
	}
	// 图搜索: 广度优先搜索
	minStep := 9
	for k, v := range m {
		if distance(k, startGene) == 1 {
			queue, step := []*Node{v}, 1
			visited := map[string]int{}
		stepSearch:
			for len(queue) > 0 {
				size := len(queue)
				// 清空队列代表一次变异
				for i := 0; i < size; i++ {
					q := queue[i]
					visited[q.code] = 1
					for _, nb := range q.neighbour {
						if nb.code == endGene {
							minStep = min(minStep, step+1)
							break stepSearch
						} else if visited[nb.code] == 0 {
							queue = append(queue, nb)
						}
					}
				}
				queue = queue[size:]
				step += 1
			}
		}
	}
	if minStep > 8 {
		return -1
	}
	return minStep
}

// 计算两个基因的距离
func distance(startGene string, bankGene string) int {
	n, dist := len(bankGene), 0
	for i := 0; i < n; i++ {
		if bankGene[i] != startGene[i] {
			dist += 1
		}
	}
	return dist
}

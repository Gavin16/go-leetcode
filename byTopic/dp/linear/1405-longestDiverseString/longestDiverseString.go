package _405_longestDiverseString

import (
	"sort"
)

// 1405. 最长快乐字符串
// 如果字符串中不含有任何 'aaa'，'bbb' 或 'ccc' 这样的字符串作为子串，那么该字符串就是一个「快乐字符串」。
//
// 给你三个整数 a，b ，c，请你返回 任意一个 满足下列全部条件的字符串 s：
//
// s 是一个尽可能长的快乐字符串。
// s 中 最多 有a 个字母 'a'、b 个字母 'b'、c 个字母 'c' 。
// s 中只含有 'a'、'b' 、'c' 三种字母。
// 如果不存在这样的字符串 s ，请返回一个空字符串 ""。
//
// 示例 1：
//
// 输入：a = 1, b = 1, c = 7
// 输出："ccaccbcc"
// 解释："ccbccacc" 也是一种正确答案。
// 示例 2：
//
// 输入：a = 2, b = 2, c = 1
// 输出："aabbc"
// 示例 3：
//
// 输入：a = 7, b = 1, c = 0
// 输出："aabaa"
// 解释：这是该测试用例的唯一正确答案。
//
// 提示：
//
// 0 <= a, b, c <= 100
// a + b + c > 0

func longestDiverseString(a int, b int, c int) string {
	// 构造字符串时每次使用最长的那个
	// 并在使用前判断末尾字符串数量是否已经达到2
	// 若最后发现最长的字符串和末尾的字符串都已经相同，则代表已经达到最大长度
	cnt := []struct {
		ch  byte
		num int
	}{{ch: 'a', num: a},
		{ch: 'b', num: b},
		{ch: 'c', num: c}}
	ans := make([]byte, 0)
	for {
		sort.Slice(cnt, func(i, j int) bool {
			return cnt[i].num > cnt[j].num
		})
		hasNext := false
		for i, p := range cnt {
			if p.num <= 0 {
				break
			}
			l := len(ans)
			if l >= 2 && ans[l-1] == p.ch && ans[l-2] == p.ch {
				continue
			}
			ans = append(ans, p.ch)
			cnt[i].num -= 1
			hasNext = true
			break
		}
		if !hasNext {
			return string(ans)
		}
	}
}

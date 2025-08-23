package _0_findSubstring

import "slices"

// 30. 串联所有单词的子串
// 给定一个字符串 s 和一个字符串数组 words。 words 中所有字符串 长度相同。
//
// s 中的 串联子串 是指一个包含  words 中所有字符串以任意顺序排列连接起来的子串。
//
// 例如，如果 words = ["ab","cd","ef"]， 那么 "abcdef"， "abefcd"，"cdabef"， "cdefab"，"efabcd"， 和 "efcdab" 都是串联子串。
// "acdbef" 不是串联子串，因为他不是任何 words 排列的连接。
// 返回所有串联子串在 s 中的开始索引。你可以以 任意顺序 返回答案。
//
// 示例 1：
//
// 输入：s = "barfoothefoobarman", words = ["foo","bar"]
// 输出：[0,9]
// 解释：因为 words.length == 2 同时 words[i].length == 3，连接的子字符串的长度必须为 6。
// 子串 "barfoo" 开始位置是 0。它是 words 中以 ["bar","foo"] 顺序排列的连接。
// 子串 "foobar" 开始位置是 9。它是 words 中以 ["foo","bar"] 顺序排列的连接。
// 输出顺序无关紧要。返回 [9,0] 也是可以的。
// 示例 2：
//
// 输入：s = "wordgoodgoodgoodbestword", words = ["word","good","best","word"]
// 输出：[]
// 解释：因为 words.length == 4 并且 words[i].length == 4，所以串联子串的长度必须为 16。
// s 中没有子串长度为 16 并且等于 words 的任何顺序排列的连接。
// 所以我们返回一个空数组。
// 示例 3：
//
// 输入：s = "barfoofoobarthefoobarman", words = ["bar","foo","the"]
// 输出：[6,9,12]
// 解释：因为 words.length == 3 并且 words[i].length == 3，所以串联子串的长度必须为 9。
// 子串 "foobarthe" 开始位置是 6。它是 words 中以 ["foo","bar","the"] 顺序排列的连接。
// 子串 "barthefoo" 开始位置是 9。它是 words 中以 ["bar","the","foo"] 顺序排列的连接。
// 子串 "thefoobar" 开始位置是 12。它是 words 中以 ["the","foo","bar"] 顺序排列的连接。
//
// 提示：
//
// 1 <= s.length <= 104
// 1 <= words.length <= 5000
// 1 <= words[i].length <= 30
// words[i] 和 s 由小写英文字母组成

// 时间复杂度O(m*n),空间复杂度O(m) 击败: 9.48%
func findSubstring(s string, words []string) []int {
	slices.Sort(words)
	num, wLen := len(words), len(words[0])
	window := num * wLen
	if len(s) < window {
		return []int(nil)
	}
	ret := make([]int, 0)
	for i := 0; i+window <= len(s); i += 1 {
		slider := s[i : i+window]
		temp := make([]string, 0)
		for j := 0; j < num; j += 1 {
			temp = append(temp, slider[j*wLen:(j+1)*wLen])
		}
		slices.Sort(temp)
		pos := 0
		for ; pos < num; pos++ {
			if words[pos] != temp[pos] {
				break
			}
		}
		if pos == num {
			ret = append(ret, i)
		}
	}
	return ret
}

// 解法2: 使用map记录滑动窗口中出现的长度为 word 长度的字符串的出现
// 初始化窗口时, 就将 words 中的计数写入到 map 中，对应所有单词的计数 - 1
// 这时在字符串s 对应窗口区间内 如果存在words 中的某一个单词, 那么map中对应单词(key) 对应的value 为0
// 接下来开始滑动窗口
//
//	对于后面等长字符串进入窗口范围内时, map对应的key的value + 1
//	对于前面等长字符串离开窗口范围时，map 对应key 的value - 1
//
// 这样做相当于把 words 内所有单词都写入到了 map 中， 移动窗口时，离开窗口的等长字符串的计数又会被减1
// 等价于map只实时记录了 s 在该窗口中的子串(+1) 和 words 中所有的单词的计数(-1)
// 当 s 在该窗口中的子串是words 中所有单词的拼接时， map 中所有的key 对应value 都为0
//
//	这里value 为0，等效于删除该k-v 对，因为golang 中的map 默认map中的value 零值可用，新进人窗口的子串可以直接进行map[key]++操作
//
// 击败: 68.63%
func findSubstring2(s string, words []string) []int {
	m, n, strLen := len(words), len(words[0]), len(s)

	ret := make([]int, 0)
	// 根据单词长度进行错位遍历
	for i := 0; i < n && i+m*n <= strLen; i += 1 {
		// 找出每次错位，各个单词的统计计数
		// 窗口范围内 s 中出现则+1, words 中出现则 -1
		diff := map[string]int{}

		// 第一个窗口初始化
		for j := 0; j < m; j += 1 {
			diff[s[i+j*n:i+(j+1)*n]] += 1
		}
		for _, v := range words {
			diff[v] -= 1
			if diff[v] == 0 {
				delete(diff, v)
			}
		}
		// 错位i 的情况下滑动
		// 进入窗口的对应计数 + 1  离开窗口的对应计数 - 1
		for k := i; k <= strLen-m*n; k += n {
			if k != i {
				word := s[k+(m-1)*n : k+m*n]
				diff[word] += 1
				if diff[word] == 0 {
					delete(diff, word)
				}

				word = s[k-n : k]
				diff[word] -= 1
				if diff[word] == 0 {
					delete(diff, word)
				}
			}

			if len(diff) == 0 {
				ret = append(ret, k)
			}
		}
	}
	return ret
}

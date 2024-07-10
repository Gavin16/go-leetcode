package _9_groupAnagrams

// 49. 字母异位词分组
// 给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
//
// 字母异位词 是由重新排列源单词的所有字母得到的一个新单词。
//
// 示例 1:
//
// 输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
// 输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
// 示例 2:
//
// 输入: strs = [""]
// 输出: [[""]]
// 示例 3:
//
// 输入: strs = ["a"]
// 输出: [["a"]]
//
// 提示：
//
// 1 <= strs.length <= 104
// 0 <= strs[i].length <= 100
// strs[i] 仅包含小写字母

// 击败: 57.72%
func groupAnagrams(strs []string) [][]string {
	statMap := make(map[[26]byte][]string)
	for i := 0; i < len(strs); i++ {
		cc := [26]byte{}
		for j := 0; j < len(strs[i]); j++ {
			cc[strs[i][j]-'a'] += 1
		}

		if v, ok := statMap[cc]; !ok {
			slice := []string{strs[i]}
			statMap[cc] = slice
		} else {
			v = append(v, strs[i])
			statMap[cc] = v
		}
	}
	ans := make([][]string, 0)
	for _, v := range statMap {
		ans = append(ans, v)
	}
	return ans
}

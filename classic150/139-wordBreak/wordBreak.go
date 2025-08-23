package _39_wordBreak

// 给你一个字符串 s 和一个字符串列表 wordDict 作为字典。如果可以利用字典中出现的一个或多个单词拼接出 s 则返回 true。
//
// 注意：不要求字典中出现的单词全部都使用，并且字典中的单词可以重复使用。
//
// 示例 1：
//
// 输入: s = "leetcode", wordDict = ["leet", "code"]
// 输出: true
// 解释: 返回 true 因为 "leetcode" 可以由 "leet" 和 "code" 拼接成。
// 示例 2：
//
// 输入: s = "applepenapple", wordDict = ["apple", "pen"]
// 输出: true
// 解释: 返回 true 因为 "applepenapple" 可以由 "apple" "pen" "apple" 拼接成。
//     注意，你可以重复使用字典中的单词。
// 示例 3：
//
// 输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
// 输出: false
//
//
// 提示：
//
// 1 <= s.length <= 300
// 1 <= wordDict.length <= 1000
// 1 <= wordDict[i].length <= 20
// s 和 wordDict[i] 仅由小写英文字母组成
// wordDict 中的所有字符串 互不相同

// 动态规划
// dp[i] 表示字符串[0,i]子串能否被wordDict中的单词拼接, len(s) = n
// 当 i = 0时, dp[0] = true
// 从前向后遍历 i, 当dp[i] = true 时, 子串[i+1, j] (j = i+1, i+2, ..., n)若能被wordDict中的单词拼接, 则dp[j] = true
// 这样使 i 从0->n-1, j=i+1 -> n 遍历, 就能找出dp[i]中的所有值, 最终的结果即为dp[n]
func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	wordDictSet := map[string]bool{}
	for _, w := range wordDict {
		wordDictSet[w] = true
	}
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			if dp[i] && wordDictSet[s[i:j]] {
				dp[j] = true
			}
		}
	}
	return dp[len(s)]
}

package _7_letterCombinations

//
// 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
// 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
//
// 示例 1：
//
// 输入：digits = "23"
// 输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
// 示例 2：
//
// 输入：digits = ""
// 输出：[]
// 示例 3：
//
// 输入：digits = "2"
// 输出：["a","b","c"]
//
// 提示：
//
// 0 <= digits.length <= 4
// digits[i] 是范围 ['2', '9'] 的一个数字。

func letterCombinations(digits string) []string {
	dict := map[byte]string{
		'2': "abc", '3': "def", '4': "ghi", '5': "jkl",
		'6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz",
	}
	var ans []string
	var temp []byte
	var dfs func(i int)
	dfs = func(i int) {
		if i == len(digits) {
			if len(temp) > 0 {
				ans = append(ans, string(temp))
			}
			return
		}
		ch := digits[i]
		str := dict[ch]
		for j := 0; j < len(str); j++ {
			temp = append(temp, str[j])
			dfs(i + 1)
			temp = temp[:len(temp)-1]
		}
	}
	dfs(0)
	return ans
}

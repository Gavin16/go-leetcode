package _7_letterCombinations

// 17. 电话号码的字母组合
// 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
//
// 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
//
// 示例 1：
//
// 输入：digits = "23"
// 输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]
//
// 示例 2：
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

// 2-> 'a','b','c'; 3-> 'd', 'e', 'f'; 4-> 'g','h','i'; 5-> 'j','k','l'
// 6-> 'm','n','o'; 7-> 'p', 'q', 'r', 's' ; 8 -> 't','u','v'; 9 -> 'w','x','y','z'
// 表示的字母组合的长度  需要与 digits 长度一致

// 击败: 100.00%
func letterCombinations(digits string) []string {
	bytes := make([]byte, 0)
	ans := make([]string, 0)
	d2c := map[int][]byte{
		2: {'a', 'b', 'c'}, 3: {'d', 'e', 'f'}, 4: {'g', 'h', 'i'}, 5: {'j', 'k', 'l'},
		6: {'m', 'n', 'o'}, 7: {'p', 'q', 'r', 's'}, 8: {'t', 'u', 'v'}, 9: {'w', 'x', 'y', 'z'},
	}
	var dfs func(index int)
	dfs = func(index int) {
		if index == len(digits) {
			if len(bytes) > 0 {
				ans = append(ans, string(bytes))
			}
			return
		}
		n := int(digits[index])
		chars := d2c[n-'0']
		for _, char := range chars {
			bytes = append(bytes, char)
			dfs(index + 1)
			bytes = bytes[:len(bytes)-1]
		}
	}
	dfs(0)
	return ans
}

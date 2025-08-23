package _83_canConstruct

// 383. 赎金信
// 给你两个字符串：ransomNote 和 magazine ，判断 ransomNote 能不能由 magazine 里面的字符构成。
//
// 如果可以，返回 true ；否则返回 false 。
//
// magazine 中的每个字符只能在 ransomNote 中使用一次。
//
// 示例 1：
//
// 输入：ransomNote = "a", magazine = "b"
// 输出：false
// 示例 2：
//
// 输入：ransomNote = "aa", magazine = "ab"
// 输出：false
// 示例 3：
//
// 输入：ransomNote = "aa", magazine = "aab"
// 输出：true
//
// 提示：
//
// 1 <= ransomNote.length, magazine.length <= 105
// ransomNote 和 magazine 由小写英文字母组成

// 击败: 29.88%
func canConstruct(ransomNote string, magazine string) bool {
	magMap := map[byte]int{}
	for i := 0; i < len(magazine); i++ {
		magMap[magazine[i]] += 1
	}

	for j := 0; j < len(ransomNote); j++ {
		_, ok := magMap[ransomNote[j]]
		if !ok {
			return false
		}
		magMap[ransomNote[j]] -= 1
		if magMap[ransomNote[j]] < 0 {
			return false
		}
	}
	return true
}

// 使用数组作为map记录
// 击败: 60.58%
func canConstruct1(ransomNote string, magazine string) bool {
	magSlice := [26]int{}
	for i := 0; i < len(magazine); i++ {
		magSlice[magazine[i]-'a'] += 1
	}

	for j := 0; j < len(ransomNote); j++ {
		magSlice[ransomNote[j]-'a'] -= 1
		if magSlice[ransomNote[j]-'a'] < 0 {
			return false
		}
	}
	return true
}

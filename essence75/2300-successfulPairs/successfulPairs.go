package _300_successfulPairs

import (
	"sort"
)

// 2300. 咒语和药水的成功对数
// 给你两个正整数数组 spells 和 potions ，长度分别为 n 和 m ，其中 spells[i] 表示第 i 个咒语的能量强度，potions[j] 表示第 j 瓶药水的能量强度。
//
// 同时给你一个整数 success 。一个咒语和药水的能量强度 相乘 如果 大于等于 success ，那么它们视为一对 成功 的组合。
//
// 请你返回一个长度为 n 的整数数组 pairs，其中 pairs[i] 是能跟第 i 个咒语成功组合的 药水 数目。
//
// 示例 1：
//
// 输入：spells = [5,1,3], potions = [1,2,3,4,5], success = 7
// 输出：[4,0,3]
// 解释：
// - 第 0 个咒语：5 * [1,2,3,4,5] = [5,10,15,20,25] 。总共 4 个成功组合。
// - 第 1 个咒语：1 * [1,2,3,4,5] = [1,2,3,4,5] 。总共 0 个成功组合。
// - 第 2 个咒语：3 * [1,2,3,4,5] = [3,6,9,12,15] 。总共 3 个成功组合。
// 所以返回 [4,0,3] 。
// 示例 2：
//
// 输入：spells = [3,1,2], potions = [8,5,8], success = 16
// 输出：[2,0,2]
// 解释：
// - 第 0 个咒语：3 * [8,5,8] = [24,15,24] 。总共 2 个成功组合。
// - 第 1 个咒语：1 * [8,5,8] = [8,5,8] 。总共 0 个成功组合。
// - 第 2 个咒语：2 * [8,5,8] = [16,10,16] 。总共 2 个成功组合。
// 所以返回 [2,0,2] 。
//
// 提示：
//
// n == spells.length
// m == potions.length
// 1 <= n, m <= 105
// 1 <= spells[i], potions[i] <= 105
// 1 <= success <= 1010

func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Slice(potions, func(i, j int) bool {
		return potions[i] < potions[j]
	})
	ans := make([]int, len(spells))
	for i, num := range spells {
		target := (int(success)-1)/num + 1
		search := sort.SearchInts(potions, target)
		ans[i] = len(potions) - search
	}
	return ans
}

func searchRange(nums []int, target int) []int {
	pos := sort.SearchInts(nums, target)
	var left, right int
	if pos < len(nums) && nums[pos] == target {
		left, right = pos, pos
		for ; right < len(nums) && nums[right] == target; right++ {

		}
		return []int{left, right - 1}
	} else {
		left, right = -1, -1
	}
	return []int{left, right}
}

func search(nums []int, target int) int {
	pos := sort.SearchInts(nums, target)
	if pos < len(nums) && nums[pos] == target {
		return pos
	}
	return -1
}

func numSmallerByFrequency(queries []string, words []string) []int {
	// 构造words 最小字母出现频次字典
	cntDict := make([][]int, len(words))
	for i, charCnt := range cntDict {
		charCnt = make([]int, 26)
		word := words[i]
		for _, ch := range word {
			charCnt[ch-'a'] += 1
		}
		for s := 0; s < len(charCnt); s++ {
			if charCnt[s] != 0 {
				charCnt = charCnt[s:]
				break
			}
		}
		cntDict[i] = charCnt
	}
	sort.Slice(cntDict, func(i, j int) bool {
		return cntDict[i][0] < cntDict[j][0]
	})
	ans := make([]int, len(queries))
	for i := 0; i < len(queries); i++ {
		query := queries[i]
		tempCnt := make([]int, 26)
		firstMin := 'z'
		for _, ch := range query {
			tempCnt[ch-'a'] += 1
			if ch < firstMin {
				firstMin = ch
			}
		}
		freq := tempCnt[firstMin-'a']

		pos := sort.Search(len(cntDict), func(j int) bool {
			return cntDict[j][0] > freq
		})
		ans[i] = len(cntDict) - pos
	}
	return ans
}

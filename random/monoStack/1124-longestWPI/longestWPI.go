package _124_longestWPI

// 1124. 表现良好的最长时间段
// 给你一份工作时间表 hours，上面记录着某一位员工每天的工作小时数。
//
// 我们认为当员工一天中的工作小时数大于 8 小时的时候，那么这一天就是「劳累的一天」。
//
// 所谓「表现良好的时间段」，意味在这段时间内，「劳累的天数」是严格 大于「不劳累的天数」。
//
// 请你返回「表现良好时间段」的最大长度。
//
// 示例 1：
//
// 输入：hours = [9,9,6,0,6,6,9]
// 输出：3
// 解释：最长的表现良好时间段是 [9,9,6]。
// 示例 2：
//
// 输入：hours = [6,6,6]
// 输出：0
//
// 提示：
//
// 1 <= hours.length <= 104
// 0 <= hours[i] <= 16

// 使用滑动窗计算
func longestWPI(hours []int) int {
	left, longest := 0, 0
	for _, v := range hours {
		if v > 8 {
			left += 1
		}
	}
	winCnt := 0
	for s, t := 0, 0; s <= t && t < len(hours); {
		if hours[t] > 8 {
			left -= 1
		}
		winCnt += sigMod(hours[t])
		if winCnt > 0 {
			longest = max(longest, t-s+1)
		} else {
			sPrev := s
			for ; s <= t && winCnt+left <= 0; s++ {
				winCnt -= sigMod(hours[s])
			}
			if s > sPrev {
				t = s
				winCnt, left = 0, 0
				for i := s; i < len(hours); i++ {
					if hours[i] > 8 {
						left += 1
					}
				}
				continue
			}
		}
		t += 1
	}
	return longest
}

func sigMod(duration int) int {
	if duration > 8 {
		return 1
	}
	return -1
}

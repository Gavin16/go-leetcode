package _266_countTexts

// 2266. 统计打字方案数
// Alice 在给 Bob 用手机打字。数字到字母的 对应 如下图所示。
//
// 为了 打出 一个字母，Alice 需要 按 对应字母 i 次，i 是该字母在这个按键上所处的位置。
//
// 比方说，为了按出字母 's' ，Alice 需要按 '7' 四次。类似的， Alice 需要按 '5' 两次得到字母  'k' 。
// 注意，数字 '0' 和 '1' 不映射到任何字母，所以 Alice 不 使用它们。
// 但是，由于传输的错误，Bob 没有收到 Alice 打字的字母信息，反而收到了 按键的字符串信息 。
//
// 比方说，Alice 发出的信息为 "bob" ，Bob 将收到字符串 "2266622" 。
// 给你一个字符串 pressedKeys ，表示 Bob 收到的字符串，请你返回 Alice 总共可能发出多少种文字信息 。
//
// 由于答案可能很大，将它对 109 + 7 取余 后返回。
//
// 示例 1：
// 输入：pressedKeys = "22233"
// 输出：8
// 解释：
// Alice 可能发出的文字信息包括：
// "aaadd", "abdd", "badd", "cdd", "aaae", "abe", "bae" 和 "ce" 。
// 由于总共有 8 种可能的信息，所以我们返回8 。
//
// 示例 2：
// 输入：pressedKeys = "222222222222222222222222222222222222"
// 输出：82876089
// 解释：
// 总共有 2082876103 种 Alice 可能发出的文字信息。
// 由于我们需要将答案对 109 + 7 取余，所以我们返回 2082876103 % (109 + 7) = 82876089 。
//
// 提示：
// 1 <= pressedKeys.length <= 105
// pressedKeys 只包含数字 '2' 到 '9' 。

func countTexts(pressedKeys string) int {
	digitMap := map[byte]int{
		'2': 3, '3': 3, '4': 3, '5': 3,
		'6': 3, '7': 4, '8': 3, '9': 4}

	mod := 1000_000_007
	repNums := cutRepeats(pressedKeys)
	ans := 1
	for _, repStr := range repNums {
		ch := repStr[0]
		num, N := digitMap[ch], len(repStr)

		dp := make([]int, N+1)
		dp[0], dp[1] = 1, 1
		for i := 2; i <= N; i++ {
			// 包含4个字母的按键,前导数量最长为 num-1
			for j := 1; j <= min(num, i); j++ {
				dp[i] += dp[i-j] % mod
			}
		}
		ans = (ans * dp[N]) % mod
	}
	return ans
}

func cutRepeats(keys string) []string {
	res := make([]string, 0)
	if len(keys) == 1 {
		return []string{keys}
	}
	i, j := 0, 1
	for j < len(keys) {
		if keys[i] == keys[j] {
			j++
		} else {
			res = append(res, keys[i:j])
			i, j = j, j+1
		}
	}
	res = append(res, keys[i:j])
	return res
}

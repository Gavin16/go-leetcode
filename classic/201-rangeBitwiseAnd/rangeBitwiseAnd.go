package _01_rangeBitwiseAnd

// 给你两个整数 left 和 right ，表示区间 [left, right] ，返回此区间内所有数字 按位与 的结果（包含 left 、right 端点）。
//
// 示例 1：
//
// 输入：left = 5, right = 7
// 输出：4
// 示例 2：
//
// 输入：left = 0, right = 0
// 输出：0
// 示例 3：
//
// 输入：left = 1, right = 2147483647
// 输出：0
//
//
// 提示：
//
// 0 <= left <= right <= 231 - 1

// 按位计算 : 超时
func rangeBitwiseAnd(left int, right int) int {
	res := left
	for i := left + 1; i <= right; i++ {
		res = res & i
	}
	return res
}

// left 到 right 的跨度较大
// 只要有一个数的对应为为0, 那么这一位就为0了
// 若 left 位数少于right位数, 那么left左侧少的部分将导致全部为0
// -- 例如 7: 111, 8: 1000 bits数的增加导致全部为0
// 若 left 位数和right位数相同, 那么最后保留下来的是公共前缀
// -- 例如: 9: 1001, 15: 1111, 靠左边的位存在差异时，右边的一定存在，因为右边的是低位，变化更频繁
// -- 因此只有当前缀相同时, 才认为是按位与不会被0覆盖

func rangeBitwiseAnd1(left int, right int) int {
	shift := 0
	for left < right {
		left >>= 1
		right >>= 1
		shift += 1
	}
	return left << shift
}

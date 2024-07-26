package _24_calculate

import (
	"strings"
)

// 224. 基本计算器
// 给你一个字符串表达式 s ，请你实现一个基本计算器来计算并返回它的值。
// 注意:不允许使用任何将字符串作为数学表达式计算的内置函数，比如 eval() 。
//
// 示例 1：
// 输入：s = "1 + 1"
// 输出：2
//
// 示例 2：
// 输入：s = " 2-1 + 2 "
// 输出：3
//
// 示例 3：
// 输入：s = "(1+(4+5+2)-3)+(6+8)"
// 输出：23
//
// 提示：
//
// 1 <= s.length <= 3 * 105
// s 由数字、'+'、'-'、'('、')'、和 ' ' 组成
// s 表示一个有效的表达式
// '+' 不能用作一元运算(例如， "+1" 和 "+(2 + 3)" 无效)
// '-' 可以用作一元运算(即 "-1" 和 "-(2 + 3)" 是有效的)
// 输入中不存在两个连续的操作符
// 每个数字和运行的计算将适合于一个有符号的 32位 整数
//

// 使用栈
// TODO 梳理计算过程
// 击败: 64.96%
func calculate(s string) int {
	ans := 0
	ops, sign := []int{1}, 1
	str := strings.Replace(s, " ", "", -1)

	for i := 0; i < len(str); {
		switch str[i] {
		case '+':
			sign = ops[len(ops)-1]
			i++
		case '-':
			sign = -1 * ops[len(ops)-1]
			i++
		case '(':
			ops = append(ops, sign)
			i++
		case ')':
			ops = ops[:len(ops)-1]
			i++
		default:
			num, j := 0, i
			for ; j < len(str) && str[j] >= '0' && str[j] <= '9'; j++ {
				num = num*10 + int(str[j]-'0')
			}
			i = j
			ans += sign * num
		}
	}
	return ans
}

package _7_addBinary

// 67. 二进制求和
// 给你两个二进制字符串 a 和 b ，以二进制字符串的形式返回它们的和。
//
// 示例 1：
//
// 输入:a = "11", b = "1"
// 输出："100"
// 示例 2：
//
// 输入：a = "1010", b = "1011"
// 输出："10101"
//
// 提示：
//
// 1 <= a.length, b.length <= 104
// a 和 b 仅由字符 '0' 或 '1' 组成
// 字符串如果不是 "0" ，就不含前导零
func addBinary(a string, b string) string {
	arrA, arrB := []byte(a), []byte(b)
	for i, j := 0, len(arrA)-1; i < j; {
		arrA[i], arrA[j] = arrA[j], arrA[i]
		i, j = i+1, j-1
	}
	for i, j := 0, len(arrB)-1; i < j; {
		arrB[i], arrB[j] = arrB[j], arrB[i]
		i, j = i+1, j-1
	}
	k, carry := 0, byte(0)
	ans := make([]byte, 0)
	for ; k < max(len(arrA), len(arrB)); k++ {
		bitA, bitB := byte('0'), byte('0')
		if k < len(arrA) {
			bitA = arrA[k]
		}
		if k < len(arrB) {
			bitB = arrB[k]
		}
		bitSum := ((bitA - '0') + (bitB - '0') + carry) % byte(2)
		carry = ((bitA - '0') + (bitB - '0') + carry) / byte(2)
		ans = append(ans, bitSum)
	}
	if carry > 0 {
		ans = append(ans, carry)
	}
	ansStr := ""
	for i := 0; i < len(ans); i++ {
		ansStr = string(ans[i]+'0') + ansStr
	}
	return ansStr
}

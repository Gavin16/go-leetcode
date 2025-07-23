package _37_singleNumber

// 给你一个整数数组 nums ，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次 。请你找出并返回那个只出现了一次的元素。
//
// 你必须设计并实现线性时间复杂度的算法且使用常数级空间来解决此问题。
//
// 示例 1：
//
// 输入：nums = [2,2,3,2]
// 输出：3
// 示例 2：
//
// 输入：nums = [0,1,0,1,0,1,99]
// 输出：99
//
// 提示：
//
// 1 <= nums.length <= 3 * 104
// -231 <= nums[i] <= 231 - 1
// nums 中，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次

func singleNumber(nums []int) int {
	bitSum := make([]int, 32)
	for i := 0; i < 32; i++ {
		for _, n := range nums {
			n >>= i
			bit := n & 1
			bitSum[i] += bit
		}
		bitSum[i] = bitSum[i] % 3
	}
	var singleNum = 0
	var sign = 1
	for i := 31; i >= 0; i-- {
		if i == 31 && bitSum[i] == 1 {
			sign *= -1
		}
		if sign == 1 {
			singleNum = singleNum*2 + bitSum[i]
		} else {
			singleNum = singleNum*2 + 1 - bitSum[i]
		}
	}
	if sign == -1 {
		singleNum += 1
	}
	return singleNum * sign
}

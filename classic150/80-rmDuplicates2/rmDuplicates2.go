package _0_rmDuplicates2

import (
	"sort"
)

// 80. 删除有序数组中的重复项 II
// 给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使得出现次数超过两次的元素只出现两次 ，返回删除后数组的新长度。
//
// 不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
//
// 说明：
//
// 为什么返回数值是整数，但输出的答案是数组呢？
//
// 请注意，输入数组是以「引用」方式传递的，这意味着在函数里修改输入数组对于调用者是可见的。
//
// 你可以想象内部操作如下:
//
// // nums 是以“引用”方式传递的。也就是说，不对实参做任何拷贝
// int len = removeDuplicates(nums);
//
// // 在函数里修改输入数组对于调用者是可见的。
// // 根据你的函数返回的长度, 它会打印出数组中 该长度范围内 的所有元素。
//
//	for (int i = 0; i < len; i++) {
//	   print(nums[i]);
//	}
//
// 示例 1：
//
// 输入：nums = [1,1,1,2,2,3]
// 输出：5, nums = [1,1,2,2,3]
// 解释：函数应返回新长度 length = 5, 并且原数组的前五个元素被修改为 1, 1, 2, 2, 3。 不需要考虑数组中超出新长度后面的元素。
//
// 示例 2：
// [0,0,1,1,1,1,2,2,2,3,3]
// 输入：nums = [0,0,1,1,1,1,2,3,3]
// 输出：7, nums = [0,0,1,1,2,3,3]
// 解释：函数应返回新长度 length = 7, 并且原数组的前七个元素被修改为 0, 0, 1, 1, 2, 3, 3。不需要考虑数组中超出新长度后面的元素。
//
// 提示：
//
// 1 <= nums.length <= 3 * 104
// -104 <= nums[i] <= 104
// nums 已按升序排列

// 解法1： 使用pre 和 post 分别表示从后向前的写入位置 和 读取位置
// 只关注 post位置写入重复数据是否超过2个即可
func removeDuplicates(nums []int) int {
	pre, post := 0, 0

	for post < len(nums) {
		cnt := 0
		for i := post; i < len(nums); i++ {
			if nums[i] == nums[post] {
				nums[pre] = nums[i]
				pre++
				cnt++
				if cnt == 2 {
					break
				}
			} else {
				break
			}
		}
		post = findPostNotEqual(nums, post)
	}
	return pre
}

func findPostNotEqual(nums []int, pos int) int {
	for i := pos; i < len(nums); i++ {
		if nums[i] != nums[pos] {
			return i
		}
	}
	return len(nums)
}

// 解法2: 使用map 对出现次数经计数: 击败 4.13%
func removeDuplicates1(nums []int) int {
	cntMap := make(map[int]int, len(nums)/2)
	for _, num := range nums {
		cntMap[num] = min(cntMap[num]+1, 2)
	}
	keys := make([]int, 0, len(cntMap))
	for key := range cntMap {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	gCnt := 0
	for _, num := range keys {
		cnt := cntMap[num]
		for i := 0; i < cnt; i++ {
			nums[gCnt] = num
			gCnt++
		}
	}
	return gCnt
}

// 解法3 一个for循环搞定
func removeDuplicates2(nums []int) int {
	slow := 0
	for fast, v := range nums {
		if fast < 2 || nums[slow-2] != v {
			nums[slow] = v
			slow++
		}
	}
	return slow
}

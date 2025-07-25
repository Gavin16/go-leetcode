package __findMedianSortedArrays

// 4. 寻找两个正序数组的中位数
// 给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
//
// 算法的时间复杂度应该为 O(log (m+n)) 。
//
// 示例 1：
//
// 输入：nums1 = [1,3], nums2 = [2]
// 输出：2.00000
// 解释：合并数组 = [1,2,3] ，中位数 2
// 示例 2：
//
// 输入：nums1 = [1,2], nums2 = [3,4]
// 输出：2.50000
// 解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
//
// 提示：
//
// nums1.length == m
// nums2.length == n
// 0 <= m <= 1000
// 0 <= n <= 1000
// 1 <= m + n <= 2000
// -106 <= nums1[i], nums2[i] <= 106

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	merged := mergeArrays(nums1, nums2)
	m, n := len(nums1), len(nums2)
	var midIdx int
	if (m+n)%2 == 1 {
		midIdx = (m + n - 1) / 2
		return float64(merged[midIdx])
	} else {
		left, right := (m+n)/2-1, (m+n)/2
		return (float64(merged[left]) + float64(merged[right])) / 2
	}
}

func mergeArrays(nums1 []int, nums2 []int) []int {
	m, n := len(nums1), len(nums2)
	merged := make([]int, m+n)
	i, j, k := 0, 0, 0
	for i < m && j < n {
		if nums1[i] < nums2[j] {
			merged[k] = nums1[i]
			i += 1
		} else {
			merged[k] = nums2[j]
			j += 1
		}
		k += 1
	}
	for i < m {
		merged[k] = nums1[i]
		i, k = i+1, k+1
	}
	for j < n {
		merged[k] = nums2[j]
		j, k = j+1, k+1
	}
	return merged
}

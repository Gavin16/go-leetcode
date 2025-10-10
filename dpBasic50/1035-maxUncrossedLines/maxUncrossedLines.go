package _035_maxUncrossedLines

// 1035. 不相交的线
// 在两条独立的水平线上按给定的顺序写下 nums1 和 nums2 中的整数。
//
// 现在，可以绘制一些连接两个数字 nums1[i] 和 nums2[j] 的直线，这些直线需要同时满足：
//
// nums1[i] == nums2[j]
// 且绘制的直线不与任何其他连线（非水平线）相交。
// 请注意，连线即使在端点也不能相交：每个数字只能属于一条连线。
//
// 以这种方法绘制线条，并返回可以绘制的最大连线数。
//
// 示例 1：
//
// 输入：nums1 = [1,4,2], nums2 = [1,2,4]
// 输出：2
// 解释：可以画出两条不交叉的线，如上图所示。
// 但无法画出第三条不相交的直线，因为从 nums1[1]=4 到 nums2[2]=4 的直线将与从 nums1[2]=2 到 nums2[1]=2 的直线相交。
// 示例 2：
//
// 输入：nums1 = [2,5,1,2,5], nums2 = [10,5,2,1,5,2]
// 输出：3
// 示例 3：
//
// 输入：nums1 = [1,3,7,1,7,5], nums2 = [1,9,2,5,1]
// 输出：2
//
// 提示：
//
// 1 <= nums1.length, nums2.length <= 500
// 1 <= nums1[i], nums2[j] <= 2000

// 问题可以直接转化为 求nums1, nums2 中相同的数字序列的最大长度
// 题设中提到连线不能相交 对应到序列上正好满足要求(序列匹配上之后都是从左往右的)
func maxUncrossedLines(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	f := make([][]int, m+1)
	for id := range f {
		f[id] = make([]int, n+1)
	}
	// 初始化 f[0][j] = 0, f[i][0] = 0
	for i, v1 := range nums1 {
		for j, v2 := range nums2 {
			if v1 == v2 {
				f[i+1][j+1] = f[i][j] + 1
			} else {
				f[i+1][j+1] = max(f[i+1][j], f[i][j+1])
			}
		}
	}
	return f[m][n]
}

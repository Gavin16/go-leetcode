package _49_maxPoints

// 149. 直线上最多的点数
// 给你一个数组 points ，其中 points[i] = [xi, yi] 表示 X-Y 平面上的一个点。求最多有多少个点在同一条直线上。
//
// 示例 1：
//
// 输入：points = [[1,1],[2,2],[3,3]]
// 输出：3
// 示例 2：
//
// 输入：points = [[1,1],[3,2],[5,3],[4,1],[2,3],[1,4]]
// 输出：4
//
// 提示：
//
// 1 <= points.length <= 300
// points[i].length == 2
// -104 <= xi, yi <= 104
// points 中的所有点 互不相同

// 两两判断斜率是否相等容易出现重复
// 采用选定两个点，然后找前面或者后面是否有第三个点和他们共线, 可以直接算出结果
func maxPoints(points [][]int) int {
	if len(points) <= 2 {
		return len(points)
	}
	ans, N := 0, len(points)

	for i := 1; i < N; i++ {
		for j := 0; j < i; j++ {
			x1, y1 := points[i][0], points[i][1]
			x2, y2 := points[j][0], points[j][1]
			temp := 2
			for k := 0; k < j; k++ {
				x3, y3 := points[k][0], points[k][1]
				if (y2-y1)*(x3-x2) == (y3-y2)*(x2-x1) {
					temp += 1
				}
			}
			ans = max(ans, temp)
		}
	}
	return ans
}

package _35_asteroidCollision

// 735. 小行星碰撞
// 给定一个整数数组 asteroids，表示在同一行的小行星。数组中小行星的索引表示它们在空间中的相对位置。
// 对于数组中的每一个元素，其绝对值表示小行星的大小，正负表示小行星的移动方向（正表示向右移动，负表示向左移动）。每一颗小行星以相同的速度移动。
//
// 找出碰撞后剩下的所有小行星。碰撞规则：两个小行星相互碰撞，较小的小行星会爆炸。
// 如果两颗小行星大小相同，则两颗小行星都会爆炸。两颗移动方向相同的小行星，永远不会发生碰撞。
//
// 示例 1：
// 输入：asteroids = [5,10,-5]
// 输出：[5,10]
// 解释：10 和 -5 碰撞后只剩下 10 。 5 和 10 永远不会发生碰撞。
// 示例 2：
//
// 输入：asteroids = [8,-8]
// 输出：[]
// 解释：8 和 -8 碰撞后，两者都发生爆炸。
// 示例 3：
//
// 输入：asteroids = [10,2,-5]
// 输出：[10]
// 解释：2 和 -5 发生碰撞后剩下 -5 。10 和 -5 发生碰撞后剩下 10 。
//
// 提示：
// 2 <= asteroids.length <= 104
// -1000 <= asteroids[i] <= 1000
// asteroids[i] != 0

func asteroidCollision(asteroids []int) []int {
	left, right := make([]int, 0), make([]int, 0)
	ansLeft := make([]int, 0)
	for i := 0; i < len(asteroids); i++ {
		num := asteroids[i]
		if num < 0 {
			if len(right) == 0 {
				left = append(left, num)
			} else {
				ansLeft = append(ansLeft, left...)
				left = left[:0]
				for len(right) > 0 && right[len(right)-1]+num < 0 {
					right = right[:len(right)-1]
				}
				if len(right) == 0 {
					left = append(left, num)
				} else if right[len(right)-1]+num == 0 {
					right = right[:len(right)-1]
				}
			}
		} else {
			right = append(right, num)
		}
	}
	ansLeft = append(ansLeft, left...)
	ans := append(ansLeft, right...)
	return ans
}

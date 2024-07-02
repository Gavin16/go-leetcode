package canJump

import (
	"testing"
)

func TestCanJump(t *testing.T) {

	nums := []int{2, 3, 1, 1, 4}
	jump := canJump2(nums)
	println(jump)

	nums1 := []int{3, 2, 1, 0, 4}
	jump1 := canJump2(nums1)
	println(jump1)
}

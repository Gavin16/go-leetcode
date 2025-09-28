package _34_increasingTriplet

import "testing"

func TestIncreasingTriplet(t *testing.T) {

	nums1 := []int{1, 2, 3, 4, 5}
	nums2 := []int{5, 4, 3, 2, 1}
	nums3 := []int{2, 1, 5, 0, 4, 6}
	nums4 := []int{20, 100, 10, 12, 5, 13}

	println(increasingTriplet(nums1))
	println(increasingTriplet(nums2))
	println(increasingTriplet(nums3))
	println(increasingTriplet(nums4))

}

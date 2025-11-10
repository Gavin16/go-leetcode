package _435_numberOfPaths

import (
	"fmt"
	"reflect"
	"slices"
	"testing"
)

func TestNumberOfPaths(t *testing.T) {
	//grid, k := [][]int{{5, 2, 4}, {3, 0, 5}, {0, 7, 2}}, 3
	//grid, k := [][]int{{3, 2}, {3, 2}}, 3
	//fmt.Println(numberOfPaths(grid, k))

	tests := []struct {
		nums []int
		k, x int
		want []int
	}{
		{[]int{1, 1, 2, 2, 3, 4, 2, 3}, 8, 3, []int{14}},
		{[]int{1, 1, 2, 2, 3, 4, 2, 3}, 6, 2, []int{6, 10, 12}},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("findXSum run: %d", i), func(t *testing.T) {
			if got := findXSum(tt.nums, tt.k, tt.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findXSum result %#v, want %#v", got, tt.want)
			}
		})
	}

}

func findXSum(nums []int, k int, x int) []int {
	n := len(nums)
	split := make([][]int, n-k+1)
	answer := make([]int, n-k+1)
	for i := range split {
		split[i] = nums[i : i+k]
		dict := map[int]int{}
		for _, v := range split[i] {
			dict[v] = dict[v] + 1
		}
		cnt := make([][2]int, 0)
		for key, val := range dict {
			cnt = append(cnt, [2]int{key, val})
		}
		slices.SortFunc(cnt, func(a, b [2]int) int {
			if b[1] == a[1] {
				return b[0] - a[0]
			}
			return b[1] - a[1]
		})
		for _, pair := range cnt[:min(x, len(cnt))] {
			answer[i] += pair[0] * pair[1]
		}
	}
	return answer
}

package _95_MedianFinder

import "testing"

func TestMedianFinder(t *testing.T) {

	finder1 := Constructor()
	finder1.AddNum(4)
	println(finder1.FindMedian())
	finder1.AddNum(3)
	println(finder1.FindMedian())
	finder1.AddNum(2)
	println(finder1.FindMedian())
	finder1.AddNum(5)
	println(finder1.FindMedian())
	finder1.AddNum(7)
	println(finder1.FindMedian())
}

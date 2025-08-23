package _80_RandomizedSet

import (
	"testing"
)

func TestRandomizedSet(t *testing.T) {
	randomizedSet := Constructor()

	println(randomizedSet.Remove(0))
	println(randomizedSet.Remove(0))
	println(randomizedSet.Insert(0))
	println(randomizedSet.GetRandom())
	println(randomizedSet.Remove(0))
	println(randomizedSet.Insert(0))

}

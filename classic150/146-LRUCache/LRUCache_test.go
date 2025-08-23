package _46_LRUCache

import "testing"

func TestLRUCache(t *testing.T) {

	//cache := Constructor(2)
	//cache.Put(1, 1)
	//cache.Put(2, 2)
	//println(cache.Get(1))
	//cache.Put(3, 3)
	//println(cache.Get(2))
	//cache.Put(4, 4)
	//println(cache.Get(1))
	//println(cache.Get(3))
	//println(cache.Get(4))

	//cache := Constructor(1)
	//cache.Put(1, 1)
	//cache.Put(2, 2)
	//println(cache.Get(1))
	//
	//cache.Put(3, 3)
	//println(cache.Get(2))
	//cache.Put(4, 4)
	//println(cache.Get(1))
	//println(cache.Get(3))
	//println(cache.Get(4))

	cache := Constructor(2)
	println(cache.Get(2))
	cache.Put(2, 6)

	println(cache.Get(1))
	cache.Put(1, 5)
	cache.Put(1, 2)

	println(cache.Get(1))
	println(cache.Get(2))

}

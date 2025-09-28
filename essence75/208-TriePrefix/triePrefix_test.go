package _08_TriePrefix

import "testing"

func TestTriePrefix(t *testing.T) {
	trieTree := Constructor()
	trieTree.Insert("app")
	trieTree.Insert("apple")
	trieTree.Insert("beer")
	trieTree.Insert("add")
	trieTree.Insert("jam")
	trieTree.Insert("rental")
	println(trieTree.Search("apps"))
	println(trieTree.Search("app"))
	println(trieTree.StartsWith("app"))
	trieTree.Insert("app")
	println(trieTree.Search("app"))
}

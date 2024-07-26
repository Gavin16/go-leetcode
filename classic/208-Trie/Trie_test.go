package _08_Trie

import "testing"

func TestTrie(t *testing.T) {

	//trie := Constructor()
	//trie.Insert("app")
	//trie.Insert("apple")
	//println(trie.Search("app"))
	//println(trie.Search("appl"))
	//println(trie.Search("apple"))
	//println(trie.StartsWith("appl"))

	trie1 := Constructor()
	trie1.Insert("h")
	trie1.Insert("he")
	println(trie1.StartsWith("h"))
	trie1.Insert("hello")
	println(trie1.Search("hell"))
	println(trie1.StartsWith("hell"))
	trie1.Insert("hell")
	println(trie1.Search("hell"))

}

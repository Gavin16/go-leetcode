package _11_WordDictionary

import "testing"

func TestWordDictionary(t *testing.T) {

	dict := Constructor()
	dict.AddWord("bad")
	dict.AddWord("dad")
	dict.AddWord("mad")
	dict.AddWord("maddan")

	println(dict.Search("pad"))
	println(dict.Search("bad"))
	println(dict.Search(".ad"))
	println(dict.Search("b.."))
	println(dict.Search("m.d.a."))
	println(dict.Search("mad..."))

}

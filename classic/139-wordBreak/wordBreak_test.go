package _39_wordBreak

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWorkBreak(t *testing.T) {
	dict := []string{"leet", "code"}
	dictSet := map[string]bool{}
	for _, word := range dict {
		dictSet[word] = true
	}

	println(dictSet["hello"])
	println(dictSet["leet"])
	
	tests := []struct {
		name     string
		str      string
		wordDict []string
		want     bool
	}{
		{"1", "leetcode", []string{"leet", "code"}, true},
		{"2", "applepenapple", []string{"apple", "pen"}, true},
		{"3", "catsandog", []string{"cats", "dog", "sand", "and", "cat"}, false},
		{"4", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab",
			[]string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}, false},
		{"5", "aaaaaaa", []string{"aaa", "aaaa"}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := wordBreak(tt.str, tt.wordDict)
			assert.Equal(t, tt.want, actual)
		})
	}
}

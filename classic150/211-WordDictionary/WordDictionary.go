package _11_WordDictionary

// WordDictionary 211. 添加与搜索单词 - 数据结构设计
// 请你设计一个数据结构，支持 添加新单词 和 查找字符串是否与任何先前添加的字符串匹配 。
//
// 实现词典类 WordDictionary ：
//
// WordDictionary() 初始化词典对象
// void addWord(word) 将 word 添加到数据结构中，之后可以对它进行匹配
// bool search(word) 如果数据结构中存在字符串与 word 匹配，则返回 true ；否则，返回  false 。word 中可能包含一些 '.' ，每个 . 都可以表示任何一个字母。
//
// 示例：
//
// 输入：
// ["WordDictionary","addWord","addWord","addWord","search","search","search","search"]
// [[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]
// 输出：
// [null,null,null,null,false,true,true,true]
//
// 解释：
// WordDictionary wordDictionary = new WordDictionary();
// wordDictionary.addWord("bad");
// wordDictionary.addWord("dad");
// wordDictionary.addWord("mad");
// wordDictionary.search("pad"); // 返回 False
// wordDictionary.search("bad"); // 返回 True
// wordDictionary.search(".ad"); // 返回 True
// wordDictionary.search("b.."); // 返回 True
//
// 提示：
//
// 1 <= word.length <= 25
// addWord 中的 word 由小写英文字母组成
// search 中的 word 由 '.' 或小写英文字母组成
// 最多调用 104 次 addWord 和 search

// WordDictionary 击败:13.18%
type WordDictionary struct {
	char     byte
	flag     bool
	children map[byte]*WordDictionary
}

func Constructor() WordDictionary {
	return WordDictionary{char: 0, flag: false, children: make(map[byte]*WordDictionary)}
}

func (dict *WordDictionary) AddWord(word string) {
	children := dict.children
	for i := 0; i < len(word); i++ {
		w := word[i]
		if children[w] == nil {
			children[w] = &WordDictionary{char: w, children: make(map[byte]*WordDictionary)}
		}
		if i == len(word)-1 {
			children[w].flag = true
		}
		children = children[w].children
	}
}

func (dict *WordDictionary) Search(word string) bool {
	children := dict.children
	return dfsSearch(children, word)
}

func dfsSearch(dict map[byte]*WordDictionary, word string) bool {
	if len(word) == 1 {
		w := word[0]
		if dict[w] != nil && dict[w].flag {
			return true
		}
		if w == '.' {
			for _, child := range dict {
				if child.flag {
					return true
				}
			}
		}
		return false
	}
	w, exist := word[0], false
	if w != '.' {
		if dict[w] == nil {
			return false
		}
		exist = exist || dfsSearch(dict[w].children, word[1:])
	} else {
		for _, child := range dict {
			if child.children == nil {
				continue
			}
			exist = exist || dfsSearch(child.children, word[1:])
		}
	}
	return exist
}

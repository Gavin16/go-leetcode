package _7_fullJustify

import "strings"

// 68. 文本左右对齐
// 给定一个单词数组 words 和一个长度 maxWidth ，重新排版单词，使其成为每行恰好有 maxWidth 个字符，且左右两端对齐的文本。
// 你应该使用 “贪心算法” 来放置给定的单词；也就是说，尽可能多地往每行中放置单词。必要时可用空格 ' ' 填充，使得每行恰好有 maxWidth 个字符。
//
// 要求尽可能均匀分配单词间的空格数量。如果某一行单词间的空格不能均匀分配，则左侧放置的空格数要多于右侧的空格数。
// 文本的最后一行应为左对齐，且单词之间不插入额外的空格。
//
// 注意:
//
// 单词是指由非空格字符组成的字符序列。
// 每个单词的长度大于 0，小于等于 maxWidth。
// 输入单词数组 words 至少包含一个单词。
//
// 示例 1:
//
// 输入: words = ["This", "is", "an", "example", "of", "text", "justification."], maxWidth = 16
// 输出:
// [
//
//	"This    is    an",
//	"example  of text",
//	"justification.  "
// ]
// 示例 2:
//
// 输入:words = ["What","must","be","acknowledgment","shall","be"], maxWidth = 16
// 输出:
// [
//
//	"What   must   be",
//	"acknowledgment  ",
//	"shall be        "
//
// ]
// 解释: 注意最后一行的格式应为 "shall be    " 而不是 "shall     be",
//
//	因为最后一行应为左对齐，而不是左右两端对齐。
//	第二行同样为左对齐，这是因为这行只包含一个单词。
//
// 示例 3:
//
// 输入:words = ["Science","is","what","we","understand","well","enough","to","explain","to","a","computer.","Art","is","everything","else","we","do"]，maxWidth = 20
// 输出:
// [
//
//	"Science  is  what we",
//	"understand      well",
//	"enough to explain to",
//	"a  computer.  Art is",
//	"everything  else  we",
//	"do                  "
//
// ]
//
// 提示:
//
// 1 <= words.length <= 300
// 1 <= words[i].length <= 20
// words[i] 由小写英文字母和符号组成
// 1 <= maxWidth <= 100
// words[i].length <= maxWidth

// 题解思路
// 正序遍历 累加切片中单词长度(累加判断方式为 总单词长度 + 空格个数[n-1] 是否大于 maxWidth)
// 若长度超过 maxWidth, 则计数 和 累加和 去掉当前值，判断长度是否为1 来采用不同的填充策略
//
// 时间复杂度 击败: 100.00%
func fullJustify(words []string, maxWidth int) []string {
	start, stop, cnt := 0, 0, 0
	lenSum, length := 0, len(words)

	result := make([]string, 0)
	for stop < length {
		lenSum, cnt = lenSum+len(words[stop]), cnt+1

		if lenSum+cnt-1 > maxWidth {
			// 去掉当前stop位置
			cnt, lenSum = cnt-1, lenSum-len(words[stop])
			slice := words[start:stop]

			padLen := maxWidth - lenSum
			if cnt == 1 {
				// 末尾填充
				for k := 0; k < padLen; k++ {
					slice[0] = slice[0] + " "
				}
			} else {
				// 均匀填充空格
				for i := 0; padLen > 0; i++ {
					index := i % cnt
					if index != 0 {
						slice[index], padLen = " "+slice[index], padLen-1
					}
				}
			}
			result = append(result, strings.Join(slice, ""))

			// 重置状态
			start, cnt, lenSum = stop, 0, 0
			continue
		}
		stop += 1
	}
	lastSlice := words[start:stop]
	padLen, padStr := maxWidth-lenSum-(cnt-1), ""
	for i := 0; i < padLen; i++ {
		padStr = padStr + " "
	}
	return append(result, strings.Join(lastSlice, " ")+padStr)
}

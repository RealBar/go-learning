package main

//给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。
//
//示例 1：
//
//输入: "babad"
//输出: "bab"
//注意: "aba" 也是一个有效答案。
//示例 2：
//
//输入: "cbbd"
//输出: "bb"
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/longest-palindromic-substring
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func longestPalindrome(s string) string {
	runes := []rune(s)
	var max string
	for i, _ := range runes {
		len1 := expandCenter(runes, i, i)
		len2 := expandCenter(runes, i, i+1)
		if len1 > len(max) {
			max = string(runes[i-len1/2 : i+len1/2+1])
		}
		if len2 > len(max) {
			max = string(runes[i-len2/2+1 : i+len2/2+1])
		}
	}
	return max
}

func expandCenter(runes []rune, l int, r int) int {
	for l >= 0 && r < len(runes) && runes[l] == runes[r] {
		l--
		r++
	}
	return r - l - 1
}

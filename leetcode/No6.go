package main

//比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：
//
//L   C   I   R
//E T O E S I I G
//E   D   H   N
//之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。
//
//请你实现这个将字符串进行指定行数变换的函数：
//
//string convert(string s, int numRows);
//示例 1:
//
//输入: s = "LEETCODEISHIRING", numRows = 3
//输出: "LCIRETOESIIGEDHN"
//示例 2:
//
//输入: s = "LEETCODEISHIRING", numRows = 4
//输出: "LDREOEIIECIHNTSG"
//解释:
//
//L     D     R
//E   O E   I I
//E C   I H   N
//T     S     G
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/zigzag-conversion
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func convert(s string, numRows int) string {
	if numRows == 1{
		return s
	}
	runes := []rune(s)
	lines := make([][]rune, numRows)
	for i := range lines {
		lines[i] = make([]rune, 0)
	}
	var goingDown = false
	var res = make([]rune, 0,len(runes))
	var row = 0
	for _, r := range runes {
		lines[row] = append(lines[row], r)
		if row == 0 || row == numRows-1 {
			goingDown = !goingDown
		}
		if goingDown {
			row++
		} else {
			row--
		}
	}
	for _,l := range lines{
		res = append(res,l...)
	}
	return string(res)
}

package main

import "fmt"

// 寻找最长不含有重复字符的子串
// 对于每一个字母X，lastOccurred[x]不存在，或者 < start -> 无需操作
// lastOccurred[x] >= start 更新start
// 更新lastOccurred[x],更新maxLength
func lengthOfNonRepeatingSubStr(s string) int {
	//lastOccured := make(map[byte]int) // 记录每个字符最后出现的位置
	lastOccured := make(map[rune]int) // 支持中文
	start, maxLength := 0, 0
	//for i, ch := range []byte(s) {
	for i, ch := range []rune(s) { // 支持中文
		lastI, ok := lastOccured[ch]
		if ok && lastI >= start {
			// 将位置移动到上一次开始的字符位置的下一位,(abca -> bca, start为b所在位置)
			start = lastOccured[ch] + 1
		}
		//fmt.Println(lastOccured, ok, lastI, start, maxLength, i)
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccured[ch] = i
	}
	return maxLength
}

func main() {
	fmt.Println(
		lengthOfNonRepeatingSubStr("abcabcbb"))
	fmt.Println(
		lengthOfNonRepeatingSubStr("bbb"))
	fmt.Println(
		lengthOfNonRepeatingSubStr("pwwkew"))
	fmt.Println(
		lengthOfNonRepeatingSubStr(""))
	fmt.Println(
		lengthOfNonRepeatingSubStr("b"))
	fmt.Println(
		lengthOfNonRepeatingSubStr("abcdefg"))
	fmt.Println(
		lengthOfNonRepeatingSubStr("黑化肥会发会发挥"))
}

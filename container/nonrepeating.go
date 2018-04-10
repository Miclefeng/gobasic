package main

import (
	"fmt"
)

// 寻找最长不含有重复字符的子串
// 对于每一个字母X，lastOccurred[x]不存在，或者 < start -> 无需操作
// lastOccurred[x] >= start 更新start
// 更新lastOccurred[x],更新maxLength
func lengthOfNonRepeatingSubStr(s string) int {
	lastOccured := make(map[byte]int) // 记录每个字符最后出现的位置
	start, maxLength := 0, 0
	for i, ch := range []byte(s) {
		lastI, ok := lastOccured[ch]
		if ok && lastI >= start {
			start = lastOccured[ch] + 1
		}
		//fmt.Println(lastOccured, ok, lastI, start, maxLength, i)
		if i - start + 1 > maxLength {
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
}

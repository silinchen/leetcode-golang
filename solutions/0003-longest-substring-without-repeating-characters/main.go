package main

import (
	"fmt"
	"math"
	"strings"
)

func maxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
/**
 * Runtime: 8~12ms
 * Memory: 3.2~3.6MB
 */
// 方法三：优化滑动窗口
func lengthOfLongestSubstring(s string) int {
	var ans = 0
	var start = 0
	var sMap = make(map[rune]int)
	for j, str := range s {
		if i, ok := sMap[str]; ok {
			start = maxInt(i, start)
		}
		ans = maxInt(ans, j-start+1)
		sMap[str] = j + 1
	}
	return ans
}
/**
 * Runtime: 0~4ms
 * Memory: 2.6MB
 */
// 方法二：滑动窗口
func lengthOfLongestSubstring2(s string) int {
	substr, ans, i, j, n := "", 0, 0, 0, len(s)
	for i < n && j < n {
		ch := s[j : j+1]
		if index := strings.Index(substr, ch); index == -1 {
			j = j + 1
			substr = s[i:j]
			ans = maxInt(ans, j-i)
		} else {
			i = i + 1
			substr = substr[1:]
		}
	}
	return ans
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))    // 3
	fmt.Println(lengthOfLongestSubstring("bbbbb"))       // 1
	fmt.Println(lengthOfLongestSubstring("pwwkew"))      // 3
	fmt.Println(lengthOfLongestSubstring("你好，世界！世界，你好")) // 6

	fmt.Println(lengthOfLongestSubstring2("abcabcbb")) // 3
	fmt.Println(lengthOfLongestSubstring2("pwwkew"))   // 3
}


